package httpServer

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"project-auction/handlers"
)

func InitWebServer(log *slog.Logger) *echo.Echo {
	router := echo.New()

	handlers.NewHandler(log, handlers.Config{
		EchoRouter: router,
	})

	return router
}
