package time

import (
	"encoding/json"
	"lanyard/pkg/lan_party/departures"
	"lanyard/pkg/lan_party/options"
	"lanyard/server/internal/store"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetStartTime(w http.ResponseWriter, r *http.Request) {
	optionValue := options.GetOptionValue(store.GetPartyConfig(), options.StartDate)
	if optionValue == nil {
		http.Error(w, "no value found for start date option", http.StatusNotFound)
		return
	}

	startTime, ok := optionValue.(time.Time)
	if !ok {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	res := SimpleTimeFormat{Time: startTime.Unix()}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func AddDepartureTime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SimpleTimeFormat

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request for departure time; must be a simple POSIX timestamp", http.StatusBadRequest)
		return
	}

	uploadErr := departures.AddDepartureTime(store.GetPartyConfig(), req.Time)
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
	n := departures.HowManyAtTime(store.GetPartyConfig(), &t)

	res := HowManyResponse{
		Time:      p,
		NumPeople: n,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}
