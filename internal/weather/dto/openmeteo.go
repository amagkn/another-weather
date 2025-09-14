package dto

type OpenmeteoWeatherOutput struct {
	CurrentWeather struct {
		Time          string  `json:"time"`
		Interval      int     `json:"interval"`
		Temperature   float32 `json:"temperature"`
		Windspeed     float64 `json:"windspeed"`
		Winddirection int     `json:"winddirection"`
		IsDay         int     `json:"is_day"`
		Weathercode   int     `json:"weathercode"`
	} `json:"current_weather"`
}
