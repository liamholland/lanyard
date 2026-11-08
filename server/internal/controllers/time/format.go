package time

const BasePath string = "/time"

type SimpleTimeFormat struct {
	Time int64 `json:"time"`
}

type HowManyResponse struct {
	Time      int64 `json:"time"`
	NumPeople int64 `json:"numPeople"`
}
