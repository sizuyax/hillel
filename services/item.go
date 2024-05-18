package services

import (
	"golang.org/x/net/context"
	"project-auction/models"
	"project-auction/repository"
)

type ItemService interface {
	CreateItem(context.Context, *models.Item) (models.Item, error)
	GetItems(context.Context) ([]models.Item, error)
	GetItemByID(context.Context, int) (models.Item, error)
	UpdateItem(context.Context, models.Item) (models.Item, error)
	DeleteItemByID(context.Context, int) error
}

type itemService struct {
	ItemRepository repository.PGItemRepository
}

type ISConfig struct {
	ItemRepository repository.PGItemRepository
}

func NewItemService(cfg ISConfig) ItemService {
	return &itemService{
		ItemRepository: cfg.ItemRepository,
	}
}

func (is *itemService) CreateItem(ctx context.Context, item *models.Item) (models.Item, error) {
	createItem, err := is.ItemRepository.InsertItem(ctx, item)
	if err != nil {
		return models.Item{}, err
	}

	return createItem, nil
}

func (is *itemService) GetItems(ctx context.Context) ([]models.Item, error) {
	items, err := is.ItemRepository.SelectItems(ctx)
	if err != nil {
		return []models.Item{}, err
	}

	return items, nil
}

func (is *itemService) GetItemByID(ctx context.Context, id int) (models.Item, error) {
	item, err := is.ItemRepository.SelectItemByID(ctx, id)
	if err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func (is *itemService) UpdateItem(ctx context.Context, item models.Item) (models.Item, error) {
	existsItem, err := is.GetItemByID(ctx, item.ID)
	if err != nil {
		return models.Item{}, err
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
		return models.Item{}, err
	}

	return updateItem, nil
}

func (is *itemService) DeleteItemByID(ctx context.Context, id int) error {
	if err := is.ItemRepository.DeleteItemByID(ctx, id); err != nil {
		return err
	}

	return nil
}
