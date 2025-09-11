package app

import (
	"github.com/amagkn/another-weather/internal/weather/adapters/city_api"
	"github.com/amagkn/another-weather/internal/weather/adapters/weather_api"
	"github.com/amagkn/another-weather/internal/weather/controller/http_router"
	"github.com/amagkn/another-weather/internal/weather/usecase"
)

func WeatherDomain(d Dependences) {
	weatherUseCase := usecase.New(weather_api.New(), city_api.New())

	http_router.WeatherRouter(d.RouterHTTP, weatherUseCase)
}
