package model

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = os.Getenv("JWT_SECRET_KEY")

func init() {
	if SecretKey == "" {
		SecretKey = "echo-rss-secret-key"
	}
}

func GenerateAccessToken(userID string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 15).Unix(),
		"type":    "access",
	})

	accessTokenString, err := accessToken.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}
