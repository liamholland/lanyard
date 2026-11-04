package service

import (
	cfg "lanyard/pkg/config"
	lp "lanyard/pkg/lan_party"
	opts "lanyard/pkg/lan_party/options"
	"log"
	"os"
)

const updateTmpConfigExtension string = ".old"

func CreateConfigFile(appConfig *cfg.AppConfig) error {
	_, err := os.OpenFile(appConfig.LanPartyDataPath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("failed to create file at path '%s'", appConfig.LanPartyDataPath)
		return err
	}
	return nil
}

func UpdateConfig(appConfig *cfg.AppConfig, partyConfig lp.LanParty) error {
	reanameErr := os.Rename(appConfig.LanPartyDataPath, appConfig.LanPartyDataPath+updateTmpConfigExtension)
	if reanameErr != nil {
		log.Print("failed to rename existing config: ", reanameErr.Error())
		return reanameErr
	}

	updateErr := update(appConfig, &partyConfig)
	if updateErr != nil {
		log.Print("update failed; restoring file")

		restoreErr := os.Rename(appConfig.LanPartyDataPath+updateTmpConfigExtension, appConfig.LanPartyDataPath)
		if restoreErr != nil {
			log.Panic("failed to restore options! application is likely in an unstable state!")
			return restoreErr
		}

		return updateErr
	}

	os.Remove(appConfig.LanPartyDataPath + updateTmpConfigExtension)
	return nil
}

func update(appConfig *cfg.AppConfig, partyConfig *lp.LanParty) error {
	createErr := CreateConfigFile(appConfig)
	if createErr != nil {
		return createErr
	}

	writeErr := writeOptions(appConfig.LanPartyDataPath, partyConfig)
	if writeErr != nil {
		log.Print("failed to write options to file: ", writeErr.Error())
		return writeErr
	}

	return nil
}

func writeOptions(filePath string, partyConfig *lp.LanParty) error {
	file, openErr := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if openErr != nil {
		log.Print("failed to open file for writing options")
		return openErr
	}
	defer file.Close()

	for option := range opts.Options {
		err := writeOption(file, option, partyConfig.Options[opts.Options[option]])
		if err != nil {
			log.Print("failed to write option '", opts.Options[option], "': ", err.Error())
			return err
		}
	}
	return nil
}

func writeOption(file *os.File, option opts.ValidLanPartyOptionName, value any) error {
	if value == nil {
		log.Print("no value for option '", opts.Options[option], "'. skipping...")
		return nil
	}

	o := opts.GetOption(option)
	inFileOption, formatErr := o.OnWrite(value, o)
	if formatErr != nil {
		return formatErr
	}

	_, writeErr := file.WriteString(inFileOption + "\n")
	if writeErr != nil {
		return writeErr
	}

	return nil
}
