package services

import (
	"log/slog"
	"project-auction/internal/adapters/postgres/repository"
	"project-auction/internal/domain/entity"
)

type SellerService interface {
	CreateSeller(entity.Seller) (entity.Seller, error)
}

type sellerService struct {
	log              *slog.Logger
	SellerRepository repository.PGSellerRepository
}

func NewSellerService(log *slog.Logger, sellerRepository repository.PGSellerRepository) SellerService {
	return &sellerService{
		log:              log,
		SellerRepository: sellerRepository,
	}
}

func (ss *sellerService) CreateSeller(seller entity.Seller) (entity.Seller, error) {
	seller, err := ss.SellerRepository.InsertSeller(seller)
	if err != nil {
		ss.log.Error("failed to insert seller", slog.String("error", err.Error()))
		return entity.Seller{}, err
	}

	return seller, nil
}
