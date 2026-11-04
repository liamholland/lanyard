package main

import (
	"fmt"
	cfg "lanyard/pkg/config"
	service "lanyard/pkg/lan_party/gamelist"
	opts "lanyard/pkg/lan_party/options"
	"log"

	"github.com/spf13/cobra"
)

var gameListCmd = &cobra.Command{
	Use:     opts.Options[opts.GameListPath],
	Short:   "interact with the game list",
	Long:    "interact with the game list. commands will check the user provided path first, followed by\nthe setting in the party configuration",
	Aliases: []string{"g"},
}

var viewGameListCmd = &cobra.Command{
	Use:   "show",
	Short: "show the games on a game list",
	Run: func(cmd *cobra.Command, args []string) {
		CheckVerboseMode(cmd)
		ViewGameList(GetAppConfig(), userGameListPathFlagValue)
	},
}

var userGameListPathFlagValue = ""

func init() {
	gameListCmd.PersistentFlags().StringVarP(&userGameListPathFlagValue, "path", "p", "", "set the path of the game list to use")

	gameListCmd.AddCommand(viewGameListCmd)
	rootCmd.AddCommand(gameListCmd)
}

func ViewGameList(appConfig *cfg.AppConfig, userPath string) {
	if userPath != "" {
		readFromUserPath(userPath)
		return
	}

	readFromSetting(appConfig)
}

func readFromUserPath(path string) {
	gl, err := service.ReadGameList(path)
	if err != nil {
		return
	}

	displayGameList(gl)
}

func readFromSetting(appConfig *cfg.AppConfig) {
	lpConfig := loadConfig(appConfig)
	if lpConfig == nil {
		fmt.Print("no path provided to read list from")
		return
	}

	optionValue := opts.GetOptionValue(lpConfig, opts.GameListPath)
	fp, ok := optionValue.(string)
	if !ok {
		log.Fatal("current option value is invalid: ", optionValue)
	}

	gl, err := service.ReadGameList(fp)
	if err != nil {
		return
	}

	displayGameList(gl)
}

func displayGameList(gl *service.GameList) {
	fmt.Println("games on the list:")
	for _, g := range *gl {
		fmt.Println("> ", g)
	}
}
