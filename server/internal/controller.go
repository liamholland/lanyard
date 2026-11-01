package internal

import (
	"io"
	"log"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	log.Print("received health check request from: ", r.RemoteAddr)
	io.WriteString(w, "hello gamer\n")
}
