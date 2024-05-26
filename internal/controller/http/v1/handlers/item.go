package handlers

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"project-auction/internal/common/apperrors"
	"project-auction/internal/controller/http/v1/dto"
	"project-auction/internal/domain/entity"
	"project-auction/internal/domain/services"
	"strconv"
)

// CreateItem 		godoc
//
//	@Summary		Create item
//	@Description	Create item
//	@Tags			Items
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		httpmodels.CreateItemRequest	true	"model for create item"
//	@Success		201		{object}	entity.Item
//	@Failure		400		{object}	apperrors.Error
//	@Failure		500		{object}	apperrors.Error
//	@Router			/items														    [post]
func (h Handler) CreateItem(c echo.Context) error {
	ctx, err := services.NewContextFromEchoContext(c)
	if err != nil {
		h.log.Error("failed get context from echo context", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	var req dto.CreateItemRequest
	if err := c.Bind(&req); err != nil {
		h.log.ErrorContext(ctx, "failed to parse request", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), apperrors.NewBadRequest("incorrect body request"))
	}

	item := entity.Item{
		Name:    req.Name,
		OwnerID: ctx.Value(entity.SellerIDKey).(int),
		Price:   req.Price,
	}

	createItem, err := h.itemService.CreateItem(ctx, item)
	if err != nil {
		h.log.ErrorContext(ctx, "failed to create item", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), err)
	}

	return c.JSON(http.StatusCreated, entity.Item{
		ID:      createItem.ID,
		OwnerID: createItem.OwnerID,
		Name:    createItem.Name,
		Price:   createItem.Price,
	})
}

// GetItems 		godoc
//
//	@Summary		Get items
//	@Description	Get all items
//	@Tags			Items
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	[]entity.Item
//	@Failure		500		{object}	apperrors.Error
//	@Router			/items 														[get]
func (h Handler) GetItems(c echo.Context) error {
	ctx, err := services.NewContextFromEchoContext(c)
	if err != nil {
		h.log.Error("failed get context from echo context", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	itemsArray, err := h.itemService.GetItems(ctx)
	if err != nil {
		h.log.ErrorContext(ctx, "failed to get items", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), err)
	}

	itemArrayResponse := make([]entity.Item, 0, len(itemsArray))

	for _, itemValue := range itemsArray {
		itemArrayResponse = append(itemArrayResponse, entity.Item{
			ID:      itemValue.ID,
			OwnerID: itemValue.OwnerID,
			Name:    itemValue.Name,
			Price:   itemValue.Price,
		})
	}

	return c.JSON(http.StatusOK, itemArrayResponse)
}

// GetItemByID 		godoc
//
//	@Summary		Get item
//	@Description	Get item by id
//	@Tags			Items
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"get item by id"
//	@Success		200				{object}	entity.Item
//	@Failure		400				{object}	apperrors.Error
//	@Failure		500				{object}	apperrors.Error
//	@Router			/items/{id} 													[get]
func (h Handler) GetItemByID(c echo.Context) error {
	ctx, err := services.NewContextFromEchoContext(c)
	if err != nil {
		h.log.Error("failed get context from echo context", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.log.ErrorContext(ctx, "failed to parse id", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), apperrors.NewBadRequest("id must be integer."))
	}

	item, err := h.itemService.GetItemByID(ctx, idInt)
	if err != nil {
		h.log.ErrorContext(ctx, "failed to get item by id with id", slog.Int("id", idInt), slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), err)
	}

	return c.JSON(http.StatusOK, entity.Item{
		ID:      item.ID,
		OwnerID: item.OwnerID,
		Name:    item.Name,
		Price:   item.Price,
	})
}

// UpdateItem 		godoc
//
//	@Summary		Update item
//	@Description	Update item
//	@Tags			Items
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			request			body		httpmodels.UpdateItemRequest	true	"model for update item"
//	@Success		200				{object}	entity.Item
//	@Failure		400				{object}	apperrors.Error
//	@Failure		500				{object}	apperrors.Error
//	@Router			/items/{id} 													[put]
func (h Handler) UpdateItem(c echo.Context) error {
	ctx, err := services.NewContextFromEchoContext(c)
	if err != nil {
		h.log.Error("failed get context from echo context", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	var req dto.UpdateItemRequest

	if err := c.Bind(&req); err != nil {
		h.log.ErrorContext(ctx, "failed to parse request", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), apperrors.NewInternal())
	}

	if err := req.Validate(); err != nil {
		h.log.ErrorContext(ctx, "validation failed", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), err)
	}

	item := entity.Item{
		ID:      req.ID,
		Name:    req.Name,
		OwnerID: ctx.Value(entity.SellerIDKey).(int),
		Price:   req.Price,
	}

	updateItem, err := h.itemService.UpdateItem(ctx, item)
	if err != nil {
		h.log.ErrorContext(ctx, "failed to update item", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), err)
	}

	return c.JSON(http.StatusOK, entity.Item{
		ID:      updateItem.ID,
		OwnerID: updateItem.OwnerID,
		Name:    updateItem.Name,
		Price:   updateItem.Price,
	})
}

// DeleteItemByID 		godoc
//
//	@Summary		Delete item
//	@Description	Delete item
//	@Tags			Items
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"delete item by id"
//	@Success		200
//	@Failure		400				{object}	apperrors.Error
//	@Failure		500				{object}	apperrors.Error
//	@Router			/items/{id} 													[delete]
func (h Handler) DeleteItemByID(c echo.Context) error {
	ctx, err := services.NewContextFromEchoContext(c)
	if err != nil {
		h.log.Error("failed get context from echo context", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.log.ErrorContext(ctx, "failed to parse id", slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), apperrors.NewBadRequest("id must be integer."))
	}

	if err := h.itemService.DeleteItemByID(ctx, idInt); err != nil {
		h.log.ErrorContext(ctx, "failed to delete item with id", slog.Int("id", idInt), slog.String("error", err.Error()))
		return c.JSON(apperrors.Status(err), err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "ok",
	})
}
