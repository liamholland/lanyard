package middleware

import (
	"lanyard/server/internal/store"
	"log"
	"net/http"
)

const LanyardIdentityHeader string = "Lanyard-Identity"

func IdentityStampedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		identityHeader := r.Header.Get(LanyardIdentityHeader)
		if identityHeader != "" {
			if identity := store.CheckIdentity(identityHeader, r.RemoteAddr); identity > -1 {
				log.Print("request is from user with id: ", identity)
			} else {
				log.Print("request identity ", identity, " is invalid")
			}
		}

		next.ServeHTTP(w, r)
	})
}
