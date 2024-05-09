package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project-auction/apperrors"
	"project-auction/models"
	"project-auction/server/http/httpmodels"
)

// RegisterUser 	godoc
//
//	@Summary		Register user
//	@Description	Register user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body	httpmodels.CreateUserRequest	true	"model for create user"
//	@Success		201
//	@Failure		400		{object}	apperrors.Error
//	@Failure		500		{object}	apperrors.Error
//	@Router			/users 											[post]
func (h Handler) RegisterUser(c echo.Context) error {

	var req httpmodels.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		h.Log.Error("failed to parse request", err)
		return c.JSON(apperrors.Status(err), apperrors.NewInternal())
	}

	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, apperrors.NewBadRequest("email or password is empty"))
	}

	user := &models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	userRes, err := h.UserService.CreateUser(user)
	if err != nil {
		return c.JSON(apperrors.Status(err), err)
	}

	return c.JSON(http.StatusCreated, httpmodels.CreateUserResponse{
		ID: userRes.ID,
		BaseUser: httpmodels.BaseUser{
			Email:    userRes.Email,
			Password: userRes.Password,
		},
	})
}
