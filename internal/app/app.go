package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/amagkn/another-weather/config"
	"github.com/amagkn/another-weather/pkg/http_server"
	"github.com/amagkn/another-weather/pkg/logger"
	"github.com/amagkn/another-weather/pkg/router"
	"github.com/go-chi/chi/v5"
)

type Dependences struct {
	RouterHTTP *chi.Mux
}

func Run(c config.Config) (err error) {
	var deps Dependences

	deps.RouterHTTP = router.New()

	WeatherDomain(deps)

	httpServer := http_server.New(deps.RouterHTTP, c.HTTP.Port)
	defer httpServer.Close()

	waiting(httpServer)

	return nil
}

func waiting(httpServer *http_server.Server) {
	logger.Info("App started")

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

	select {
	case i := <-wait:
		logger.Info("App got signal: " + i.String())
	case err := <-httpServer.Notify():
		logger.Error(err, "App got notify: httpServer.Notify")
	}

	logger.Info("App is stopping")
}
