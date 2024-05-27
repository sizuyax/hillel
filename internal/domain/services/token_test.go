package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateJWTAccessToken(t *testing.T) {
	profileID := 1
	profileType := "seller"

	accessToken, err := GenerateJWTAccessToken(profileID, profileType)

	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)

	fmt.Println(accessToken, "---> error: ", err)
}

func TestParseJWTAccessToken(t *testing.T) {
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTY4NTg3NDMsInByb2ZpbGVJRCI6MSwicHJvZmlsZVR5cGUiOiJzZWxsZXIifQ.x7NnP88AnaYf5cwfsn28mVcNsKRDHCYim1YucCnFe6A"

	profileID, profileType, err := ParseJWTAccessToken(accessToken)

	assert.NoError(t, err)
	assert.Equal(t, 1, profileID)

	fmt.Println(profileID)
	fmt.Println(profileType, "---> error: ", err)
}

func TestGenerateJWTRefreshToken(t *testing.T) {
	profileID := 1
	profileType := "seller"

	refreshToken, err := GenerateJWTRefreshToken(profileID, profileType)

	assert.NoError(t, err)
	assert.NotEmpty(t, refreshToken)

	fmt.Println("refresh token: ", refreshToken, "---> error: ", err)
}

func TestRefreshAccessJWTToken(t *testing.T) {
	refreshToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYzMzM2MTcsInNlbGxlcklEIjoxfQ.wE95gpKM6bJkWOCbx__atFQjSso5ODKEQgYcnSzj6To"

	jwtPair, err := RefreshAccessJWTToken(refreshToken)

	assert.NoError(t, err)
	assert.NotEmpty(t, jwtPair)

	fmt.Println("access token: ", jwtPair.AccessToken)
	fmt.Println("refresh token: ", jwtPair.RefreshToken, "---> error: ", err)
}

func TestGenerateJWTPairTokens(t *testing.T) {
	profileID := 1
	profileType := "seller"

	jwtPair, err := GenerateJWTPairTokens(profileID, profileType)

	assert.NoError(t, err)
	assert.NotEmpty(t, jwtPair)

	fmt.Println("access token: ", jwtPair.AccessToken)
	fmt.Println("refresh token: ", jwtPair.RefreshToken, "---> error: ", err)
}
