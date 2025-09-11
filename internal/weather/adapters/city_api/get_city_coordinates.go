package city_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/amagkn/another-weather/internal/weather/dto"
	"github.com/amagkn/another-weather/internal/weather/entity"
	"github.com/amagkn/another-weather/pkg/common_error"
)

func (c *CityAPI) GetCityCoordinates(cityName string) ([]entity.CityWithCoordinates, error) {
	var citiesWithCoordinates []entity.CityWithCoordinates

	nominatimURL := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", cityName)
	res, err := http.Get(nominatimURL)
	if err != nil {
		return citiesWithCoordinates, common_error.WithPath("http.Get", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return citiesWithCoordinates, common_error.WithPath("io.ReadAll", err)
	}

	var cities []dto.GetCityCoordinates
	if err := json.Unmarshal(body, &cities); err != nil {
		return citiesWithCoordinates, common_error.WithPath("json.Unmarshal", err)
	}

	for _, city := range cities {
		citiesWithCoordinates = append(citiesWithCoordinates, entity.CityWithCoordinates{
			Name: city.Name,
			Lot:  city.Lot,
			Lat:  city.Lat,
		})
	}

	return citiesWithCoordinates, nil
}
