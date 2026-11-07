package internal

import (
	"embed"
	"errors"
	"io/fs"
	cfg "lanyard/pkg/config"
	hc "lanyard/server/internal/controllers/health"
	partyTime "lanyard/server/internal/controllers/time"
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

	http.HandleFunc(hc.BasePath, hc.HealthCheck)
	http.HandleFunc(partyTime.BasePath+"/start", partyTime.GetStartTime)

	subFS, htmlPathErr := fs.Sub(lanyardFs, "lanyard")
	if htmlPathErr != nil {
		log.Fatal("failed to create paths for webpage:", htmlPathErr)
	}
	webPageHandler := http.StripPrefix("/lanyard/", http.FileServer(http.FS(subFS)))
	http.Handle("/lanyard/", webPageHandler)

	err := http.ListenAndServe(":"+strconv.Itoa(appConfig.ServerPort), nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("server closed")
	} else if err != nil {
		log.Fatal("server encountered error: ", err.Error())
	}
}
