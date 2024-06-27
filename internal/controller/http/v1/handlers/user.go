package handlers

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"project-auction/internal/common/apperrors"
	"project-auction/internal/controller/http/v1/dto"
	"project-auction/internal/domain/entity"
)

// RegisterUser 	godoc
//
//	@Summary		Register user
//	@Description	Register user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.CreateUserRequest	true	"model for create user"
//	@Success		201
//	@Failure		400		{object}	apperrors.Error
//	@Failure		500		{object}	apperrors.Error
//	@Router			/users 																																																					[post]
func (h Handler) RegisterUser(c echo.Context) error {
	var req dto.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		h.log.Error("failed to parse request", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), apperrors.NewInternal())
	}

	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, apperrors.NewBadRequest("email or password is empty"))
	}

	user := entity.User{
		Email:    req.Email,
		Password: req.Password,
		Type:     "user",
	}

	userRes, err := h.userService.CreateUser(user)
	if err != nil {
		h.log.Error("failed to create user", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), err)
	}

	jwtPair, err := h.tokenService.GenerateJWTPairTokens(userRes.ID, userRes.Type)
	if err != nil {
		h.log.Error("failed to generate jwt pair tokens", slog.String("error", err.Error()))
		return apperrors.NewInternal()
	}

	return c.JSON(http.StatusCreated, entity.PairJWTClaims{
		AccessToken:  jwtPair.AccessToken,
		RefreshToken: jwtPair.RefreshToken,
	})
}
