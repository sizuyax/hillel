package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log/slog"
	"project-auction/internal/adapters/postgres/adaperrors"
	"project-auction/internal/common/apperrors"
	"project-auction/internal/domain/entity"
	"strconv"
)

type PGBidRepository interface {
	Insert(bid entity.Bid) (entity.Bid, error)
}

type pgBidRepository struct {
	log      *slog.Logger
	database *sqlx.DB
}

func NewBidRepository(log *slog.Logger, db *sqlx.DB) PGBidRepository {
	return &pgBidRepository{
		log:      log,
		database: db,
	}
}

func (br pgBidRepository) Insert(bid entity.Bid) (entity.Bid, error) {
	q := `
	INSERT INTO bids(item_id, owner_id, points)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	if err := br.database.QueryRowx(q, bid.ItemID, bid.OwnerID, bid.Points).Scan(&bid.ID); err != nil {
		br.log.Error("failed to execute request to db", slog.String("error", err.Error()))
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == adaperrors.UniqueViolation {
				br.log.Error("failed to execute request to db", slog.String("error", "unique violation"))
				return entity.Bid{}, apperrors.NewConflict("item id, points", strconv.Itoa(bid.ItemID), strconv.FormatFloat(bid.Points, 'f', -1, 64))
			}
		}

		return entity.Bid{}, err
	}

	return bid, nil
}
