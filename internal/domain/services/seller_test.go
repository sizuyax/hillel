package services

import (
	"github.com/stretchr/testify/assert"
	"project-auction/internal/adapters/postgres/repository/mocks"
	"project-auction/internal/domain/entity"
	"testing"
)

func TestInsertSeller(t *testing.T) {
	mockRepo := mocks.NewPGSellerRepository(t)
	svc := &sellerService{
		SellerRepository: mockRepo,
	}

	inputSeller := entity.Seller{
		Email:    "test@test.com",
		Password: "1234",
	}

	expectedSeller := entity.Seller{
		ID:       1,
		Email:    "test@test.com",
		Password: "1234",
	}

	mockRepo.On("InsertSeller", inputSeller).Return(expectedSeller, nil)

	createSeller, err := svc.CreateSeller(inputSeller)

	assert.NoError(t, err)
	assert.Equal(t, expectedSeller, createSeller)

	mockRepo.AssertExpectations(t)
}
