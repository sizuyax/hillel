package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project-auction/apperrors"
	"project-auction/models"
	"project-auction/server/http/httpmodels"
	"project-auction/services"
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
//	@Success		201		{object}	models.Item
//	@Failure		400		{object}	apperrors.Error
//	@Failure		500		{object}	apperrors.Error
//	@Router			/items										    [post]
func (h Handler) CreateItem(c echo.Context) error {
	ctx, err := services.NewContextFromEchoContext(c)
	if err != nil {
		h.Log.Error("failed get context from echo context", err)
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	var req httpmodels.CreateItemRequest

	if err := c.Bind(&req); err != nil {
		h.Log.ErrorContext(ctx, "failed to parse request", err)
		return c.JSON(apperrors.Status(err), apperrors.NewBadRequest("incorrect body request"))
	}

	item := &models.Item{
		Name:    req.Name,
		OwnerID: ctx.Value(models.SellerIDKey).(int),
		Price:   req.Price,
	}

	createItem, err := h.ItemService.CreateItem(ctx, item)
	if err != nil {
		h.Log.ErrorContext(ctx, "failed to create item", err)
		return c.JSON(apperrors.Status(err), err)
	}

	return c.JSON(http.StatusCreated, httpmodels.CreateItemResponse{
		ID:      createItem.ID,
		OwnerID: createItem.OwnerID,
		BaseItem: httpmodels.BaseItem{
			Name:  createItem.Name,
			Price: createItem.Price,
		},
	})
}

// GetItems 		godoc
//
//	@Summary		Get items
//	@Description	Get all items
//	@Tags			Items
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	[]models.Item
//	@Failure		500		{object}	apperrors.Error
//	@Router			/items 										[get]
func (h Handler) GetItems(c echo.Context) error {
	ctx, err := services.NewContextFromEchoContext(c)
	if err != nil {
		h.Log.Error("failed get context from echo context", err)
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	itemsArray, err := h.ItemService.GetItems(ctx)
	if err != nil {
		h.Log.ErrorContext(ctx, "failed to get items", err)
		return c.JSON(apperrors.Status(err), err)
	}

	itemArrayResponse := make([]httpmodels.GetItemResponse, 0, len(itemsArray))

	for _, itemValue := range itemsArray {
		itemArrayResponse = append(itemArrayResponse, httpmodels.GetItemResponse{
			ID:      itemValue.ID,
			OwnerID: itemValue.OwnerID,
			BaseItem: httpmodels.BaseItem{
				Name:  itemValue.Name,
				Price: itemValue.Price,
			},
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
//	@Success		200				{object}	models.Item
//	@Failure		400				{object}	apperrors.Error
//	@Failure		500				{object}	apperrors.Error
//	@Router			/items/{id} 									[get]
func (h Handler) GetItemByID(c echo.Context) error {
	ctx, err := services.NewContextFromEchoContext(c)
	if err != nil {
		h.Log.Error("failed get context from echo context", err)
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.Log.ErrorContext(ctx, "failed to parse id", err)
		return c.JSON(apperrors.Status(err), apperrors.NewBadRequest("id must be integer."))
	}

	item, err := h.ItemService.GetItemByID(ctx, idInt)
	if err != nil {
		h.Log.ErrorContext(ctx, "failed to get item by id with id", idInt, "error", err)
		return c.JSON(apperrors.Status(err), err)
	}

	return c.JSON(http.StatusOK, httpmodels.GetItemResponse{
		ID:      item.ID,
		OwnerID: item.OwnerID,
		BaseItem: httpmodels.BaseItem{
			Name:  item.Name,
			Price: item.Price,
		},
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
//	@Success		200				{object}	models.Item
//	@Failure		400				{object}	apperrors.Error
//	@Failure		500				{object}	apperrors.Error
//	@Router			/items/{id} 									[put]
func (h Handler) UpdateItem(c echo.Context) error {
	ctx, err := services.NewContextFromEchoContext(c)
	if err != nil {
		h.Log.Error("failed get context from echo context", err)
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	var req httpmodels.UpdateItemRequest

	if err := c.Bind(&req); err != nil {
		h.Log.ErrorContext(ctx, "failed to parse request", err)
		return c.JSON(apperrors.Status(err), apperrors.NewInternal())
	}

	if req.ID == 0 {
		return c.JSON(http.StatusBadRequest, apperrors.NewBadRequest("id is require."))
	}

	item := models.Item{
		ID:      req.ID,
		Name:    req.Name,
		OwnerID: ctx.Value(models.SellerIDKey).(int),
		Price:   req.Price,
	}

	updateItem, err := h.ItemService.UpdateItem(ctx, item)
	if err != nil {
		h.Log.ErrorContext(ctx, "failed to update item", err)
		return c.JSON(apperrors.Status(err), err)
	}

	return c.JSON(http.StatusOK, httpmodels.UpdateItemResponse{
		ID:      updateItem.ID,
		OwnerID: updateItem.OwnerID,
		BaseItem: httpmodels.BaseItem{
			Name:  updateItem.Name,
			Price: updateItem.Price,
		},
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
//	@Router			/items/{id} 									[delete]
func (h Handler) DeleteItemByID(c echo.Context) error {
	ctx, err := services.NewContextFromEchoContext(c)
	if err != nil {
		h.Log.Error("failed get context from echo context", err)
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.Log.ErrorContext(ctx, "failed to parse id", err)
		return c.JSON(apperrors.Status(err), apperrors.NewBadRequest("id must be integer."))
	}

	if err := h.ItemService.DeleteItemByID(ctx, idInt); err != nil {
		h.Log.ErrorContext(ctx, "failed to delete item with id", idInt, "error", err)
		return c.JSON(apperrors.Status(err), err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "ok",
	})
}
