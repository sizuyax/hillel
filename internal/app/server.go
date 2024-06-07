package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"log/slog"
	"project-auction/internal/adapters/postgres/repository"
	"project-auction/internal/controller/http/v1/handlers"
	"project-auction/internal/controller/http/v1/routes"
	"project-auction/internal/domain/services"
)

func InitWebServer(log *slog.Logger, db *sqlx.DB) *echo.Echo {
	router := echo.New()

	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(services.USConfig{
		UserRepository: userRepository,
	})

	sellerRepository := repository.NewSellerRepository(db)
	sellerService := services.NewSellerService(services.SSConfig{
		SellerRepository: sellerRepository,
	})

	itemRepository := repository.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)

	commentRepository := repository.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepository)

	handler := handlers.NewHandler(handlers.Config{
		EchoRouter:     router,
		Log:            log,
		UserService:    userService,
		SellerService:  sellerService,
		ItemService:    itemService,
		CommentService: commentService,
	})

	routes.SetupRoutes(handlers.Config{
		EchoRouter: router,
	}, handler)

	return router
}
