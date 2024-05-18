package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"project-auction/models"
	"time"
)

func GenerateJWTAccessToken(sellerID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.AccessJWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Hour).UTC().Unix(),
		},
		SellerID: sellerID,
	})

	return token.SignedString([]byte("test"))
}

func ParseJWTAccessToken(accessToken string) (int, error) {
	const op = "services.ParseJWTAccessToken"

	token, err := jwt.ParseWithClaims(accessToken, &models.AccessJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s: unexpected signing method: %v", op, token.Header["alg"])
		}

		return []byte("test"), nil
	})
	if err != nil {
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	claims, ok := token.Claims.(*models.AccessJWTClaims)
	if !ok {
		return 0, fmt.Errorf("%s: failed to parse claims: %v", op, err)
	}

	return claims.SellerID, nil
}
