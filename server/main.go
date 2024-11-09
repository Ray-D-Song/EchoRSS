package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

	api.Post("/login", controller.Login)
	api.Post("/refresh-token", controller.RefreshToken)

	api.Use(jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(model.SecretKey)},
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
	}))

	api.Get("/folders", controller.ListFoldersHdl)
	api.Post("/folders", controller.CreateFolderHdl)
	api.Delete("/folders", controller.DeleteFolderHdl)
	api.Put("/folders", controller.RenameFolderHdl)
	api.Get("/contents", controller.ListContentsHdl)
	api.Post("/contents", controller.CreateContentHdl)

	app.Listen(":8080")
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

func jwtSuccess(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if userID, ok := claims["user_id"]; ok {
		c.Set("userId", userID.(string))
		return c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Invalid token claims",
	})
}
