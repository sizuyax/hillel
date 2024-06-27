package handlers

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"project-auction/internal/common/apperrors"
	"project-auction/internal/controller/http/v1/dto"
	"project-auction/internal/domain/entity"
	"project-auction/internal/domain/services"
	"strconv"
)

// CreateComment 		godoc
//
//	@Summary		Comment item
//	@Description	Comment item
//	@Tags			Comments
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			id						path		string						true	"comment item with id"
//	@Param			request					body		dto.CreateCommentRequest	true	"model for create comment"
//	@Success		200						{object}	entity.Comment
//	@Failure		400						{object}	apperrors.Error
//	@Failure		500						{object}	apperrors.Error
//	@Router			/items/{id}/comments 																													[post]																	[get]
func (h Handler) CreateComment(c echo.Context) error {
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

	itemRes, err := h.itemService.GetItemByID(ctx, itemIDInt)
	if err != nil {
		h.log.Error("failed to get item", slog.Int("id", itemIDInt), slog.String("error", err.Error()))
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(apperrors.Status(err), apperrors.NewNoRows())
		}
		return c.JSON(apperrors.Status(err), apperrors.NewInternal())
	}

	var req dto.CreateCommentRequest

	if err := c.Bind(&req); err != nil {
		h.log.Error("failed to parse request", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), apperrors.NewInternal())
	}

	inputComment := entity.Comment{
		ItemID:  itemRes.ID,
		OwnerID: ctx.Value(entity.ProfileIDKey).(int),
		Body:    req.Body,
	}

	commentRes, err := h.commentService.CreateComment(inputComment)
	if err != nil {
		h.log.Error("failed to create comment", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), err)
	}

	return c.JSON(http.StatusCreated, entity.Comment{
		ID:      commentRes.ID,
		ItemID:  commentRes.ItemID,
		OwnerID: commentRes.OwnerID,
		Body:    commentRes.Body,
	})
}
