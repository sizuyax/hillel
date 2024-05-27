package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"project-auction/internal/common/apperrors"
	"project-auction/internal/domain/entity"
)

//go:generate mockery --name=PGSellerRepository
type PGSellerRepository interface {
	InsertSeller(entity.Seller) (entity.Seller, error)
}

type pgSellerRepository struct {
	database *sqlx.DB
}

func NewSellerRepository(db *sqlx.DB) PGSellerRepository {
	return &pgSellerRepository{
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
		// TODO
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return entity.Seller{}, apperrors.NewConflict("email", seller.Email)
			}
		}

		return entity.Seller{}, err
	}

	return seller, nil
}
