package app

import (
	"github.com/amagkn/another-weather/internal/weather/controller/http_router"
	"github.com/amagkn/another-weather/internal/weather/usecase"
)

func WeatherDomain(d Dependences) {
	weatherUseCase := usecase.New()

	http_router.WeatherRouter(d.RouterHTTP, weatherUseCase)
}
