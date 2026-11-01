package main

import (
	"fmt"
	types "lanyard/pkg/common_types"
	service "lanyard/pkg/lan_party_options"
	opts "lanyard/pkg/lan_party_options/options"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var partyConfigCommand = &cobra.Command{
	Use:     "party",
	Short:   "modify the party settings",
	Aliases: []string{"p"},
}

var createConfigFileCommand = &cobra.Command{
	Use:   "create",
	Short: "create the settings file for the party",
	Run: func(cmd *cobra.Command, args []string) {
		CheckVerboseMode(cmd)
		CreatePartyConfigFileWithDefaults(GetAppConfig())
	},
}

var getSettingValueCommand = &cobra.Command{
	Use:   "get",
	Short: "get and display the current value of a setting",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CheckVerboseMode(cmd)
		GetCurrentSettingForOption(GetAppConfig(), args[0])
	},
}

var updateSettingValueCommand = &cobra.Command{
	Use:   "set",
	Short: "update the value of a setting",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		CheckVerboseMode(cmd)
		UpdateCurrentSettingForOption(GetAppConfig(), args[0], args[1])
	},
}

func init() {
	partyConfigCommand.PersistentFlags().StringVarP(&GetAppConfig().LanPartyDataPath, "path", "p", "party_config", "the path of the config to use")

	partyConfigCommand.AddCommand(createConfigFileCommand)
	partyConfigCommand.AddCommand(getSettingValueCommand)
	partyConfigCommand.AddCommand(updateSettingValueCommand)

	rootCmd.AddCommand(partyConfigCommand)
}

func CreatePartyConfigFileWithDefaults(appConfig *types.AppConfig) {
	log.Print("creating file at: ", appConfig.LanPartyDataPath)

	createErr := service.CreateConfigFile(appConfig)
	if createErr != nil {
		log.Print("failed to create file, please try again. error: ", createErr.Error())
		return
	}

	writeErr := service.UpdateConfig(appConfig, types.LanParty{
		Options: map[string]any{
			opts.Options[opts.StartDate]: time.Now(), // start date
		},
	})
	if writeErr != nil {
		log.Print("failed to write defaults to file, please try again. error: ", writeErr.Error())
	}
}

func GetCurrentSettingForOption(appConfig *types.AppConfig, arg string) {
	log.Print("getting value of setting '", arg, "' in config at: ", appConfig.LanPartyDataPath)

	settings := loadConfig(appConfig)
	if settings == nil {
		return
	}

	if !opts.IsValidOptionName(arg) {
		log.Print("invalid setting name: '", arg, "'")
		return
	}

	v := opts.GetOptionValueFromName(settings, arg)
	if v == nil {
		log.Print("got an empty value")
		return
	}

	fmt.Println("Setting '", arg, "' is currently set as: ", v)
}

func UpdateCurrentSettingForOption(appConfig *types.AppConfig, arg string, v string) {
	log.Print("updating the value of setting '", arg, "' to '", v, "'")

	ok := opts.IsValidOptionName(arg)
	if !ok {
		fmt.Printf("'%s' is not a valid setting\n", arg)
		return
	}

	option := opts.GetOptionFromName(arg)
	if option == nil {
		log.Print("failed to fetch option: ", arg)
		return
	}

	settings := loadConfig(appConfig)
	if settings == nil {
		return
	}

	newSettings := opts.Copy(settings)
	validErr := option.OnRead(newSettings, v)
	if validErr != nil {
		fmt.Print("'", v, "' is not a valid value for setting '", arg, "'")
		log.Print(validErr.Error())
		return
	}

	updateErr := service.UpdateConfig(appConfig, *newSettings)
	if updateErr != nil {
		log.Print("update via cli failed")
		return
	}

	fmt.Println("Set value of '", arg, "' to '", v, "'")
}

func loadConfig(appConfig *types.AppConfig) *types.LanParty {

	settings, loadError := service.LoadConfig(appConfig)
	if os.IsNotExist(loadError) {
		fmt.Println("config does not exist. crete one with 'lanyard party create'")
		return nil
	} else if loadError != nil {
		log.Print("failed to load config, please try again: ", loadError.Error())
		return nil
	}

	return settings
}
