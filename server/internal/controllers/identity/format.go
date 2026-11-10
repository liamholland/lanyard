package identity

const BasePath string = "/identity"

type IdentityFormat struct {
	Username string `json:"username"`
	Identity string `json:"lanyardId"`
}
