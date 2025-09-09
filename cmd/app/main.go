package main

import (
	"github.com/amagkn/another-weather/config"
	"github.com/amagkn/another-weather/internal/app"
	"github.com/amagkn/another-weather/pkg/logger"
	"github.com/amagkn/another-weather/pkg/validation"
)

func main() {
	c, err := config.New()
	if err != nil {
		logger.Fatal(err, "config.New")
	}

	logger.Init(c.Logger)
	validation.Init()

	err = app.Run(c)
	if err != nil {
		logger.Fatal(err, "app.Run")
	}

	logger.Info("App stopped!")
}
