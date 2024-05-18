package models

import "github.com/dgrijalva/jwt-go"

type AccessJWTClaims struct {
	jwt.StandardClaims
	SellerID int `json:"sellerID"`
}

type RefreshJWTClaims struct {
	jwt.StandardClaims
	SellerID int `json:"sellerID"`
}

type PairJWTClaims struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
