package httpmodels

type RefreshTokensRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokensResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
