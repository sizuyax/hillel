package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log/slog"
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
	log      *slog.Logger
	database *sqlx.DB
}

func NewItemRepository(log *slog.Logger, db *sqlx.DB) PGItemRepository {
	return &pgItemRepository{
		log:      log,
		database: db,
	}
}

func (ir *pgItemRepository) InsertItem(ctx context.Context, item entity.Item) (entity.Item, error) {
	const op = "repository.InsertItem"

	q := `
	INSERT INTO items (owner_id, name, price)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	if err := ir.database.QueryRowxContext(ctx, q, item.OwnerID, item.Name, item.Price).Scan(&item.ID); err != nil {
		ir.log.ErrorContext(ctx, fmt.Sprintf("%s: %v", op, err))

		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return entity.Item{}, apperrors.NewConflict("name", item.Name)
			}
		}
		return entity.Item{}, fmt.Errorf("%s:%v", op, err)
	}

	return item, nil
}

func (ir *pgItemRepository) SelectItems(ctx context.Context) ([]entity.Item, error) {
	const op = "repository.SelectItems"

	q := `
	SELECT * FROM items;
	`

	var itemArray []entity.Item
	rows, err := ir.database.QueryxContext(ctx, q)
	if err != nil {
		ir.log.ErrorContext(ctx, fmt.Sprintf("%s:%v", op, err))
		return []entity.Item{}, fmt.Errorf("%s:%v", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var item entity.Item
		if err := rows.Scan(&item.ID, &item.OwnerID, &item.Name, &item.Price); err != nil {
			ir.log.ErrorContext(ctx, fmt.Sprintf("%s:%v", op, err))
			return []entity.Item{}, fmt.Errorf("%s:%v", op, err)
		}

		itemArray = append(itemArray, item)
	}

	if err := rows.Err(); err != nil {
		ir.log.ErrorContext(ctx, fmt.Sprintf("%s:%v", op, err))
		return []entity.Item{}, fmt.Errorf("%s:%v", op, err)
	}

	if len(itemArray) == 0 {
		return []entity.Item{}, apperrors.NewNoRows()
	}

	return itemArray, nil
}

func (ir *pgItemRepository) SelectItemByID(ctx context.Context, id int) (entity.Item, error) {
	const op = "repository.SelectItemByID"

	q := `
	SELECT * FROM items
	WHERE id=$1;
	`

	var item entity.Item
	if err := ir.database.QueryRowxContext(ctx, q, id).Scan(&item.ID, &item.OwnerID, &item.Name, &item.Price); err != nil {
		ir.log.ErrorContext(ctx, fmt.Sprintf("%s:%v", op, err))

		if errors.Is(err, sql.ErrNoRows) {
			return entity.Item{}, apperrors.NewNoRows()
		}

		return entity.Item{}, fmt.Errorf("%s:%v", op, err)
	}

	return item, nil
}

func (ir *pgItemRepository) UpdateItem(ctx context.Context, item entity.Item) (entity.Item, error) {
	const op = "repository.UpdateItem"

	q := `
	UPDATE items
	SET owner_id=$1, name=$2, price=$3
	WHERE id=$4
	RETURNING id, owner_id, name, price
	`

	var updateItem entity.Item
	if err := ir.database.QueryRowxContext(ctx, q, item.OwnerID, item.Name, item.Price, item.ID).Scan(&updateItem.ID, &updateItem.OwnerID, &updateItem.Name, &updateItem.Price); err != nil {
		ir.log.ErrorContext(ctx, fmt.Sprintf("%s: %v", op, err))

		return entity.Item{}, fmt.Errorf("%s:%v", op, err)
	}

	return updateItem, nil
}

func (ir *pgItemRepository) DeleteItemByID(ctx context.Context, id int) error {
	const op = "repository.DeleteItemByID"

	q := `
	DELETE FROM items
	WHERE id=$1
	`

	res, err := ir.database.ExecContext(ctx, q, id)
	if err != nil {
		ir.log.Error(fmt.Sprintf("%s:%v", op, err))
		return fmt.Errorf("%s:%v", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		ir.log.Error(fmt.Sprintf("%s: %v", op, err))
		return fmt.Errorf("%s:%v", op, err)
	}

	if rowsAffected == 0 {
		return apperrors.NewNoRows()
	}

	return nil
}
