package services

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"project-auction/internal/adapters/repository/postgres/mocks"
	"project-auction/internal/domain/entity"
	"testing"
)

func TestCreateItem(t *testing.T) {
	type TestCase struct {
		Name           string
		InsertItemCall *mock.Call
		InputItem      entity.Item
		ExpectedItem   entity.Item
		ExpectedError  error
	}

	mockRepo := mocks.NewPGItemRepository(t)
	svc := &itemService{
		ItemRepository: mockRepo,
	}

	ctx := context.Background()
	inputItem := entity.Item{
		OwnerID: 1,
		Name:    "Test",
		Price:   20,
	}
	expectedItem := entity.Item{
		ID:      1,
		OwnerID: 1,
		Name:    "Test",
		Price:   10,
	}

	var testCases = []TestCase{
		{
			Name:           "successfully create item",
			InsertItemCall: mockRepo.On("InsertItem", ctx, inputItem).Return(expectedItem, nil),
			InputItem:      inputItem,
			ExpectedItem:   expectedItem,
			ExpectedError:  nil,
		},
	}

	for _, tc := range testCases {
		createItem, err := svc.CreateItem(ctx, tc.InputItem)

		if assert.NoError(t, err) {
			assert.Equal(t, tc.ExpectedItem, createItem)
			assert.Equal(t, 1, createItem.ID)
			assert.Equal(t, 1, createItem.OwnerID)
			assert.Equal(t, float64(10), createItem.Price)
		}
		mockRepo.AssertExpectations(t)

		fmt.Println(createItem, "---> error: ", err)
	}
}

func TestUpdateItem(t *testing.T) {
	type TestCase struct {
		Name               string
		SelectItemByIDCall *mock.Call
		UpdateItemCall     *mock.Call
		InputItem          entity.Item
		ExpectedItem       entity.Item
		ExpectedError      error
	}

	mockRepo := mocks.NewPGItemRepository(t)
	svc := &itemService{
		ItemRepository: mockRepo,
	}

	ctx := context.Background()
	inputItem := entity.Item{
		ID:   1,
		Name: "Test new",
	}

	existsItem := entity.Item{
		ID:      1,
		OwnerID: 1,
		Name:    "Test",
		Price:   20,
	}

	updatedModelItem := entity.Item{
		ID:      1,
		OwnerID: 1,
		Name:    inputItem.Name,
		Price:   20,
	}

	var testCases = []TestCase{
		{
			Name:               "successfully update item",
			SelectItemByIDCall: mockRepo.On("SelectItemByID", ctx, 1).Return(existsItem, nil),
			UpdateItemCall:     mockRepo.On("UpdateItem", ctx, updatedModelItem).Return(updatedModelItem, nil),
			InputItem:          inputItem,
			ExpectedItem:       updatedModelItem,
			ExpectedError:      nil,
		},
	}

	for _, tc := range testCases {
		updateItem, err := svc.UpdateItem(ctx, tc.InputItem)

		if assert.NoError(t, err) {
			assert.Equal(t, tc.ExpectedItem, updateItem)
			assert.Equal(t, inputItem.Name, updateItem.Name)
		}

		mockRepo.AssertExpectations(t)

		fmt.Println(updateItem, "---> error: ", err)
	}

}