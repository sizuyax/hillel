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

	userRepository := repository.NewUserRepository(log, db)
	userService := services.NewUserService(log, userRepository)

	sellerRepository := repository.NewSellerRepository(log, db)
	sellerService := services.NewSellerService(log, sellerRepository)

	itemRepository := repository.NewItemRepository(log, db)
	itemService := services.NewItemService(log, itemRepository)

	commentRepository := repository.NewCommentRepository(log, db)
	commentService := services.NewCommentService(log, commentRepository)

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
