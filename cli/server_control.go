package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	cfg "lanyard/pkg/config"
	gl "lanyard/pkg/lan_party/gamelist"
	server "lanyard/server/pkg"
	"log"
	"net/http"
	"os"
	"strings"

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

var serverStatusCmd = &cobra.Command{
	Use:   "up",
	Short: "check the status of the server",
	Run: func(cmd *cobra.Command, args []string) {
		CheckVerboseMode(cmd)
		IsServerUp()
	},
}

var updateServerSettingCmd = &cobra.Command{
	Use:   "update",
	Short: "update a setting on the running server",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		CheckVerboseMode(cmd)
		IsServerUp()
		UpdateServerSetting(args[0], args[1])
	},
}

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "interact with server user information",
}

var getUsersCmd = &cobra.Command{
	Use:   "list",
	Short: "list the users on the server",
	Long:  "lists users on the server. returns their username and the ip at which the server last recorded them",
	Run: func(cmd *cobra.Command, args []string) {
		CheckVerboseMode(cmd)
		IsServerUp()
		GetUsers()
	},
}

var getUserResolutionsCmd = &cobra.Command{
	Use:   "blame",
	Short: "display the status of users' resolved games",
	Long:  "display the status of each registered user and the games they have resolved with the server. lets you blame someone for not being prepared",
	Run: func(cmd *cobra.Command, args []string) {
		CheckVerboseMode(cmd)
		IsServerUp()
		GetUserResolutionStatus()
	},
}

var getGameResolutionStatus = &cobra.Command{
	Use:   "status",
	Short: "check the resolution status for a game",
	Long:  "check the resolution status for a game across all users on the server. takes the game name as an argument",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CheckVerboseMode(cmd)
		IsServerUp()
		GetResolutionStatusForGame(args[0])
	},
}

const serverLocation string = "http://localhost:1311"

type ListUsersResponse struct {
	Users map[string]string `json:"users"`
}

type ListUserResolutionsResponse struct {
	Resolutions map[string]gl.GameListResolution `json:"resolutions"`
}

func init() {
	serverCmd.PersistentFlags().IntVarP(&GetAppConfig().ServerPort, "port", "p", 1311, "set the port on which to run the server")
	serverCmd.PersistentFlags().StringVarP(&GetAppConfig().LanPartyDataPath, "config", "P", "party_config", "the path of the config to use")

	usersCmd.AddCommand(getUsersCmd)
	usersCmd.AddCommand(getUserResolutionsCmd)
	usersCmd.AddCommand(getGameResolutionStatus)

	serverCmd.AddCommand(startServerCmd)
	serverCmd.AddCommand(serverStatusCmd)
	serverCmd.AddCommand(updateServerSettingCmd)
	serverCmd.AddCommand(usersCmd)

	rootCmd.AddCommand(serverCmd)
}

func StartServer(appConfig *cfg.AppConfig) {
	server.StartServer(appConfig)
}

func IsServerUp() {
	request := serverLocation + "/health/check"
	serverRes, serverErr := http.Get(request)
	if serverErr != nil || serverRes == nil {
		fmt.Println("server is not running; start it with 'server start'")
		log.Fatal("server not running")
	}

	if serverRes.StatusCode != 200 {
		fmt.Println("server is not healthy, responded with ", serverRes.StatusCode, "; restart it")
		log.Fatal("server not healthy")
	}

	fmt.Println("server is running :)")

	resBody, readErr := io.ReadAll(serverRes.Body)
	if readErr != nil {
		log.Fatal("failed to read server response")
	}
	fmt.Println("server responded with: ", string(resBody))
}

func UpdateServerSetting(optionName string, optionValue string) {
	requestPath := serverLocation + "/admin/options/update"
	requestBody := map[string]string{
		"option": optionName,
		"value":  optionValue,
	}

	encodedBody, encodeErr := json.Marshal(requestBody)
	if encodeErr != nil {
		log.Fatal("failed to encode command options to json request")
	}

	bodyReader := bytes.NewReader(encodedBody)
	serverRes, serverErr := http.Post(requestPath, "application/json", bodyReader)
	if serverErr != nil {
		fmt.Println("failed to update option on server")
		log.Fatal("server encountered error while updating option: ", serverErr)
	}

	if serverRes.StatusCode != 200 {
		fmt.Println("failed to update option on server")
		log.Fatal("server responded with:", serverRes.StatusCode)
	}

	fmt.Println("option updated on server")
}

func GetUsers() {
	request := serverLocation + "/admin/users"
	serverRes, serverErr := http.Get(request)
	if serverErr != nil {
		log.Fatal("user list fetch from server failed: ", serverErr)
		return
	}

	if serverRes.StatusCode != 200 {
		fmt.Println("failed to get users on server")
		os.Exit(1)
	}

	defer serverRes.Body.Close()

	var res ListUsersResponse

	decodeErr := json.NewDecoder(serverRes.Body).Decode(&res)
	if decodeErr != nil {
		log.Fatal("server returned a malformed response")
	}

	fmt.Println("Current Users:")
	for u, ip := range res.Users {
		fmt.Printf("> %-20s%s\n", u, ip)
	}
}

func GetUserResolutionStatus() {
	request := serverLocation + "/admin/users/resolved"
	serverRes, serverErr := http.Get(request)
	if serverErr != nil {
		log.Fatal("user list fetch from server failed: ", serverErr)
		return
	}

	if serverRes.StatusCode != 200 {
		fmt.Println("failed to get users on server")
		os.Exit(1)
	}

	defer serverRes.Body.Close()

	var res ListUserResolutionsResponse

	decodeErr := json.NewDecoder(serverRes.Body).Decode(&res)
	if decodeErr != nil {
		log.Fatal("server returned a malformed response")
	}

	fmt.Println("Current User Game Status:")
	for u, r := range res.Resolutions {
		fmt.Printf("\n> %s:\n", u)
		if len(r) > 0 {
			displayGameResolution(&r)
		} else {
			fmt.Println("no resolution status!")
		}
	}
}

func GetResolutionStatusForGame(game string) {
	request := serverLocation + "/admin/users/resolved"
	serverRes, serverErr := http.Get(request)
	if serverErr != nil {
		log.Fatal("user list fetch from server failed: ", serverErr)
		return
	}

	if serverRes.StatusCode != 200 {
		fmt.Println("failed to get users on server")
		os.Exit(1)
	}

	defer serverRes.Body.Close()

	var res ListUserResolutionsResponse

	decodeErr := json.NewDecoder(serverRes.Body).Decode(&res)
	if decodeErr != nil {
		log.Fatal("server returned a malformed response")
	}

	var gameStatus = make(map[string]bool)
	var unknown = make([]string, 0)

	for u, r := range res.Resolutions {
		if uGameStatus, ok := r[game]; ok {
			gameStatus[u] = uGameStatus
		} else {
			unknown = append(unknown, u)
		}
	}

	if len(unknown) == len(res.Resolutions) {
		log.Print("failed to resolve any data on game ", game, " for any users")
		fmt.Println("no resolution data for '", game, "'. are you sure its spelled correctly?")
		fmt.Println("use 'gamelist show' to list the games on the game list")
	} else {
		fmt.Printf("result of game resolution for %s:\n", game)
		displayGameResolution((*gl.GameListResolution)(&gameStatus))
		if len(unknown) > 0 {
			fmt.Println("users with no resolution: ", strings.Join(unknown, ", "))
		}
	}
}
