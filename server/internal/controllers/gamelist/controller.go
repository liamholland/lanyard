package gamelist

import (
	"encoding/json"
	gl "lanyard/pkg/lan_party/gamelist"
	opts "lanyard/pkg/lan_party/options"
	"lanyard/server/internal/store"
	"log"
	"net/http"
)

var BasePath string = "/games"

type GameResolutionRequest struct {
	Games    map[string]int64 `json:"games"`
	Resolved map[string]bool  `json:"resolved"`
}

func ResolveGames(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req GameResolutionRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request for games list; must be a map of game names and install sizes", http.StatusBadRequest)
		return
	}

	glPath, ok := opts.GetOptionValue(store.GetPartyConfig(), opts.GameListPath).(string)
	if !ok {
		http.Error(w, "no games list configured on server", http.StatusBadGateway)
		return
	}

	list, glErr := gl.ReadGameList(glPath)
	if glErr != nil {
		http.Error(w, "no games list configured on server", http.StatusBadGateway)
		return
	}

	resolution := gl.ResolveGameListFromMap(list, req.Games, req.Resolved)

	ctxId := r.Context().Value("lan-id")
	if userId, ok := ctxId.(int64); ok {
		store.UpdateUserResolution(userId, *resolution)
	} else {
		log.Print("user id on context is invalid, skipping save of user resolution: ", ctxId)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resolution)
}
