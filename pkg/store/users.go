package store

import "lanyard/pkg/lan_party/gamelist"

func ensureUserResolutions() *map[int64]gamelist.GameListResolution {
	if GetParty().UserResolutions == nil {
		GetParty().UserResolutions = make(map[int64]gamelist.GameListResolution)
	}

	return &GetParty().UserResolutions
}

func UpdateUserResolution(id int64, resolution map[string]bool) {
	ur := ensureUserResolutions()
	(*ur)[id] = resolution
}

func GetUserResolutions() *map[int64]gamelist.GameListResolution {
	return &GetParty().UserResolutions
}
