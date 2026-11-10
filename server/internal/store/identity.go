package store

import (
	"log"
	"strconv"
)

var userIdentities map[string]string
var usernames map[int64]string

func InitIdentities() {
	userIdentities = make(map[string]string)
	usernames = make(map[int64]string)
}

func EnterIdentity(id string, ip string) {
	userIdentities[id] = ip
}

func CheckIdentity(id string, ip string) int64 {
	if _, ok := userIdentities[id]; !ok {
		EnterIdentity(id, ip)
	}

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Print("failed to parse user id")
		return -1
	}

	return userId
}

func RegisterUsername(id int64, username string) {
	if _, ok := usernames[id]; ok {
		return
	}

	usernames[id] = username
}

func GetUsername(id int64) string {
	if name, ok := usernames[id]; !ok {
		return ""
	} else {
		return name
	}
}

func GetUsernames() map[int64]string {
	return usernames
}

func GetUsers() map[string]string {
	return userIdentities
}
