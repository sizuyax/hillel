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

type PGUserRepository interface {
	InsertUser(*models.User) (*models.User, error)
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

func (ur *pgUserRepository) InsertUser(user *models.User) (*models.User, error) {
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
				return &models.User{}, apperrors.NewConflict("email", user.Email)
			}
		}

		return &models.User{}, fmt.Errorf("%s:%v", op, err)
	}

	return user, nil
}
