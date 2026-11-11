package time

import (
	"encoding/json"
	"lanyard/pkg/lan_party/departures"
	"lanyard/pkg/lan_party/gamelist"
	"lanyard/pkg/lan_party/options"
	"lanyard/pkg/store"
	"lanyard/server/internal/format"
	"lanyard/server/internal/middleware"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetStartTime(w http.ResponseWriter, r *http.Request) {
	optionValue := options.GetOptionValue(store.GetParty(), options.StartDate)
	if optionValue == nil {
		http.Error(w, "no value found for start date option", http.StatusNotFound)
		return
	}

	startTime, ok := optionValue.(time.Time)
	if !ok {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	res := format.SimpleTimeFormat{Time: startTime.Unix()}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func AddDepartureTime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req format.SimpleTimeFormat

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request for departure time; must be a simple POSIX timestamp", http.StatusBadRequest)
		return
	}

	var assocId int64 = 0
	userId, idErr := middleware.GetIdentity(r)
	if idErr == nil {
		assocId = userId
	}

	uploadErr := departures.AddDepartureTime(store.GetParty(), req.Time, assocId)
	if uploadErr != nil {
		http.Error(w, "bad request for departure time; must be a simple POSIX timestamp", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetHowManyAtTime(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	if len(queryParams) < 1 || queryParams.Get("time") == "" {
		log.Print("did not find 'time' parameter on request")
		http.Error(w, "'time' is a required qeury parameter", http.StatusBadRequest)
		return
	}

	v := queryParams.Get("time")
	p, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Print("time is not a POSIX timestamp: ", p)
		http.Error(w, "'time' should be a POSIX timestamp", http.StatusBadRequest)
		return
	}

	t := time.Unix(p, 0)
	n := departures.HowManyAtTime(store.GetParty(), &t)

	res := format.HowManyResponse{
		Time:      p,
		NumPeople: n,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}

func GetWhenToPlayGame(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	if len(queryParams) < 1 || queryParams.Get("game") == "" {
		log.Print("did not find 'game' parameter on request")
		http.Error(w, "'game' is a required qeury parameter", http.StatusBadRequest)
		return
	}

	g := queryParams.Get("game")

	glPath, ok := options.GetOptionValue(store.GetParty(), options.GameListPath).(string)
	if !ok {
		log.Print("no games list")
		http.Error(w, "no games list on server", http.StatusNotFound)
		return
	}

	numP, err := gamelist.GetIdealPlayersForGame(g, glPath)
	if err != nil || numP < 0 {
		log.Print("no ideal players for game, finding last time where anyone is around")
		numP = 1
	}

	res := format.HowManyResponse{
		Time:      departures.GetTimeWithNPeople(store.GetParty(), numP).Unix(),
		NumPeople: numP,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}
