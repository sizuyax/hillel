package models

type SellString string

const SellerIDKey SellString = "seller_id"

type Seller struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
