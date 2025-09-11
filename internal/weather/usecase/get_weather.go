package usecase

import (
	"github.com/amagkn/another-weather/internal/weather/dto"
	"github.com/amagkn/another-weather/internal/weather/entity"
	"github.com/amagkn/another-weather/pkg/common_error"
)

func (u *UseCase) GetWeather(city string) (dto.GetWeatherOutput, error) {
	var output dto.GetWeatherOutput

	cities, err := u.cityAPI.GetCityCoordinates(city)
	if err != nil {
		return output, common_error.WithPath("u.cityAPI.GetCityCoordinates", err)
	}

	if len(cities) == 0 {
		return output, entity.ErrCityNotFound
	}

	//_, _ := u.weatherAPI.GetWeatherByCity(cityWithCoordinates.Lot, cityWithCoordinates.Lat)

	return dto.GetWeatherOutput{}, nil
}
