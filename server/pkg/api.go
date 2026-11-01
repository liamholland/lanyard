package pkg

import (
	commontypes "lanyard/pkg/common_types"
	server "lanyard/server/internal"
)

func StartServer(appConfig *commontypes.AppConfig) {
	server.Start(appConfig)
}
