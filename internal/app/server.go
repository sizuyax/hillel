package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"log/slog"
	"project-auction/internal/adapters/repository/postgres"
	"project-auction/internal/controller/http/v1/handlers"
	"project-auction/internal/controller/http/v1/routes"
	"project-auction/internal/domain/services"
)

func InitWebServer(log *slog.Logger, db *sqlx.DB) *echo.Echo {
	router := echo.New()

	userRepository := postgres.NewUserRepository(log, db)
	userService := services.NewUserService(services.USConfig{
		UserRepository: userRepository,
	})

	sellerRepository := postgres.NewSellerRepository(log, db)
	sellerService := services.NewSellerService(services.SSConfig{
		SellerRepository: sellerRepository,
	})

	itemRepository := postgres.NewItemRepository(log, db)
	itemService := services.NewItemService(itemRepository)

	handler := handlers.NewHandler(handlers.Config{
		EchoRouter:    router,
		Log:           log,
		UserService:   userService,
		SellerService: sellerService,
		ItemService:   itemService,
	})

	routes.SetupRoutes(handlers.Config{
		EchoRouter: router,
	}, handler)

	return router
}
