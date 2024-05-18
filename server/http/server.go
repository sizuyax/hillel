package httpServer

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"log/slog"
	"project-auction/handlers"
	"project-auction/repository"
	"project-auction/services"
)

func InitWebServer(log *slog.Logger, db *sqlx.DB) *echo.Echo {
	router := echo.New()

	userRepository := repository.NewUserRepository(log, db)
	userService := services.NewUserService(services.USConfig{
		UserRepository: userRepository,
	})

	sellerRepository := repository.NewSellerRepository(log, db)
	sellerService := services.NewSellerService(services.SSConfig{
		SellerRepository: sellerRepository,
	})

	itemRepository := repository.NewItemRepository(log, db)
	itemService := services.NewItemService(services.ISConfig{
		ItemRepository: itemRepository,
	})

	handler := handlers.NewHandler(handlers.Config{
		EchoRouter:    router,
		Log:           log,
		UserService:   userService,
		SellerService: sellerService,
		ItemService:   itemService,
	})

	handlers.SetupRoutes(handlers.Config{
		EchoRouter: router,
	}, handler)

	return router
}
