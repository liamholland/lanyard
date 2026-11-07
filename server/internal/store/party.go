package store

import (
	cfg "lanyard/pkg/config"
	lp "lanyard/pkg/lan_party"
	"lanyard/pkg/lan_party/service"
	"log"
	"os"
)

var lanParty lp.LanParty

func InitPartyConfig(appConfig *cfg.AppConfig) {
	lanParty = *loadConfig(appConfig)
}

func GetPartyConfig() *lp.LanParty {
	return &lanParty
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
