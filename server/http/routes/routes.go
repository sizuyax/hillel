package routes

import (
	"github.com/labstack/echo/v4"
	"hillel/server/http/handlers"
)

func SetupBookRoutes(e *echo.Echo) {
	e.GET("/books", handlers.GetBooks)

	e.POST("/newbooks", handlers.PostBooks)

	e.PUT("/updatebooks", handlers.PutBooks)

	e.DELETE("/deletebooks", handlers.DeleteBooks)
}
