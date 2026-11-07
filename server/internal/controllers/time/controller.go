package time

import (
	"encoding/json"
	"lanyard/pkg/lan_party/options"
	"lanyard/server/internal/logging"
	"lanyard/server/internal/store"
	"net/http"
	"time"
)

const BasePath string = "/time"

type TimeResponse struct {
	Time int64 `json:"time"`
}

func GetStartTime(w http.ResponseWriter, r *http.Request) {
	logging.LogControllerRequest(BasePath+"/start", "GET", r.RemoteAddr)

	optionValue := options.GetOptionValue(store.GetPartyConfig(), options.StartDate)
	if optionValue == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	startTime, ok := optionValue.(time.Time)
	if !ok {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	res := TimeResponse{Time: startTime.Unix()}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
