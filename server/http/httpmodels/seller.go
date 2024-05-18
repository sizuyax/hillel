package httpmodels

import (
	"github.com/go-playground/validator/v10"
	"project-auction/apperrors"
)

type BaseSeller struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateSellerRequest struct {
	BaseSeller
}

func (r *CreateSellerRequest) Validate() error {
	validate := validator.New()

	if err := validate.Var(r.Email, "required"); err != nil {
		return apperrors.NewBadRequest("email is required")
	}

	if err := validate.Var(r.Password, "required"); err != nil {
		return apperrors.NewBadRequest("password is required")
	}

	return nil
}

type CreateSellerResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
