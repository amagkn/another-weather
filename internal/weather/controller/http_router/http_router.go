package http_router

import (
	ver1 "github.com/amagkn/another-weather/internal/weather/controller/http_router/ver1"
	"github.com/amagkn/another-weather/internal/weather/usecase"
	"github.com/go-chi/chi/v5"
)

func WeatherRouter(r *chi.Mux, uc *usecase.UseCase) {
	v1 := ver1.New(uc)

	r.Route("/api/v1/weather", func(r chi.Router) {
		r.Get("/{city}", v1.GetWeather)
	})
}
