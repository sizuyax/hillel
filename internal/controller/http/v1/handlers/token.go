package handlers

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"project-auction/internal/common/apperrors"
	"project-auction/internal/controller/http/v1/dto"
	"project-auction/internal/domain/entity"
	"project-auction/internal/domain/services"
)

// RefreshTokens 	godoc
//
//	@Summary		Refresh access token
//	@Description	Refresh access token
//	@Tags			Tokens
//	@Accept			json
//	@Produce		json
//	@Param			request			body		dto.RefreshTokensRequest	true	"model for refresh access token"
//	@Success		200				{object}	entity.PairJWTClaims
//	@Failure		400				{object}	apperrors.Error
//	@Failure		500				{object}	apperrors.Error
//	@Router			/auth/tokens 																																						[post]
func (h Handler) RefreshTokens(c echo.Context) error {
	ctx, err := services.NewContextFromEchoContext(c)
	if err != nil {
		h.log.ErrorContext(ctx, "failed get context from echo context", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	var req dto.RefreshTokensRequest

	if err := c.Bind(&req); err != nil {
		h.log.ErrorContext(ctx, "failed to parse request", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), apperrors.NewInternal())
	}

	jwtPairs, err := h.tokenService.RefreshAccessJWTToken(req.RefreshToken)
	if err != nil {
		h.log.ErrorContext(ctx, "failed to refresh jwt access token", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), apperrors.NewInternal())
	}

	return c.JSON(http.StatusCreated, entity.PairJWTClaims{
		AccessToken:  jwtPairs.AccessToken,
		RefreshToken: jwtPairs.RefreshToken,
	})
}
