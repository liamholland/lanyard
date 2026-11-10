package admin

const BasePath string = "/admin"

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
