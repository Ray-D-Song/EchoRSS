package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"ray-d-song.com/echo-rss/controller"
	"ray-d-song.com/echo-rss/db"
	"ray-d-song.com/echo-rss/model"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"

	jwtware "github.com/gofiber/contrib/jwt"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	err := db.Migrate()
	if err != nil {
		panic(err)
	}

	// if admin user not exists, create it
	admin := new(model.User)
	if err := admin.GetByUsername("admin"); err != nil {
		admin.ID = uuid.New().String()
		admin.Username = "admin"
		admin.Password = "admin"
		admin.Role = "admin"
		admin.Create()
	}

	app := fiber.New()
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/refresh-token", controller.RefreshToken)

	api.Use(jwtware.New(jwtware.Config{
		Filter: func(c *fiber.Ctx) bool {
			return c.Path() == "/api/auth/login" || c.Path() == "/api/auth/refresh-token"
		},
		SigningKey:   jwtware.SigningKey{Key: []byte(model.SecretKey)},
		ErrorHandler: jwtError,
	}))

	app.Listen(":8080")
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}
