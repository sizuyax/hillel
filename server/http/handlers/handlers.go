package handlers

import "github.com/labstack/echo/v4"

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// GetItems 		godoc
// @Summary 		Get items
// @Description 	Get all items
// @Tags 			models.Item
// @Accept  		json
// @Produce  		json
// @Success 		200 				{object} 	[]Item
// @Router 			/items 				[get]
func GetItems(c echo.Context) error {
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
func GetItemByID(c echo.Context) error {
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
func CreateItem(c echo.Context) error {
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
func UpdateItem(c echo.Context) error {
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
func DeleteItem(c echo.Context) error {
	panic("implement me")
}
