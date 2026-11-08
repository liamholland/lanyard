package health

import (
	"io"
	"net/http"
)

const BasePath string = "/health/check"

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello gamer\n")
}
