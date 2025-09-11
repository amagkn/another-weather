package weather_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/amagkn/another-weather/internal/weather/dto"
	"github.com/amagkn/another-weather/pkg/common_error"
)

func (w *WeatherApi) GetWeatherByCity(lot, lat string) (dto.OpenmeteoWeather, error) {
	var weather dto.OpenmeteoWeather

	openmeteoURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current_weather=true", lat, lot)
	res, err := http.Get(openmeteoURL)
	if err != nil {
		return weather, common_error.WithPath("http.Get", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return weather, common_error.WithPath("io.ReadAll", err)
	}

	if err := json.Unmarshal(body, &weather); err != nil {
		return weather, common_error.WithPath("json.Unmarshal", err)
	}

	return weather, nil
}
