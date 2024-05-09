package httpmodels

type BaseSeller struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateSellerRequest struct {
	BaseSeller
}

type CreateSellerResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
