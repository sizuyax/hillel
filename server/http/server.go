package httpServer

import (
	"github.com/labstack/echo/v4"
	"project-auction/server/http/routes"
)

func InitWebServer() *echo.Echo {
	e := echo.New()

	routes.SetupRoutes(e)

	return e
}
