package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"project-auction/internal/adapters/repository/postgres/mocks"
	"project-auction/internal/domain/entity"
	"testing"
)

func TestInsertUser(t *testing.T) {
	mockRepo := mocks.NewPGUserRepository(t)
	svc := &userService{
		UserRepository: mockRepo,
	}

	inputUser := &entity.User{
		Email:    "test@test.com",
		Password: "test",
	}

	expectedUser := &entity.User{
		ID:       1,
		Email:    "test@test.com",
		Password: "test",
	}

	mockRepo.On("InsertUser", inputUser).Return(expectedUser, nil)

	createUser, err := svc.CreateUser(inputUser)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, createUser)

	mockRepo.AssertExpectations(t)

	fmt.Println(createUser, "---> error: ", err)
}
