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

//go:generate mockery --name=PGUserRepository
type PGUserRepository interface {
	InsertUser(entity.User) (entity.User, error)
}

type pgUserRepository struct {
	log      *slog.Logger
	database *sqlx.DB
}

func NewUserRepository(log *slog.Logger, db *sqlx.DB) PGUserRepository {
	return &pgUserRepository{
		log:      log,
		database: db,
	}
}

func (ur pgUserRepository) InsertUser(user entity.User) (entity.User, error) {
	q := `
	INSERT INTO profiles (email, password, type)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	if err := ur.database.QueryRowx(q, user.Email, user.Password, user.Type).Scan(&user.ID); err != nil {
		ur.log.Error("failed to execute request to db", slog.String("error", err.Error()))
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == adaperrors.UniqueViolation {
				ur.log.Error("conflict", slog.String("error", err.Error()))
				return entity.User{}, apperrors.NewConflict("email", user.Email)
			}
		}

		return entity.User{}, err
	}

	return user, nil
}
