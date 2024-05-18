package httpmodels

import (
	"github.com/go-playground/validator/v10"
	"project-auction/apperrors"
)

type BaseItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateItemRequest struct {
	BaseItem
}

type CreateItemResponse struct {
	ID      int `json:"id"`
	OwnerID int `json:"owner_id"`
	BaseItem
}

type GetItemResponse struct {
	ID      int `json:"id"`
	OwnerID int `json:"owner_id"`
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

type UpdateItemResponse struct {
	ID      int `json:"id"`
	OwnerID int `json:"owner_id"`
	BaseItem
}
