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

//go:generate mockery --name=PGCommentRepository
type PGCommentRepository interface {
	InsertComment(entity.Comment) (entity.Comment, error)
}

type pgCommentRepository struct {
	log      *slog.Logger
	database *sqlx.DB
}

func NewCommentRepository(log *slog.Logger, db *sqlx.DB) PGCommentRepository {
	return &pgCommentRepository{
		log:      log,
		database: db,
	}
}

func (cr pgCommentRepository) InsertComment(comment entity.Comment) (entity.Comment, error) {
	q := `
	INSERT INTO comments (item_id, owner_id, body)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	if err := cr.database.QueryRowx(q, comment.ItemID, comment.OwnerID, comment.Body).Scan(&comment.ID); err != nil {
		cr.log.Error("failed to execute request to db", slog.String("error", err.Error()))
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == adaperrors.UniqueViolation {
				cr.log.Error("failed to execute request to db", slog.String("error", "unique violation"))
				return entity.Comment{}, apperrors.NewConflict("body", comment.Body)
			}
		}

		return entity.Comment{}, err
	}

	return comment, nil
}
