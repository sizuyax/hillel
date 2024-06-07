package services

import (
	"log/slog"
	"project-auction/internal/adapters/postgres/repository"
	"project-auction/internal/domain/entity"
)

type UserService interface {
	CreateUser(entity.User) (entity.User, error)
}

type userService struct {
	log            *slog.Logger
	UserRepository repository.PGUserRepository
}

func NewUserService(log *slog.Logger, userRepository repository.PGUserRepository) UserService {
	return &userService{
		log:            log,
		UserRepository: userRepository,
	}
}

func (us *userService) CreateUser(inputUser entity.User) (entity.User, error) {
	user, err := us.UserRepository.InsertUser(inputUser)
	if err != nil {
		us.log.Error("failed insert user", slog.String("error", err.Error()))
		return entity.User{}, err
	}

	return user, nil
}
