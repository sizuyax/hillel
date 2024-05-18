package models

const SellerIDKey = "seller_id"

type Seller struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
