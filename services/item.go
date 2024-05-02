package services

import (
	"project-auction/models"
	"project-auction/repository"
)

type ItemService interface {
	CreateItem(*models.Item) (models.Item, error)
	GetItems() ([]models.Item, error)
	GetItemByID(int) (models.Item, error)
	UpdateItem(models.Item) (models.Item, error)
	DeleteItemByID(id int) error
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

func (is *itemService) CreateItem(item *models.Item) (models.Item, error) {
	createItem, err := is.ItemRepository.InsertItem(item)
	if err != nil {
		return models.Item{}, err
	}

	return createItem, nil
}

func (is *itemService) GetItems() ([]models.Item, error) {
	items, err := is.ItemRepository.SelectItems()
	if err != nil {
		return []models.Item{}, err
	}

	return items, nil
}

func (is *itemService) GetItemByID(id int) (models.Item, error) {
	item, err := is.ItemRepository.SelectItemByID(id)
	if err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func (is *itemService) UpdateItem(item models.Item) (models.Item, error) {
	existsItem, err := is.GetItemByID(item.ID)
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

	updateItem, err := is.ItemRepository.UpdateItem(item)
	if err != nil {
		return models.Item{}, err
	}

	return updateItem, nil
}

func (is *itemService) DeleteItemByID(id int) error {
	if err := is.ItemRepository.DeleteItemByID(id); err != nil {
		return err
	}

	return nil
}
