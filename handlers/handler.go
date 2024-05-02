package handlers

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log/slog"
	"project-auction/services"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Handler struct {
	Log         *slog.Logger
	UserService services.UserService
	ItemService services.ItemService
}

type Config struct {
	EchoRouter  *echo.Echo
	Log         *slog.Logger
	UserService services.UserService
	ItemService services.ItemService
}

func NewHandler(cfg Config) Handler {
	h := Handler{
		Log:         cfg.Log,
		UserService: cfg.UserService,
		ItemService: cfg.ItemService,
	}
	return h
}

func SetupRoutes(cfg Config, handler Handler) {
	itemGroup := cfg.EchoRouter.Group("/items")
	itemGroup.GET("", handler.GetItems)
	itemGroup.GET("/:id", handler.GetItemByID)
	itemGroup.POST("", handler.CreateItem)
	itemGroup.PUT("/:id", handler.UpdateItem)
	itemGroup.DELETE("/:id", handler.DeleteItemByID)

	userGroup := cfg.EchoRouter.Group("/users")
	userGroup.POST("", handler.RegisterUser)

	cfg.EchoRouter.GET("/swagger/*", echoSwagger.WrapHandler)
}
