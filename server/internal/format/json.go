package format

type SimpleTimeFormat struct {
	Time int64 `json:"time"`
}

type HowManyResponse struct {
	Time      int64 `json:"time"`
	NumPeople int64 `json:"numPeople"`
}

type GameResolutionRequest struct {
	Games    map[string]int64 `json:"games"`
	Resolved map[string]bool  `json:"resolved"`
}

type IdentityFormat struct {
	Username string `json:"username"`
	Identity string `json:"lanyardId"`
}

type OptionUpdateRequest struct {
	OptionName  string `json:"option"`
	OptionValue any    `json:"value"`
}

type ListUsersResponse struct {
	Users map[string]string `json:"users"`
}

type ListUserResolutionsResponse struct {
	Resolutions map[string]map[string]bool `json:"resolutions"`
}
