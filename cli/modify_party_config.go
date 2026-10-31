package main

import (
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

func init() {
	partyConfigCommand.AddCommand(createConfigFileCommand)
	rootCmd.AddCommand(partyConfigCommand)
}

func CreatePartyConfigFileWithDefaults(appConfig *types.AppConfig) {
	createErr := service.CreateConfigFile(appConfig)
	if createErr != nil {
		log.Print("failed to create file, please try again. error: ", createErr.Error())
	}

	writeErr := service.UpdateConfig(appConfig, types.LanParty{
		Options: map[string]any{
			opts.Options[opts.StartDate]: time.Now(), // start date
		},
	})
	if writeErr != nil {
		log.Print("failed to write defaults to file, please try again. error: ", createErr.Error())
	}
}
