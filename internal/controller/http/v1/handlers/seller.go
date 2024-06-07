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

// RegisterSeller 	godoc
//
//	@Summary		User can become a seller
//	@Description	Create seller
//	@Tags			Sellers
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.CreateSellerRequest	true	"model for create seller"
//	@Success		201
//	@Failure		400			{object}	apperrors.Error
//	@Failure		500			{object}	apperrors.Error
//	@Router			/sellers 																																				[post]
func (h Handler) RegisterSeller(c echo.Context) error {
	var req dto.CreateSellerRequest

	if err := c.Bind(&req); err != nil {
		h.log.Error("failed to parse request", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), apperrors.NewInternal())
	}

	if err := req.Validate(); err != nil {
		h.log.Error("validations failed", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), err)
	}

	inputSeller := entity.Seller{
		Email:    req.Email,
		Password: req.Password,
		Type:     "seller",
	}

	sellerRes, err := h.sellerService.CreateSeller(inputSeller)
	if err != nil {
		h.log.Error("failed to create seller", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), err)
	}

	jwtPairs, err := services.GenerateJWTPairTokens(sellerRes.ID, sellerRes.Type)
	if err != nil {
		h.log.Error("failed to generate jwt pair tokens", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, dto.CreateSellerResponse{
		AccessToken:  jwtPairs.AccessToken,
		RefreshToken: jwtPairs.RefreshToken,
	})
}
