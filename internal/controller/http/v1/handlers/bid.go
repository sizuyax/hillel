package handlers

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
	"log/slog"
	"net/http"
	"project-auction/internal/common/apperrors"
	"project-auction/internal/controller/http/v1/dto"
	"project-auction/internal/domain/entity"
	"project-auction/internal/domain/services"
	"project-auction/internal/domain/services/files"
	"strconv"
)

// CreateBid 		godoc
//
//	@Summary		Create bid for item
//	@Description	Create bid for item
//	@Tags			Bids
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			id					path		string					true	"item id for bid"
//	@Param			request				body		dto.CreateBidRequest	true	"model for create bid"
//	@Success		200					{object}	entity.Bid
//	@Failure		400					{object}	apperrors.Error
//	@Failure		500					{object}	apperrors.Error
//	@Router			/items/{id}/bids 																																		[post]
func (h Handler) CreateBid(c echo.Context) error {
	ctx, err := services.NewContextFromEchoContext(c)
	if err != nil {
		h.log.Error("failed get context from echo context", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	itemID := c.Param("id")
	itemIDInt, err := strconv.Atoi(itemID)
	if err != nil {
		h.log.Error("failed to parse id", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), apperrors.NewBadRequest("id must be integer."))
	}

	var req dto.CreateBidRequest

	if err := c.Bind(&req); err != nil {
		h.log.Error("failed to parse request", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), apperrors.NewInternal())
	}

	if err := req.Validate(); err != nil {
		h.log.Error("failed to validate req", slog.Any("request", req))
		return c.JSON(apperrors.Status(err), err)
	}

	ownerID, ok := ctx.Value(entity.ProfileIDKey).(int)
	if !ok {
		h.log.Error("failed to set the owner ID")
		return c.NoContent(http.StatusBadRequest)
	}

	inputBid := entity.Bid{
		ItemID:  itemIDInt,
		OwnerID: ownerID,
		Points:  req.Points,
	}

	bidRes, err := h.bidService.Create(inputBid)
	if err != nil {
		h.log.Error("failed to create bid", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), err)
	}

	fileName, err := files.CreateAndWrite(context.TODO(), h.itemService, inputBid.OwnerID, inputBid.ItemID, inputBid.Points, h.log)
	if err != nil {
		h.log.Error("failed to create and write data to file", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	if err := h.minioStorage.Upload(fileName); err != nil {
		h.log.Error("failed to upload file to storage", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	return c.JSON(http.StatusCreated, entity.Bid{
		ID:      bidRes.ID,
		ItemID:  bidRes.ItemID,
		OwnerID: bidRes.OwnerID,
		Points:  bidRes.Points,
	})
}
