package entity

type Bid struct {
	ID      int     `json:"id"`
	ItemID  int     `json:"item_id"`
	OwnerID int     `json:"owner_id"`
	Points  float64 `json:"points"`
}
