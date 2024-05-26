package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"project-auction/internal/domain/entity"
	"time"
)

func GenerateJWTAccessToken(sellerID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &entity.AccessJWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).UTC().Unix(),
		},
		SellerID: sellerID,
	})

	return token.SignedString([]byte("test"))
}

func ParseJWTAccessToken(accessToken string) (int, error) {
	const op = "services.ParseJWTAccessToken"

	token, err := jwt.ParseWithClaims(accessToken, &entity.AccessJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s: unexpected signing method: %v", op, token.Header["alg"])
		}

		return []byte("test"), nil
	})
	if err != nil {
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	claims, ok := token.Claims.(*entity.AccessJWTClaims)
	if !ok {
		return 0, fmt.Errorf("%s: failed to parse claims: %v", op, err)
	}

	return claims.SellerID, nil
}

func GenerateJWTRefreshToken(sellerID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &entity.AccessJWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Hour).UTC().Unix(),
		},
		SellerID: sellerID,
	})

	return token.SignedString([]byte("test-1"))
}

func RefreshAccessJWTToken(refreshToken string) (*entity.PairJWTClaims, error) {
	const op = "services.RefreshAccessJWTToken"

	token, err := jwt.ParseWithClaims(refreshToken, &entity.AccessJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s: unexpected signing method: %v", op, token.Header["alg"])
		}

		return []byte("test-1"), nil
	})
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	claims, ok := token.Claims.(*entity.AccessJWTClaims)
	if !ok {
		return nil, fmt.Errorf("%s: failed to parse claims: %v", op, err)
	}

	return GenerateJWTPairTokens(claims.SellerID)
}

func GenerateJWTPairTokens(sellerID int) (*entity.PairJWTClaims, error) {
	accessToken, err := GenerateJWTAccessToken(sellerID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateJWTRefreshToken(sellerID)
	if err != nil {
		return nil, err
	}

	jwtPairs := &entity.PairJWTClaims{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return jwtPairs, nil
}
