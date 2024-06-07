package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"project-auction/internal/common/apperrors"
	"project-auction/internal/domain/entity"
)

//go:generate mockery --name=PGCommentRepository
type PGCommentRepository interface {
	InsertComment(entity.Comment) (entity.Comment, error)
}

type pgCommentRepository struct {
	database *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) PGCommentRepository {
	return &pgCommentRepository{
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
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return entity.Comment{}, apperrors.NewConflict("body", comment.Body)
			}
		}

		return entity.Comment{}, err
	}

	return comment, nil
}
