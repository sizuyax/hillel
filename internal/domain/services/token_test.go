package services

import (
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
)

func TestGenerateJWTAccessToken(t *testing.T) {
	profileID := 1
	profileType := "seller"
	var log *slog.Logger

	ts := NewTokenService(log, "test", "test")

	accessToken, err := ts.GenerateJWTAccessToken(profileID, profileType)

	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)
}

func TestParseJWTAccessToken(t *testing.T) {
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTY4NTg3NDMsInByb2ZpbGVJRCI6MSwicHJvZmlsZVR5cGUiOiJzZWxsZXIifQ.x7NnP88AnaYf5cwfsn28mVcNsKRDHCYim1YucCnFe6A"
	var log *slog.Logger

	ts := NewTokenService(log, "test", "test")

	profileID, _, err := ts.ParseJWTAccessToken(accessToken)

	assert.NoError(t, err)
	assert.Equal(t, 1, profileID)
}

func TestGenerateJWTRefreshToken(t *testing.T) {
	profileID := 1
	profileType := "seller"
	var log *slog.Logger

	ts := NewTokenService(log, "test", "test")

	refreshToken, err := ts.GenerateJWTRefreshToken(profileID, profileType)

	assert.NoError(t, err)
	assert.NotEmpty(t, refreshToken)
}

func TestRefreshAccessJWTToken(t *testing.T) {
	refreshToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYzMzM2MTcsInNlbGxlcklEIjoxfQ.wE95gpKM6bJkWOCbx__atFQjSso5ODKEQgYcnSzj6To"
	var log *slog.Logger

	ts := NewTokenService(log, "test", "test")

	jwtPair, err := ts.RefreshAccessJWTToken(refreshToken)

	assert.NoError(t, err)
	assert.NotEmpty(t, jwtPair)
}

func TestGenerateJWTPairTokens(t *testing.T) {
	profileID := 1
	profileType := "seller"
	var log *slog.Logger

	ts := NewTokenService(log, "test", "test")

	jwtPair, err := ts.GenerateJWTPairTokens(profileID, profileType)

	assert.NoError(t, err)
	assert.NotEmpty(t, jwtPair)
}
