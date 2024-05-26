package dto

import (
	"github.com/go-playground/validator/v10"
	"project-auction/internal/common/apperrors"
)

type BaseItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateItemRequest struct {
	BaseItem
}

type UpdateItemRequest struct {
	ID int `json:"id"`
	BaseItem
}

func (uir *UpdateItemRequest) Validate() error {
	validate := validator.New()

	if err := validate.Var(uir.ID, "required"); err != nil {
		return apperrors.NewBadRequest("id is require.")
	}

	return nil
}
