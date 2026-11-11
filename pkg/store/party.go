package store

import (
	cfg "lanyard/pkg/config"
	lp "lanyard/pkg/lan_party"
	"lanyard/pkg/lan_party/service"
	"log"
	"os"
	"time"
)

var lanParty lp.LanParty

func InitParty(appConfig *cfg.AppConfig) {
	lanParty = *loadConfig(appConfig)
	lanParty.DepartureTimes = make(map[int64]time.Time)
	lanParty.Identities = make(map[string]string)
	lanParty.Usernames = make(map[int64]string)
}

func GetParty() *lp.LanParty {
	return &lanParty
}

func SetParty(newSettings *lp.LanParty) {
	lanParty = *newSettings
}

func loadConfig(appConfig *cfg.AppConfig) *lp.LanParty {
	settings, loadError := service.LoadConfig(appConfig)
	if os.IsNotExist(loadError) {
		log.Fatal("config does not exist. create one with 'lanyard party create'")
	} else if loadError != nil {
		log.Fatal("failed to load config, please try again: ", loadError.Error())
	}

	return settings
}
