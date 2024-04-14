package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"project-auction/server/http/handlers"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/items", handlers.GetItems)

	e.GET("/item/{id}", handlers.GetItemByID)

	e.POST("/create-item", handlers.CreateItem)

	e.PUT("/update-item", handlers.UpdateItem)

	e.DELETE("/delete-item", handlers.DeleteItem)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
