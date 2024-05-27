package entity

type ProfileKey string

const (
	ProfileIDKey   ProfileKey = "ProfileID"
	ProfileTypeKey ProfileKey = "ProfileType"
)

type Seller struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     string `json:"type"`
}
