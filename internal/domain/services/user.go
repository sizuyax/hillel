package services

import (
	"project-auction/internal/adapters/repository/postgres"
	"project-auction/internal/domain/entity"
)

type UserService interface {
	CreateUser(*entity.User) (*entity.User, error)
}

type userService struct {
	UserRepository postgres.PGUserRepository
}

type USConfig struct {
	UserRepository postgres.PGUserRepository
}

func NewUserService(cfg USConfig) UserService {
	return &userService{
		UserRepository: cfg.UserRepository,
	}
}

func (us *userService) CreateUser(user *entity.User) (*entity.User, error) {

	user, err := us.UserRepository.InsertUser(user)
	if err != nil {
		return &entity.User{}, err
	}

	return user, nil
}
