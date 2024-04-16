package handlers

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log/slog"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Handler struct {
	log *slog.Logger
}

type Config struct {
	EchoRouter *echo.Echo
}

func NewHandler(log *slog.Logger, cfg Config) {
	h := Handler{log: log}

	cfg.EchoRouter.GET("/items", h.GetItems)

	itemGroup := cfg.EchoRouter.Group("/item")
	itemGroup.GET("/:id", h.GetItemByID)
	itemGroup.POST("", h.CreateItem)
	itemGroup.PUT("/:id", h.UpdateItem)
	itemGroup.DELETE("/:id", h.DeleteItem)

	cfg.EchoRouter.GET("/swagger/*", echoSwagger.WrapHandler)
}
