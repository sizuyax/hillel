package models

import "github.com/dgrijalva/jwt-go"

type AccessJWTClaims struct {
	jwt.StandardClaims
	SellerID int
}
