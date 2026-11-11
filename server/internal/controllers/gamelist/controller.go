package gamelist

import (
	"encoding/json"
	gl "lanyard/pkg/lan_party/gamelist"
	opts "lanyard/pkg/lan_party/options"
	"lanyard/pkg/store"
	"lanyard/server/internal/format"
	"lanyard/server/internal/middleware"
	"net/http"
)

func ResolveGames(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req format.GameResolutionRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request for games list; must be a map of game names and install sizes", http.StatusBadRequest)
		return
	}

	glPath, ok := opts.GetOptionValue(store.GetParty(), opts.GameListPath).(string)
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

	userId, idErr := middleware.GetIdentity(r)
	if idErr == nil {
		store.UpdateUserResolution(userId, *resolution)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resolution)
}
