package internal

import (
	"embed"
	"errors"
	"io/fs"
	cfg "lanyard/pkg/config"
	admin "lanyard/server/internal/controllers/admin"
	games "lanyard/server/internal/controllers/gamelist"
	hc "lanyard/server/internal/controllers/health"
	"lanyard/server/internal/controllers/identity"
	partyTime "lanyard/server/internal/controllers/time"
	"lanyard/server/internal/middleware"
	"lanyard/server/internal/store"
	"log"
	"net/http"
	"strconv"
)

//go:embed lanyard
var lanyardFs embed.FS

func Start(appConfig *cfg.AppConfig) {
	log.Print("starting server on port ", appConfig.ServerPort)

	store.InitPartyConfig(appConfig)
	store.InitIdentities()

	mux := http.NewServeMux()

	mux.HandleFunc(hc.BasePath, hc.HealthCheck)
	mux.HandleFunc(partyTime.BasePath+"/start", partyTime.GetStartTime)
	mux.HandleFunc(partyTime.BasePath+"/departure", partyTime.AddDepartureTime)
	mux.HandleFunc(partyTime.BasePath+"/howmany", partyTime.GetHowManyAtTime)
	mux.HandleFunc(games.BasePath+"/resolve", games.ResolveGames)
	mux.HandleFunc(identity.BasePath+"/register", identity.CreateUsername)
	mux.HandleFunc(identity.BasePath+"/retrieve", identity.RetrieveUsername)

	// admin control endpoints - only localhost accepted
	mux.HandleFunc(admin.BasePath+"/options/update", admin.UpdateOption)
	mux.HandleFunc(admin.BasePath+"/users", admin.GetUsers)
	mux.HandleFunc(admin.BasePath+"/users/resolved", admin.GetUserResolutions)

	subFS, htmlPathErr := fs.Sub(lanyardFs, "lanyard")
	if htmlPathErr != nil {
		log.Fatal("failed to create paths for webpage:", htmlPathErr)
	}
	webPageHandler := http.StripPrefix("/lanyard/", http.FileServer(http.FS(subFS)))
	mux.Handle("/lanyard/", webPageHandler)

	loggedMux := middleware.LoggingMiddleware(mux)
	identityStampedMux := middleware.IdentityStampedMiddleware(loggedMux)

	err := http.ListenAndServe(":"+strconv.Itoa(appConfig.ServerPort), identityStampedMux)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("server closed")
	} else if err != nil {
		log.Fatal("server encountered error: ", err.Error())
	}
}
