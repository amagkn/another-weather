package city_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/amagkn/another-weather/internal/weather/dto"
	"github.com/amagkn/another-weather/pkg/common_error"
)

func (c *CityAPI) GetCityCoordinates(cityName string) ([]dto.NominatimCityOutput, error) {
	var cities []dto.NominatimCityOutput

	nominatimURL := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", cityName)
	res, err := http.Get(nominatimURL)
	if err != nil {
		return cities, common_error.WithPath("http.Get", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return cities, common_error.WithPath("io.ReadAll", err)
	}

	if err := json.Unmarshal(body, &cities); err != nil {
		return cities, common_error.WithPath("json.Unmarshal", err)
	}

	return cities, nil
}
