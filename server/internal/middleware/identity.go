package middleware

import (
	"lanyard/server/internal/controllers/identity"
	"lanyard/server/internal/store"
	"log"
	"net/http"
)

const LanyardIdentityHeader string = "Lanyard-Identity"

var identityRequired map[string]bool = map[string]bool{
	identity.BasePath + "/register": true,
	identity.BasePath + "/retrieve": true,
}

func IdentityStampedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		identityHeader := r.Header.Get(LanyardIdentityHeader)
		if identityHeader != "" {
			if identity := store.CheckIdentity(identityHeader, r.RemoteAddr); identity > -1 {
				log.Print("request is from user with id: ", identity)
			} else {
				if _, ok := identityRequired[r.URL.Path]; ok {
					http.Error(w, r.URL.Path+" requires an identity", http.StatusForbidden)
					return
				}

				log.Print("request identity ", identity, " is missing or invalid")
			}
		}

		next.ServeHTTP(w, r)
	})
}
