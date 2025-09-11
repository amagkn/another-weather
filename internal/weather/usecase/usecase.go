package usecase

import "github.com/amagkn/another-weather/internal/weather/entity"

type WeatherAPI interface {
	GetWeatherByCity(lot, lat string) (entity.Weather, error)
}

type CityAPI interface {
	GetCityCoordinates(cityName string) ([]entity.CityWithCoordinates, error)
}

type UseCase struct {
	weatherAPI WeatherAPI
	cityAPI    CityAPI
}

func New(w WeatherAPI, c CityAPI) *UseCase {
	return &UseCase{weatherAPI: w, cityAPI: c}
}
