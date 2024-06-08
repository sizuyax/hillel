package handlers

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"project-auction/internal/domain/services"
)

type Handler struct {
	log            *slog.Logger
	userService    services.UserService
	sellerService  services.SellerService
	itemService    services.ItemService
	commentService services.CommentService
	tokenService   services.TokenService
}

type Config struct {
	EchoRouter     *echo.Echo
	Log            *slog.Logger
	UserService    services.UserService
	SellerService  services.SellerService
	ItemService    services.ItemService
	CommentService services.CommentService
	TokenService   services.TokenService
}

func NewHandler(cfg Config) Handler {
	return Handler{
		log:            cfg.Log,
		userService:    cfg.UserService,
		sellerService:  cfg.SellerService,
		itemService:    cfg.ItemService,
		commentService: cfg.CommentService,
		tokenService:   cfg.TokenService,
	}
}
