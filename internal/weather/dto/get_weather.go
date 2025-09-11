package dto

type GetWeatherOutput struct {
	City        string `json:"city"`
	Temperature string `json:"temperature"`
}
