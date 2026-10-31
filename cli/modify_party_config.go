package main

import (
	"fmt"
	types "lanyard/pkg/common_types"
	service "lanyard/pkg/lan_party_options"
	opts "lanyard/pkg/lan_party_options/options"
	"log"
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
		CreatePartyConfigFileWithDefaults(GetAppConfig())
	},
}

var getSettingValueCommand = &cobra.Command{
	Use:   "get",
	Short: "get and display the current value of a setting",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		GetCurrentSettingForOption(GetAppConfig(), args[0])
	},
}

func init() {
	partyConfigCommand.AddCommand(createConfigFileCommand)
	partyConfigCommand.AddCommand(getSettingValueCommand)
	rootCmd.AddCommand(partyConfigCommand)
}

func CreatePartyConfigFileWithDefaults(appConfig *types.AppConfig) {
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
	settings, loadError := service.LoadConfig(appConfig)
	if loadError != nil {
		log.Print("failed to load config, please try again")
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
