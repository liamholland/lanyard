package internal

import (
	"embed"
	"errors"
	"io/fs"
	cfg "lanyard/pkg/config"
	"lanyard/pkg/store"
	admin "lanyard/server/internal/controllers/admin"
	games "lanyard/server/internal/controllers/gamelist"
	hc "lanyard/server/internal/controllers/health"
	"lanyard/server/internal/controllers/identity"
	partyTime "lanyard/server/internal/controllers/time"
	paths "lanyard/server/internal/format"
	"lanyard/server/internal/middleware"
	"log"
	"net/http"
	"strconv"
)

//go:embed lanyard
var lanyardFs embed.FS

func Start(appConfig *cfg.AppConfig) {
	log.Print("starting server on port ", appConfig.ServerPort)

	store.InitParty(appConfig)

	mux := http.NewServeMux()

	mux.HandleFunc(paths.HealthCheckPath, hc.HealthCheck)

	mux.HandleFunc(paths.TimeStartPath, partyTime.GetStartTime)
	mux.HandleFunc(paths.TimeDepPath, partyTime.AddDepartureTime)
	mux.HandleFunc(paths.TimeHowManyPath, partyTime.GetHowManyAtTime)
	mux.HandleFunc(paths.TimeGameIdealWhenPath, partyTime.GetWhenToPlayGame)

	mux.HandleFunc(paths.GamesResolvePath, games.ResolveGames)

	mux.HandleFunc(paths.IdentityRegisterPath, identity.CreateUsername)
	mux.HandleFunc(paths.IdentityRetrievePath, identity.RetrieveUsername)

	// admin control endpoints - only localhost accepted
	mux.HandleFunc(paths.AdminOptionsPath, admin.UpdateOption)
	mux.HandleFunc(paths.AdminUsersPath, admin.GetUsers)
	mux.HandleFunc(paths.AdminUsersResolvedPath, admin.GetUserResolutions)

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
