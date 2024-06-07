package services

import (
	"project-auction/internal/adapters/postgres/repository"
	"project-auction/internal/domain/entity"
)

type UserService interface {
	CreateUser(entity.User) (entity.User, error)
}

type userService struct {
	UserRepository repository.PGUserRepository
}

type USConfig struct {
	UserRepository repository.PGUserRepository
}

func NewUserService(cfg USConfig) UserService {
	return &userService{
		UserRepository: cfg.UserRepository,
	}
}

func (us *userService) CreateUser(inputUser entity.User) (entity.User, error) {
	user, err := us.UserRepository.InsertUser(inputUser)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
