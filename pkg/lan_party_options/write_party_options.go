package lanpartyoptions

import (
	types "lanyard/pkg/common_types"
	"log"
	"os"
)

func CreateConfigFile(appConfig *types.AppConfig) error {
	_, err := os.OpenFile(appConfig.LanPartyDataPath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("failed to create file at path '%s'", appConfig.LanPartyDataPath)
		return err
	}
	return nil
}
