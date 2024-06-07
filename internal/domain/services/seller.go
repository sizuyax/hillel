package services

import (
	"project-auction/internal/adapters/postgres/repository"
	"project-auction/internal/domain/entity"
)

type SellerService interface {
	CreateSeller(entity.Seller) (entity.Seller, error)
}

type sellerService struct {
	SellerRepository repository.PGSellerRepository
}

type SSConfig struct {
	SellerRepository repository.PGSellerRepository
}

func NewSellerService(cfg SSConfig) SellerService {
	return &sellerService{
		SellerRepository: cfg.SellerRepository,
	}
}

func (ss *sellerService) CreateSeller(seller entity.Seller) (entity.Seller, error) {
	seller, err := ss.SellerRepository.InsertSeller(seller)
	if err != nil {
		return entity.Seller{}, err
	}

	return seller, nil
}
