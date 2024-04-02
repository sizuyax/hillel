package handlers

import (
	"github.com/labstack/echo/v4"
	"hillel/services"
)

func AllHandlers(e *echo.Echo) {
	e.GET("/books", services.GetBooks)

	e.POST("/newbooks", services.PostBooks)

	e.PUT("/updatebooks", services.PutBooks)

	e.DELETE("/deletebooks", services.DeleteBooks)
}
