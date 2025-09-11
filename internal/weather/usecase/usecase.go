package usecase

import (
	"github.com/amagkn/another-weather/internal/weather/dto"
)

type WeatherAPI interface {
	GetWeatherByCity(lot, lat string) (dto.OpenmeteoWeather, error)
}

type CityAPI interface {
	GetCityCoordinates(cityName string) ([]dto.NominatimCity, error)
}

type UseCase struct {
	weatherAPI WeatherAPI
	cityAPI    CityAPI
}

func New(w WeatherAPI, c CityAPI) *UseCase {
	return &UseCase{weatherAPI: w, cityAPI: c}
}
