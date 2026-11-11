package store

import (
	"log"
	"strconv"
)

func EnterIdentity(id string, ip string) {
	GetParty().Identities[id] = ip
}

func CheckIdentity(id string, ip string) int64 {
	if _, ok := GetParty().Identities[id]; !ok {
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
	if _, ok := GetParty().Usernames[id]; ok {
		return
	}

	GetParty().Usernames[id] = username
}

func GetUsername(id int64) string {
	if name, ok := GetParty().Usernames[id]; !ok {
		return ""
	} else {
		return name
	}
}

func GetUsernames() map[int64]string {
	return GetParty().Usernames
}

func GetUsers() map[string]string {
	return GetParty().Identities
}
