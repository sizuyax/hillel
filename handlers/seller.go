package handlers

import (
	"github.com/labstack/echo/v4"
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
//	@Router			/sellers 										[post]
func (h Handler) RegisterSeller(c echo.Context) error {
	var req httpmodels.CreateSellerRequest

	if err := c.Bind(&req); err != nil {
		h.Log.Error("failed to parse request", err)
		return c.JSON(apperrors.Status(err), apperrors.NewInternal())
	}

	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, apperrors.NewBadRequest("email or password is empty"))
	}

	seller := &models.Seller{
		Email:    req.Email,
		Password: req.Password,
	}

	sellerRes, err := h.SellerService.CreateSeller(seller)
	if err != nil {
		return c.JSON(apperrors.Status(err), err)
	}

	accessToken, err := services.GenerateJWTAccessToken(sellerRes.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	return c.JSON(http.StatusCreated, httpmodels.CreateSellerResponse{
		AccessToken: accessToken,
	})
}
