package admin

import (
	"encoding/json"
	"fmt"
	opts "lanyard/pkg/lan_party/options"
	"lanyard/server/internal/store"
	"log"
	"net/http"
	"strings"
)

func UpdateOption(w http.ResponseWriter, r *http.Request) {
	if !isFromLocalHost(r.RemoteAddr) {
		http.Error(w, "ask the server admin to do this", http.StatusForbidden)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req OptionUpdateRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Print("option update request was malformed")
		http.Error(w, "bad option request", http.StatusBadRequest)
		return
	}

	option := opts.GetOptionFromName(req.OptionName)
	if option == nil {
		log.Print(req.OptionName, " is not a valid option name")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newSettings := opts.Copy(store.GetPartyConfig())
	validErr := option.OnRead(newSettings, req.OptionValue)
	if validErr != nil {
		log.Printf("option value %v is invalid for option %s: %s", req.OptionValue, req.OptionName, validErr.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Print("value for ", req.OptionName, " validated; updating settings")

	store.SetPartyConfig(newSettings)

	log.Print("settings successfully updated!")
	w.WriteHeader(http.StatusOK)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if !isFromLocalHost(r.RemoteAddr) {
		http.Error(w, "ask the server admin to do this", http.StatusForbidden)
		return
	}

	identities := store.GetUsers()
	usernames := store.GetUsernames()

	userIps := func() map[string]string {
		m := make(map[string]string)
		for id, name := range usernames {
			m[name] = identities[fmt.Sprintf("%v", id)]
		}
		return m
	}()

	res := ListUsersResponse{Users: userIps}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}

func isFromLocalHost(addr string) bool {
	log.Print("checking if ", addr, " is a local host")

	parts := strings.Split(addr, ":")
	host := strings.Join(parts[:len(parts)-1], "")

	log.Print("host is ", host)

	return host == "127.0.0.1" || host == "[1]"
}
