package entity

import "github.com/dgrijalva/jwt-go"

type AccessJWTClaims struct {
	jwt.StandardClaims
	ProfileID   int    `json:"profileID"`
	ProfileType string `json:"profileType"`
}

type RefreshJWTClaims struct {
	jwt.StandardClaims
	SellerID int `json:"sellerID"`
}

type PairJWTClaims struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
