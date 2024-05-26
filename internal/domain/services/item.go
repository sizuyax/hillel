package services

import (
	"golang.org/x/net/context"
	"project-auction/internal/adapters/repository/postgres"
	"project-auction/internal/domain/entity"
)

type ItemService interface {
	CreateItem(context.Context, entity.Item) (entity.Item, error)
	GetItems(context.Context) ([]entity.Item, error)
	GetItemByID(context.Context, int) (entity.Item, error)
	UpdateItem(context.Context, entity.Item) (entity.Item, error)
	DeleteItemByID(context.Context, int) error
}

type itemService struct {
	ItemRepository postgres.PGItemRepository
}

func NewItemService(itemRepository postgres.PGItemRepository) ItemService {
	return &itemService{
		ItemRepository: itemRepository,
	}
}

func (is itemService) CreateItem(ctx context.Context, item entity.Item) (entity.Item, error) {
	createItem, err := is.ItemRepository.InsertItem(ctx, item)
	if err != nil {
		return entity.Item{}, err
	}

	return createItem, nil
}

func (is itemService) GetItems(ctx context.Context) ([]entity.Item, error) {
	items, err := is.ItemRepository.SelectItems(ctx)
	if err != nil {
		return []entity.Item{}, err
	}

	return items, nil
}

func (is itemService) GetItemByID(ctx context.Context, id int) (entity.Item, error) {
	item, err := is.ItemRepository.SelectItemByID(ctx, id)
	if err != nil {
		return entity.Item{}, err
	}

	return item, nil
}

func (is itemService) UpdateItem(ctx context.Context, item entity.Item) (entity.Item, error) {
	existsItem, err := is.GetItemByID(ctx, item.ID)
	if err != nil {
		return entity.Item{}, err
	}

	if item.OwnerID == 0 {
		item.OwnerID = existsItem.OwnerID
	}

	if item.Name == "" {
		item.Name = existsItem.Name
	}

	if item.Price == 0 {
		item.Price = existsItem.Price
	}

	updateItem, err := is.ItemRepository.UpdateItem(ctx, item)
	if err != nil {
		return entity.Item{}, err
	}

	return updateItem, nil
}

func (is itemService) DeleteItemByID(ctx context.Context, id int) error {
	if err := is.ItemRepository.DeleteItemByID(ctx, id); err != nil {
		return err
	}

	return nil
}
