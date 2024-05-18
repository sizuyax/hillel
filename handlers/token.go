package handlers

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"project-auction/apperrors"
	"project-auction/server/http/httpmodels"
	"project-auction/services"
)

// RefreshTokens 	godoc
//
//	@Summary		Refresh access token
//	@Description	Refresh access token
//	@Tags			Tokens
//	@Accept			json
//	@Produce		json
//	@Param			request			body		httpmodels.RefreshTokensRequest	true	"model for refresh access token"
//	@Success		200				{object}	httpmodels.RefreshTokensResponse
//	@Failure		400				{object}	apperrors.Error
//	@Failure		500				{object}	apperrors.Error
//	@Router			/auth/tokens 														[post]
func (h Handler) RefreshTokens(c echo.Context) error {
	ctx, err := services.NewContextFromEchoContext(c)
	if err != nil {
		h.log.ErrorContext(ctx, "failed get context from echo context", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	var req httpmodels.RefreshTokensRequest

	if err := c.Bind(&req); err != nil {
		h.log.ErrorContext(ctx, "failed to parse request", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), apperrors.NewInternal())
	}

	jwtPairs, err := services.RefreshAccessJWTToken(req.RefreshToken)
	if err != nil {
		h.log.ErrorContext(ctx, "failed to refresh jwt access token", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), apperrors.NewInternal())
	}

	return c.JSON(http.StatusCreated, httpmodels.RefreshTokensResponse{
		AccessToken:  jwtPairs.AccessToken,
		RefreshToken: jwtPairs.RefreshToken,
	})
}
