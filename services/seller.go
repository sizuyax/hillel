package services

import (
	"project-auction/models"
	"project-auction/repository"
)

type SellerService interface {
	CreateSeller(*models.Seller) (*models.Seller, error)
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

func (ss *sellerService) CreateSeller(seller *models.Seller) (*models.Seller, error) {
	seller, err := ss.SellerRepository.InsertSeller(seller)
	if err != nil {
		return &models.Seller{}, err
	}

	return seller, nil
}
