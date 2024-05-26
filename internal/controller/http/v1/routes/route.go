package routes

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	"project-auction/internal/controller/http/v1/handlers"
	"project-auction/internal/domain/services/middlewares"
)

func SetupRoutes(cfg handlers.Config, handler handlers.Handler) {
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

	authGroup := cfg.EchoRouter.Group("/auth")
	authGroup.POST("/tokens", handler.RefreshTokens)

	cfg.EchoRouter.GET("/swagger/*", echoSwagger.WrapHandler)
}
