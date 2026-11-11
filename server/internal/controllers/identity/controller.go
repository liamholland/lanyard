package identity

import (
	"encoding/json"
	"lanyard/pkg/store"
	"log"
	"net/http"
	"strconv"
)

func CreateUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req IdentityFormat

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Username == "" {
		http.Error(w, "invalid identity request", http.StatusBadRequest)
		return
	}

	identity, parseErr := strconv.ParseInt(r.Header.Get("Lanyard-Identity"), 10, 64)
	if parseErr != nil {
		http.Error(w, "invalid identity request", http.StatusBadRequest)
		return
	}

	store.RegisterUsername(identity, req.Username)

	w.WriteHeader(http.StatusCreated)
}

func RetrieveUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req IdentityFormat

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Print("username request is malformed: ", err)
		http.Error(w, "invalid identity request", http.StatusBadRequest)
		return
	}

	identity, parseErr := strconv.ParseInt(req.Identity, 10, 64)
	if parseErr != nil {
		http.Error(w, "invalid identity request", http.StatusBadRequest)
		return
	}

	res := IdentityFormat{Username: store.GetUsername(identity)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
