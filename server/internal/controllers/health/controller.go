package health

import (
	"io"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello gamer\n")
}
