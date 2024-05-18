package handlers

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"project-auction/apperrors"
	"project-auction/models"
	"project-auction/server/http/httpmodels"
	"project-auction/services"
)

// RegisterSeller 	godoc
//
//	@Summary		User can become a seller
//	@Description	Create seller
//	@Tags			Sellers
//	@Accept			json
//	@Produce		json
//	@Param			request	body	httpmodels.CreateSellerRequest	true	"model for create seller"
//	@Success		201
//	@Failure		400			{object}	apperrors.Error
//	@Failure		500			{object}	apperrors.Error
//	@Router			/sellers 														[post]
func (h Handler) RegisterSeller(c echo.Context) error {
	var req httpmodels.CreateSellerRequest

	if err := c.Bind(&req); err != nil {
		h.log.Error("failed to parse request", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), apperrors.NewInternal())
	}

	if err := req.Validate(); err != nil {
		h.log.Error("validations failed", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), err)
	}

	seller := &models.Seller{
		Email:    req.Email,
		Password: req.Password,
	}

	sellerRes, err := h.sellerService.CreateSeller(seller)
	if err != nil {
		h.log.Error("failed to create seller", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), err)
	}

	jwtPairs, err := services.GenerateJWTPairTokens(sellerRes.ID)
	if err != nil {
		h.log.Error("failed to generate jwt pair tokens", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, httpmodels.CreateSellerResponse{
		AccessToken:  jwtPairs.AccessToken,
		RefreshToken: jwtPairs.RefreshToken,
	})
}
