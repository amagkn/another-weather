package ver1

import (
	"fmt"
	"net/http"

	"github.com/amagkn/another-weather/pkg/response"
	"github.com/go-chi/chi/v5"
)

func (h *Handlers) GetWeather(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")

	response.Success(w, http.StatusOK, fmt.Sprintf("%s cool!", city))
}
