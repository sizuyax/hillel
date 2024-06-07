package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateJWTAccessToken(t *testing.T) {
	sellerID := 1

	accessToken, err := GenerateJWTAccessToken(sellerID)

	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)

	fmt.Println(accessToken, "---> error: ", err)
}

func TestParseJWTAccessToken(t *testing.T) {
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYyOTc5MDYsInNlbGxlcklEIjoxfQ.mbD_3q13p3Qw7vwF_CatfDWl61ru0x-iJnRA9bw3VH8"

	sellerID, err := ParseJWTAccessToken(accessToken)

	assert.NoError(t, err)
	assert.Equal(t, 1, sellerID)

	fmt.Println(sellerID, "---> error: ", err)
}

func TestGenerateJWTRefreshToken(t *testing.T) {
	sellerID := 1

	refreshToken, err := GenerateJWTRefreshToken(sellerID)

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
	sellerID := 1

	jwtPair, err := GenerateJWTPairTokens(sellerID)

	assert.NoError(t, err)
	assert.NotEmpty(t, jwtPair)

	fmt.Println("access token: ", jwtPair.AccessToken)
	fmt.Println("refresh token: ", jwtPair.RefreshToken, "---> error: ", err)
}
