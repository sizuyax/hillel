package postgres

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log/slog"
	"project-auction/internal/common/apperrors"
	"project-auction/internal/domain/entity"
)

type PGUserRepository interface {
	InsertUser(*entity.User) (*entity.User, error)
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

func (ur *pgUserRepository) InsertUser(user *entity.User) (*entity.User, error) {
	const op = "repository.InsertUser"

	q := `
	INSERT INTO users (email, password)
	VALUES ($1, $2)
	RETURNING id
	`

	if err := ur.database.QueryRowx(q, user.Email, user.Password).Scan(&user.ID); err != nil {
		ur.log.Error(fmt.Sprintf("%s: %v", op, err))

		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return &entity.User{}, apperrors.NewConflict("email", user.Email)
			}
		}

		return &entity.User{}, fmt.Errorf("%s:%v", op, err)
	}

	return user, nil
}
