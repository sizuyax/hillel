package httpmodels

type BaseItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateItemRequest struct {
	OwnerID int     `json:"owner_id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
}

type CreateItemResponse struct {
	ID      int `json:"id"`
	OwnerID int `json:"owner_id"`
	BaseItem
}

type GetItemResponse struct {
	ID      int `json:"id"`
	OwnerID int `json:"owner_id"`
	BaseItem
}

type UpdateItemRequest struct {
	ID      int `json:"id"`
	OwnerID int `json:"owner_id"`
	BaseItem
}

type UpdateItemResponse struct {
	ID      int `json:"id"`
	OwnerID int `json:"owner_id"`
	BaseItem
}
