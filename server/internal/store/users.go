package store

type UserResolutions map[int64]map[string]bool

var userResolutions UserResolutions

func ensureUserResolutions() *UserResolutions {
	if userResolutions == nil {
		userResolutions = make(UserResolutions)
	}

	return &userResolutions
}

func UpdateUserResolution(id int64, resolution map[string]bool) {
	ur := ensureUserResolutions()
	(*ur)[id] = resolution
}

func GetUserResolutions() *UserResolutions {
	return &userResolutions
}
