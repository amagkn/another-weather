package ver1

import (
	"github.com/amagkn/another-weather/internal/weather/usecase"
)

type Handlers struct {
	uc *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handlers {
	return &Handlers{uc: uc}
}
