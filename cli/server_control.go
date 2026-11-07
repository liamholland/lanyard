package main

import (
	cfg "lanyard/pkg/config"
	server "lanyard/server/pkg"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "control the lanyard server",
	Aliases: []string{"s"},
}

var startServerCmd = &cobra.Command{
	Use:   "start",
	Short: "start the lanyard server",
	Run: func(cmd *cobra.Command, args []string) {
		CheckVerboseMode(cmd)
		StartServer(GetAppConfig())
	},
}

func init() {
	serverCmd.PersistentFlags().IntVarP(&GetAppConfig().ServerPort, "port", "p", 1311, "set the port on which to run the server")
	serverCmd.PersistentFlags().StringVarP(&GetAppConfig().LanPartyDataPath, "config", "P", "party_config", "the path of the config to use")

	serverCmd.AddCommand(startServerCmd)
	rootCmd.AddCommand(serverCmd)
}

func StartServer(appConfig *cfg.AppConfig) {
	server.StartServer(appConfig)
}
