package handlers

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"project-auction/internal/adapters/storage"
	"project-auction/internal/domain/services"
)

type Handler struct {
	log            *slog.Logger
	userService    services.UserService
	sellerService  services.SellerService
	itemService    services.ItemService
	commentService services.CommentService
	tokenService   services.TokenService
	bidService     services.BidService
	minioStorage   storage.MinioStorage
}

type Config struct {
	EchoRouter     *echo.Echo
	Log            *slog.Logger
	UserService    services.UserService
	SellerService  services.SellerService
	ItemService    services.ItemService
	CommentService services.CommentService
	TokenService   services.TokenService
	BidService     services.BidService
	MinioStorage   storage.MinioStorage
}

func NewHandler(cfg Config) Handler {
	return Handler{
		log:            cfg.Log,
		userService:    cfg.UserService,
		sellerService:  cfg.SellerService,
		itemService:    cfg.ItemService,
		commentService: cfg.CommentService,
		tokenService:   cfg.TokenService,
		bidService:     cfg.BidService,
		minioStorage:   cfg.MinioStorage,
	}
}
