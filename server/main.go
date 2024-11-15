package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"ray-d-song.com/echo-rss/controller"
	"ray-d-song.com/echo-rss/db"
	"ray-d-song.com/echo-rss/middleware"
	"ray-d-song.com/echo-rss/model"
	"ray-d-song.com/echo-rss/utils"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func init() {
	err := utils.EnsureDir()
	if err != nil {
		panic(err)
	}
	utils.InitLogger()
}

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
	app.Use(middleware.AuthMdl)
	app.Use(middleware.LoggerMiddleware())
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/refresh-token", controller.RefreshToken)

	users := api.Group("/users")
	users.Get("/", controller.ListUsersHdl)
	users.Post("/", controller.CreateUserHdl)
	users.Delete("/", controller.DeleteUserHdl)
	users.Put("/restore", controller.RestoreUserHdl)

	feeds := api.Group("/feeds")
	feeds.Post("/", controller.CreateFeedHdl)
	feeds.Get("/", controller.ListFeedsHdl)
	feeds.Post("/refresh", controller.RefreshFeedsHdl)

	items := api.Group("/items")
	items.Get("/", controller.GetItemsHdl)
	items.Put("/read", controller.SetItemReadHdl)

	tools := api.Group("/tools")
	tools.Get("/fetch-remote-content", controller.FetchRemoteContent)

	app.Listen(":8080")
}
