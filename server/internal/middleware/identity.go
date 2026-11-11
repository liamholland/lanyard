package middleware

import (
	"context"
	"lanyard/pkg/store"
	"lanyard/server/internal/controllers/gamelist"
	"lanyard/server/internal/controllers/identity"
	"log"
	"net/http"
)

const LanyardIdentityHeader string = "Lanyard-Identity"

var identityRequired map[string]bool = map[string]bool{
	identity.BasePath + "/register": true,
	identity.BasePath + "/retrieve": true,
	gamelist.BasePath + "/resolve":  true,
}

func IdentityStampedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		identityHeader := r.Header.Get(LanyardIdentityHeader)
		if identityHeader != "" {
			if identity := store.CheckIdentity(identityHeader, r.RemoteAddr); identity > -1 {
				log.Print("request is from user with id: ", identity)
				r = withIdentity(r, identity)
			} else {
				if _, ok := identityRequired[r.URL.Path]; ok {
					http.Error(w, r.URL.Path+" requires an identity", http.StatusUnauthorized)
					return
				}

				log.Print("request identity ", identity, " is missing or invalid")
			}
		}

		next.ServeHTTP(w, r)
	})
}

func withIdentity(r *http.Request, identity int64) *http.Request {
	identityCtx := context.WithValue(r.Context(), "lan-id", identity)
	return r.WithContext(identityCtx)
}
