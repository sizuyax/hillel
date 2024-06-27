package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log/slog"
	"project-auction/internal/domain/entity"
	"time"
)

type TokenService interface {
	GenerateJWTAccessToken(profileID int, profileType string) (string, error)
	ParseJWTAccessToken(accessToken string) (int, string, error)
	GenerateJWTRefreshToken(profileID int, profileType string) (string, error)
	RefreshAccessJWTToken(refreshToken string) (*entity.PairJWTClaims, error)
	GenerateJWTPairTokens(profileID int, profileType string) (*entity.PairJWTClaims, error)
}

type tokenService struct {
	log                 *slog.Logger
	accessSignedString  string
	refreshSignedString string
}

func NewTokenService(log *slog.Logger, accessSignedString, refreshSignedString string) TokenService {
	return &tokenService{
		log:                 log,
		accessSignedString:  accessSignedString,
		refreshSignedString: refreshSignedString,
	}
}

func (ts tokenService) GenerateJWTAccessToken(profileID int, profileType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &entity.AccessJWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Hour).UTC().Unix(),
		},
		ProfileID:   profileID,
		ProfileType: profileType,
	})

	return token.SignedString([]byte(ts.accessSignedString))
}

func (ts tokenService) ParseJWTAccessToken(accessToken string) (int, string, error) {
	const op = "services.ParseJWTAccessToken"

	token, err := jwt.ParseWithClaims(accessToken, &entity.AccessJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			ts.log.Error("unexpected signing method", slog.String("op", op))
			return nil, fmt.Errorf("%s: unexpected signing method: %v", op, token.Header["alg"])
		}

		return []byte(ts.accessSignedString), nil
	})
	if err != nil {
		ts.log.Error("failed to parse access token", slog.String("error", err.Error()), slog.String("op", op))
		return 0, "", fmt.Errorf("%s:%v", op, err)
	}

	claims, ok := token.Claims.(*entity.AccessJWTClaims)
	if !ok {
		ts.log.Error("failed to get claims", slog.String("op", op))
		return 0, "", fmt.Errorf("%s: failed to parse claims: %v", op, err)
	}

	return claims.ProfileID, claims.ProfileType, nil
}

func (ts tokenService) GenerateJWTRefreshToken(profileID int, profileType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &entity.AccessJWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Hour).UTC().Unix(),
		},
		ProfileID:   profileID,
		ProfileType: profileType,
	})

	return token.SignedString([]byte(ts.refreshSignedString))
}

func (ts tokenService) RefreshAccessJWTToken(refreshToken string) (*entity.PairJWTClaims, error) {
	const op = "services.RefreshAccessJWTToken"

	token, err := jwt.ParseWithClaims(refreshToken, &entity.AccessJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			ts.log.Error("unexpected signing method", slog.String("op", op))
			return nil, fmt.Errorf("%s: unexpected signing method: %v", op, token.Header["alg"])
		}

		return []byte(ts.refreshSignedString), nil
	})
	if err != nil {
		ts.log.Error("failed to parse refresh token", slog.String("error", err.Error()), slog.String("op", op))
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	claims, ok := token.Claims.(*entity.AccessJWTClaims)
	if !ok {
		ts.log.Error("failed to get claims", slog.String("op", op))
		return nil, fmt.Errorf("%s: failed to parse claims: %v", op, err)
	}

	return ts.GenerateJWTPairTokens(claims.ProfileID, claims.ProfileType)
}

func (ts tokenService) GenerateJWTPairTokens(profileID int, profileType string) (*entity.PairJWTClaims, error) {
	accessToken, err := ts.GenerateJWTAccessToken(profileID, profileType)
	if err != nil {
		return nil, err
	}

	refreshToken, err := ts.GenerateJWTRefreshToken(profileID, profileType)
	if err != nil {
		return nil, err
	}

	jwtPairs := &entity.PairJWTClaims{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return jwtPairs, nil
}
