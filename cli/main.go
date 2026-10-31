package main

import (
	"fmt"
	types "lanyard/pkg/common_types"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "ly",
	Aliases: []string{"lanyard"},
	Short:   "cli tool for managing lanyard server",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var appConfig types.AppConfig = types.AppConfig{
	LanPartyDataPath: "./party_config",
}

func main() {
	log.Default().SetOutput(os.Stderr)
	runCli()
}

func runCli() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred while running lanyard")
		os.Exit(1)
	}
}

func GetAppConfig() *types.AppConfig {
	return &appConfig
}
