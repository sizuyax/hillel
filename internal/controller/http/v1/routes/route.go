package routes

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	"project-auction/internal/controller/http/v1/handlers"
	"project-auction/internal/domain/services/middlewares"
)

func SetupRoutes(cfg handlers.Config, handler handlers.Handler) {
	itemGroupForSeller := cfg.EchoRouter.Group("/items", middlewares.ParseAccessSellerToken(cfg.TokenService))
	itemGroupForSeller.POST("", handler.CreateItem)
	itemGroupForSeller.PUT("/:id", handler.UpdateItem)
	itemGroupForSeller.DELETE("/:id", handler.DeleteItemByID)

	itemGroupForUser := cfg.EchoRouter.Group("/items")
	itemGroupForUser.GET("", handler.GetItems)
	itemGroupForUser.GET("/:id", handler.GetItemByID)
	itemGroupForUser.POST("/:id/comments", handler.CreateComment, middlewares.ParseAccessUserToken(cfg.TokenService))
	itemGroupForUser.POST("/:id/bids", handler.CreateBid, middlewares.ParseAccessUserToken(cfg.TokenService))

	cfg.EchoRouter.GET("/ws/:id", handler.WebSocket)

	sellerGroup := cfg.EchoRouter.Group("/sellers")
	sellerGroup.POST("", handler.RegisterSeller)

	userGroup := cfg.EchoRouter.Group("/users")
	userGroup.POST("", handler.RegisterUser)

	authGroup := cfg.EchoRouter.Group("/auth")
	authGroup.POST("/tokens", handler.RefreshTokens)

	cfg.EchoRouter.GET("/swagger/*", echoSwagger.WrapHandler)
}
