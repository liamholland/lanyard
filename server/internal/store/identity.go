package store

import (
	"log"
	"strconv"
)

var userIdentities map[string]string

func InitIdentities() {
	userIdentities = make(map[string]string)
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
