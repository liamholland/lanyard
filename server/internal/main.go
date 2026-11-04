package internal

import (
	"errors"
	cfg "lanyard/pkg/config"
	"log"
	"net/http"
	"strconv"
)

func Start(appConfig *cfg.AppConfig) {
	log.Print("starting server on port ", appConfig.ServerPort)

	http.HandleFunc("/health/check", healthCheck)

	err := http.ListenAndServe(":"+strconv.Itoa(appConfig.ServerPort), nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("server closed")
	} else if err != nil {
		log.Fatal("server encountered error: ", err.Error())
	}
}
