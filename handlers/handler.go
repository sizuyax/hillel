package handlers

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log/slog"
	"project-auction/services"
	"project-auction/services/middlewares"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Handler struct {
	Log           *slog.Logger
	UserService   services.UserService
	SellerService services.SellerService
	ItemService   services.ItemService
}

type Config struct {
	EchoRouter    *echo.Echo
	Log           *slog.Logger
	UserService   services.UserService
	SellerService services.SellerService
	ItemService   services.ItemService
}

func NewHandler(cfg Config) Handler {
	h := Handler{
		Log:           cfg.Log,
		UserService:   cfg.UserService,
		SellerService: cfg.SellerService,
		ItemService:   cfg.ItemService,
	}
	return h
}

func SetupRoutes(cfg Config, handler Handler) {
	itemGroupWithToken := cfg.EchoRouter.Group("/items", middlewares.ParseAccessToken)
	itemGroupWithToken.POST("", handler.CreateItem)
	itemGroupWithToken.PUT("/:id", handler.UpdateItem)
	itemGroupWithToken.DELETE("/:id", handler.DeleteItemByID)

	itemGroup := cfg.EchoRouter.Group("/items")
	itemGroup.GET("", handler.GetItems)
	itemGroup.GET("/:id", handler.GetItemByID)

	sellerGroup := cfg.EchoRouter.Group("/sellers")
	sellerGroup.POST("", handler.RegisterSeller)

	userGroup := cfg.EchoRouter.Group("/users")
	userGroup.POST("", handler.RegisterUser)

	cfg.EchoRouter.GET("/swagger/*", echoSwagger.WrapHandler)
}
