package models

type Item struct {
	ID      int     `json:"id"`
	OwnerID int     `json:"owner_id" db:"owner_id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
}
