package model

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"ray-d-song.com/echo-rss/db"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var SecretKey = os.Getenv("JWT_SECRET_KEY")
var RefreshKey = os.Getenv("JWT_REFRESH_KEY")

func init() {
	if SecretKey == "" {
		SecretKey = "echo-rss-secret-key"
	}
	if RefreshKey == "" {
		RefreshKey = "echo-rss-refresh-key"
	}
}

func GenerateTokenPair(userID string) (*TokenResponse, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
		"type":    "access",
	})

	accessTokenString, err := accessToken.SignedString([]byte(SecretKey))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
		"type":    "refresh",
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(RefreshKey))
	if err != nil {
		return nil, err
	}

	_, err = db.Bind.NamedExec(`
		UPDATE users SET refresh_token = :refresh_token WHERE id = :user_id
	`, map[string]interface{}{
		"refresh_token": refreshTokenString,
		"user_id":       userID,
	})
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func RefreshToken(c *fiber.Ctx, refreshToken string) error {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(RefreshKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid refresh token",
		})
	}

	var raw string
	userID := token.Claims.(jwt.MapClaims)["user_id"].(string)
	err = db.Bind.Get(&raw, `
		SELECT refresh_token FROM users WHERE id = :user_id
	`, map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	if raw != refreshToken {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid refresh token",
		})
	}

	newTokens, err := GenerateTokenPair(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Generate token failed",
		})
	}

	db.Bind.NamedExec(`
		UPDATE users SET refresh_token = :refresh_token WHERE id = :user_id
	`, map[string]interface{}{
		"refresh_token": newTokens.RefreshToken,
		"user_id":       userID,
	})

	return c.JSON(newTokens)
}
