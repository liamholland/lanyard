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

var resolveGameListCmd = &cobra.Command{
	Use:   "resolve",
	Short: "resolve the games list across game directories",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CheckVerboseMode(cmd)
		ResolveGameList(GetAppConfig(), userGameListPathFlagValue, args)
	},
}

var userGameListPathFlagValue = ""

func init() {
	gameListCmd.PersistentFlags().StringVarP(&userGameListPathFlagValue, "path", "p", "", "set the path of the game list to use")

	gameListCmd.AddCommand(viewGameListCmd)
	gameListCmd.AddCommand(resolveGameListCmd)
	rootCmd.AddCommand(gameListCmd)
}

func ViewGameList(appConfig *cfg.AppConfig, userPath string) {
	if userPath != "" {
		readFromPath(userPath)
		return
	}

	readFromSetting(appConfig)
}

func ResolveGameList(appConfig *cfg.AppConfig, userPath string, gamePaths []string) {
	log.Print("resolving games")
	if userPath != "" {
		resolveFromPath(userPath, gamePaths)
		return
	}

	resolveFromSetting(appConfig, gamePaths)
}

func readFromPath(path string) {
	gl, err := service.ReadGameList(path)
	if err != nil {
		return
	}

	displayGameList(gl)
}

func readFromSetting(appConfig *cfg.AppConfig) {
	fp := getGameListPathFromSetting(appConfig)
	readFromPath(fp)
}

func displayGameList(gl *service.GameList) {
	fmt.Println("games on the list:")
	for _, g := range *gl {
		fmt.Println("> ", g)
	}
}

func resolveFromPath(listPath string, gamePaths []string) {
	gl, err := service.ReadGameList(listPath)
	if err != nil {
		log.Print("error while read game list")
		return
	}

	resolvedGames := service.ResolveGameListFromPaths(gl, gamePaths)
	if resolvedGames == nil {
		log.Print("error occurred during game resolution")
		return
	}

	displayGameResolution(resolvedGames)
}

func resolveFromSetting(config *cfg.AppConfig, gamePaths []string) {
	fp := getGameListPathFromSetting(config)
	resolveFromPath(fp, gamePaths)
}

func displayGameResolution(gameRes *service.GameListResolution) {
	fmt.Println("result of game resolution:")
	for g, r := range *gameRes {
		if g == "" {
			continue
		}

		fmt.Printf("> %-30s%s\n", g, func() string {
			if r {
				return "ok!"
			}
			return "missing"
		}())
	}
}

func getGameListPathFromSetting(appConfig *cfg.AppConfig) string {
	lpConfig := loadConfig(appConfig)
	if lpConfig == nil {
		fmt.Print("no path provided to read list from")
		return ""
	}

	optionValue := opts.GetOptionValue(lpConfig, opts.GameListPath)
	fp, ok := optionValue.(string)
	if !ok {
		log.Fatal("current option value is invalid: ", optionValue)
	}

	return fp
}
