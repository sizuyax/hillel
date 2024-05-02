package services

import (
	"project-auction/models"
	"project-auction/repository"
)

type UserService interface {
	CreateUser(models.User) (models.User, error)
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

func (us *userService) CreateUser(user models.User) (models.User, error) {

	user, err := us.UserRepository.InsertUser(user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
