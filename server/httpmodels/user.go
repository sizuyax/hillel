package httpmodels

type BaseUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID int `json:"id"`
	BaseUser
}
