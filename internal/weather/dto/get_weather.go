package dto

type GetWeatherOutput struct {
	City        string  `json:"city"`
	Temperature float32 `json:"temperature"`
}
