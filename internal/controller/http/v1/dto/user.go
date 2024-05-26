package dto

type BaseUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	BaseUser
}

type CreateUserResponse struct {
	ID int `json:"id"`
	BaseUser
}
