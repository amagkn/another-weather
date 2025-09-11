package weather_api

import "github.com/amagkn/another-weather/internal/weather/entity"

func (w *WeatherApi) GetWeatherByCity(lot, lat string) (entity.Weather, error) {
	return entity.Weather{}, nil
}
