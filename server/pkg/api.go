package pkg

import (
	cfg "lanyard/pkg/config"
	server "lanyard/server/internal"
)

func StartServer(appConfig *cfg.AppConfig) {
	server.Start(appConfig)
}
