package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log/slog"
	"project-auction/apperrors"
	"project-auction/models"
)

type PGSellerRepository interface {
	InsertSeller(*models.Seller) (*models.Seller, error)
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

func (sr *pgSellerRepository) InsertSeller(seller *models.Seller) (*models.Seller, error) {
	const op = "repository.InsertSeller"

	q := `
	INSERT INTO sellers (email, password) 
	VALUES ($1, $2)
	RETURNING id
	`

	if err := sr.database.QueryRowx(q, seller.Email, seller.Password).Scan(&seller.ID); err != nil {
		sr.log.Error(fmt.Sprintf("%s: %v", op, err))

		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return &models.Seller{}, apperrors.NewConflict("email", seller.Email)
			}
		}

		return &models.Seller{}, fmt.Errorf("%s:%v", op, err)
	}

	return seller, nil
}
