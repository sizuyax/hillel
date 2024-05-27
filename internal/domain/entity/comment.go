package entity

type Comment struct {
	ID      int    `json:"id"`
	ItemID  int    `json:"item_id"`
	OwnerID int    `json:"owner_id"`
	Body    string `json:"body"`
}
