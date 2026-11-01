package main

import (
	"fmt"
	"io"
	types "lanyard/pkg/common_types"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "ly",
	Aliases: []string{"lanyard"},
	Short:   "cli tool for managing lanyard server",
	Run:     func(cmd *cobra.Command, args []string) {},
}

var appConfig types.AppConfig = types.AppConfig{
	LanPartyDataPath: "./party_config",
	VerboseMode:      false,
	ServerPort:       1311,
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&appConfig.VerboseMode, "verbose", "v", false, "configure the application to run in verbose mode")
}

func main() {
	log.Default().SetOutput(io.Discard)
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

func CheckVerboseMode(cmd *cobra.Command) {
	if appConfig.VerboseMode {
		log.Default().SetOutput(os.Stdout)
	}
}
