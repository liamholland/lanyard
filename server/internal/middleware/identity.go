package middleware

import (
	"context"
	"errors"
	"lanyard/pkg/store"
	"lanyard/server/internal/format"
	"log"
	"net/http"
)

const LanyardIdentityHeader string = "Lanyard-Identity"

var identityRequired map[string]bool = map[string]bool{
	format.IdentityRegisterPath: true,
	format.IdentityRetrievePath: true,
	format.GamesResolvePath:     true,
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

func GetIdentity(r *http.Request) (int64, error) {
	ctxId := r.Context().Value("lan-id")
	if userId, ok := ctxId.(int64); ok {
		return userId, nil
	} else {
		log.Print("user id on context is invalid, skipping save of user resolution: ", ctxId)
		return -1, errors.New("no identity on request")
	}
}

func withIdentity(r *http.Request, identity int64) *http.Request {
	identityCtx := context.WithValue(r.Context(), "lan-id", identity)
	return r.WithContext(identityCtx)
}
