package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log/slog"
	"project-auction/apperrors"
	"project-auction/models"
)

type PGItemRepository interface {
	InsertItem(*models.Item) (models.Item, error)
	SelectItems() ([]models.Item, error)
	SelectItemByID(int) (models.Item, error)
	UpdateItem(models.Item) (models.Item, error)
	DeleteItemByID(int) error
}

type pgItemRepository struct {
	log      *slog.Logger
	database *sqlx.DB
}

func NewItemRepository(log *slog.Logger, db *sqlx.DB) PGItemRepository {
	return &pgItemRepository{
		log:      log,
		database: db,
	}
}

func (ir *pgItemRepository) InsertItem(item *models.Item) (models.Item, error) {
	const op = "repository.InsertItem"

	q := `
	INSERT INTO items (name, owner_id, price)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	var id int
	if err := ir.database.QueryRowx(q, item.Name, item.OwnerID, item.Price).Scan(&id); err != nil {
		ir.log.Error(fmt.Sprintf("%s: %v", op, err))

		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return models.Item{}, apperrors.NewConflict("name", item.Name)
			}
		}
		return models.Item{}, apperrors.NewInternal()
	}

	item.ID = id

	return *item, nil
}

func (ir *pgItemRepository) SelectItems() ([]models.Item, error) {
	const op = "repository.SelectItems"

	q := `
	SELECT * FROM items;
	`

	var itemArray []models.Item
	rows, err := ir.database.Queryx(q)
	if err != nil {
		ir.log.Error(fmt.Sprintf("%s:%v", op, err))
		return []models.Item{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.OwnerID, &item.Name, &item.Price); err != nil {
			ir.log.Error(fmt.Sprintf("%s:%v", op, err))
			return []models.Item{}, err
		}

		itemArray = append(itemArray, item)
	}

	if err := rows.Err(); err != nil {
		ir.log.Error(fmt.Sprintf("%s:%v", op, err))
		return []models.Item{}, err
	}

	if len(itemArray) == 0 {
		return []models.Item{}, apperrors.NewNoRows()
	}

	return itemArray, nil
}

func (ir *pgItemRepository) SelectItemByID(id int) (models.Item, error) {
	const op = "repository.SelectItemByID"

	q := `
	SELECT * FROM items
	WHERE id=$1;
	`

	var item models.Item
	if err := ir.database.QueryRowx(q, id).Scan(&item.ID, &item.OwnerID, &item.Name, &item.Price); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Item{}, apperrors.NewNoRows()
		}

		ir.log.Error(fmt.Sprintf("%s:%v", op, err))

		return models.Item{}, apperrors.NewInternal()
	}

	return item, nil
}

func (ir *pgItemRepository) UpdateItem(item models.Item) (models.Item, error) {
	const op = "repository.UpdateItem"

	q := `
	UPDATE items
	SET owner_id=$1, name=$2, price=$3
	WHERE id=$4
	RETURNING id, owner_id, name, price
	`

	var id int
	var updateItem models.Item
	if err := ir.database.QueryRowx(q, item.OwnerID, item.Name, item.Price, item.ID).Scan(&updateItem.ID, &updateItem.OwnerID, &updateItem.Name, &updateItem.Price); err != nil {
		ir.log.Error(fmt.Sprintf("%s: %v", op, err))

		return models.Item{}, apperrors.NewInternal()
	}

	item.ID = id

	return updateItem, nil
}

func (ir *pgItemRepository) DeleteItemByID(id int) error {
	const op = "repository.DeleteItemByID"

	q := `
	DELETE FROM items
	WHERE id=$1
	`

	res, err := ir.database.Exec(q, id)
	if err != nil {
		ir.log.Error(fmt.Sprintf("%s:%v", op, err))
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		ir.log.Error(fmt.Sprintf("%s: %v", op, err))
		return err
	}

	if rowsAffected == 0 {
		return apperrors.NewNoRows()
	}

	return nil
}
