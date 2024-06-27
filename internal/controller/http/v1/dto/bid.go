package dto

import (
	"github.com/go-playground/validator/v10"
	"project-auction/internal/common/apperrors"
)

type CreateBidRequest struct {
	Points float64 `json:"points"`
}

func (cbr CreateBidRequest) Validate() error {
	validate := validator.New()

	if err := validate.Var(cbr.Points, "required"); err != nil {
		return apperrors.NewBadRequest("id is require.")
	}

	return nil
}
