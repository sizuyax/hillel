package repository

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"project-auction/internal/common/apperrors"
	"project-auction/internal/domain/entity"
)

//go:generate mockery --name=PGItemRepository
type PGItemRepository interface {
	InsertItem(context.Context, entity.Item) (entity.Item, error)
	SelectItems(context.Context) ([]entity.Item, error)
	SelectItemByID(context.Context, int) (entity.Item, error)
	UpdateItem(context.Context, entity.Item) (entity.Item, error)
	DeleteItemByID(context.Context, int) error
}

type pgItemRepository struct {
	database *sqlx.DB
}

func NewItemRepository(db *sqlx.DB) PGItemRepository {
	return &pgItemRepository{
		database: db,
	}
}

func (ir *pgItemRepository) InsertItem(ctx context.Context, item entity.Item) (entity.Item, error) {
	q := `
	INSERT INTO items (owner_id, name, price)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	if err := ir.database.QueryRowxContext(ctx, q, item.OwnerID, item.Name, item.Price).Scan(&item.ID); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return entity.Item{}, apperrors.NewConflict("name", item.Name)
			}
			if pqErr.Code == "23503" {
				return entity.Item{}, apperrors.NewUnprocessable()
			}
		}
		return entity.Item{}, err
	}

	return item, nil
}

func (ir *pgItemRepository) SelectItems(ctx context.Context) ([]entity.Item, error) {
	q := `
	SELECT * FROM items;
	`

	var itemArray []entity.Item
	rows, err := ir.database.QueryxContext(ctx, q)
	if err != nil {
		return []entity.Item{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var item entity.Item
		if err := rows.Scan(&item.ID, &item.OwnerID, &item.Name, &item.Price); err != nil {
			return []entity.Item{}, err
		}

		itemArray = append(itemArray, item)
	}

	if err := rows.Err(); err != nil {
		return []entity.Item{}, err
	}

	if len(itemArray) == 0 {
		return []entity.Item{}, apperrors.NewNoRows()
	}

	return itemArray, nil
}

func (ir *pgItemRepository) SelectItemByID(ctx context.Context, id int) (entity.Item, error) {
	q := `
	SELECT * FROM items
	WHERE id=$1;
	`

	var item entity.Item
	if err := ir.database.QueryRowxContext(ctx, q, id).Scan(&item.ID, &item.OwnerID, &item.Name, &item.Price); err != nil {
		return entity.Item{}, err
	}

	return item, nil
}

func (ir *pgItemRepository) UpdateItem(ctx context.Context, item entity.Item) (entity.Item, error) {
	q := `
	UPDATE items
	SET owner_id=$1, name=$2, price=$3
	WHERE id=$4
	RETURNING id, owner_id, name, price
	`

	var updateItem entity.Item
	if err := ir.database.QueryRowxContext(ctx, q, item.OwnerID, item.Name, item.Price, item.ID).Scan(&updateItem.ID, &updateItem.OwnerID, &updateItem.Name, &updateItem.Price); err != nil {
		return entity.Item{}, err
	}

	return updateItem, nil
}

func (ir *pgItemRepository) DeleteItemByID(ctx context.Context, id int) error {
	q := `
	DELETE FROM items
	WHERE id=$1
	`

	res, err := ir.database.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.NewNoRows()
	}

	return nil
}
