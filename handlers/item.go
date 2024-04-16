package handlers

import "github.com/labstack/echo/v4"

// GetItems 		godoc
// @Summary 		Get items
// @Description 	Get all items
// @Tags 			models.Item
// @Accept  		json
// @Produce  		json
// @Success 		200 				{object} 	[]Item
// @Router 			/items 				[get]
func (h Handler) GetItems(c echo.Context) error {
	panic("implement me")
}

// GetItemByID 		godoc
// @Summary 		Get item
// @Description 	Get item by id
// @Tags 			models.Item
// @Accept  		json
// @Produce  		json
// @Success 		200 				{object} 	Item
// @Router 			/item{id} 			[get]
func (h Handler) GetItemByID(c echo.Context) error {
	panic("implement me")
}

// CreateItem 		godoc
// @Summary 		Create item
// @Description 	Create item
// @Tags 			models.Item
// @Accept  		json
// @Produce  		json
// @Success 		201
// @Router 			/create-item 		[post]
func (h Handler) CreateItem(c echo.Context) error {
	panic("implement me")
}

// UpdateItem 		godoc
// @Summary 		Update item
// @Description 	Update item
// @Tags 			models.Item
// @Accept  		json
// @Produce  		json
// @Success 		200
// @Router 			/update-item 		[put]
func (h Handler) UpdateItem(c echo.Context) error {
	panic("implement me")
}

// DeleteItem 		godoc
// @Summary 		Delete item
// @Description 	Delete item
// @Tags 			models.Item
// @Accept  		json
// @Produce  		json
// @Success 		200
// @Router 			/delete-item 		[delete]
func (h Handler) DeleteItem(c echo.Context) error {
	panic("implement me")
}
