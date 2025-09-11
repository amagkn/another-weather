package ver1

import (
	"errors"
	"net/http"

	"github.com/amagkn/another-weather/internal/weather/entity"
	"github.com/amagkn/another-weather/pkg/common_error"
	"github.com/amagkn/another-weather/pkg/logger"
	"github.com/amagkn/another-weather/pkg/response"
	"github.com/go-chi/chi/v5"
)

func (h *Handlers) GetWeather(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")

	weather, err := h.uc.GetWeather(city)
	if err != nil {
		logger.Error(err, "h.uc.GetWeather")

		if errors.Is(err, entity.ErrCityNotFound) {
			response.Error(w, http.StatusBadRequest, response.ErrorPayload{Type: entity.ErrCityNotFound})
			return
		}

		response.Error(w, http.StatusBadRequest, response.ErrorPayload{Type: common_error.InternalServer})
		return
	}

	response.Success(w, http.StatusOK, weather)
}
