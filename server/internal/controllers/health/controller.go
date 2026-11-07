package health

import (
	"io"
	"lanyard/server/internal/logging"
	"net/http"
)

const BasePath string = "/health/check"

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	logging.LogControllerRequest(BasePath, "GET", r.RemoteAddr)
	io.WriteString(w, "hello gamer\n")
}
