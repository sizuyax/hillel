package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log/slog"
	"project-auction/internal/adapters/postgres/adaperrors"
	"project-auction/internal/common/apperrors"
	"project-auction/internal/domain/entity"
)

//go:generate mockery --name=PGSellerRepository
type PGSellerRepository interface {
	InsertSeller(entity.Seller) (entity.Seller, error)
}

type pgSellerRepository struct {
	log      *slog.Logger
	database *sqlx.DB
}

func NewSellerRepository(log *slog.Logger, db *sqlx.DB) PGSellerRepository {
	return &pgSellerRepository{
		log:      log,
		database: db,
	}
}

func (sr pgSellerRepository) InsertSeller(seller entity.Seller) (entity.Seller, error) {
	q := `
	INSERT INTO profiles (email, password, type) 
	VALUES ($1, $2, $3)
	RETURNING id
	`

	if err := sr.database.QueryRowx(q, seller.Email, seller.Password, seller.Type).Scan(&seller.ID); err != nil {
		sr.log.Error("failed to execute request to db", slog.String("error", err.Error()))
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == adaperrors.UniqueViolation {
				sr.log.Error("conflicts", slog.String("error", err.Error()))
				return entity.Seller{}, apperrors.NewConflict("email", seller.Email)
			}
		}

		return entity.Seller{}, err
	}

	return seller, nil
}
