package options

const BasePath string = "/options"

type OptionUpdateRequest struct {
	OptionName  string `json:"option"`
	OptionValue any    `json:"value"`
}
