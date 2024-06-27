package services

import (
	"log/slog"
	"project-auction/internal/adapters/postgres/repository"
	"project-auction/internal/domain/entity"
)

type BidService interface {
	Create(bid entity.Bid) (entity.Bid, error)
}

type bidService struct {
	log           *slog.Logger
	bidRepository repository.PGBidRepository
}

func NewBidService(log *slog.Logger, bidRepository repository.PGBidRepository) BidService {
	return &bidService{
		log:           log,
		bidRepository: bidRepository,
	}
}

func (bs bidService) Create(bid entity.Bid) (entity.Bid, error) {
	expectBid, err := bs.bidRepository.Insert(bid)
	if err != nil {
		bs.log.Error("failed to insert bid", slog.String("error", err.Error()))
		return entity.Bid{}, err
	}

	return expectBid, nil
}
