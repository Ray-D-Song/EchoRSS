package main

import (
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"ray-d-song.com/echo-rss/controller"
	"ray-d-song.com/echo-rss/db"
	"ray-d-song.com/echo-rss/middleware"
	"ray-d-song.com/echo-rss/model"
	"ray-d-song.com/echo-rss/utils"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "ray-d-song.com/echo-rss/docs"
)

// @title Echo RSS API
// @version 1.0
// @description Echo RSS API
// @host localhost:8080
// @BasePath /api
func main() {
	utils.InitLogger()
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

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Use(middleware.AuthMdl)
	app.Use(middleware.LoggerMiddleware())
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)

	users := api.Group("/users")
	users.Get("/", controller.ListUsersHdl)
	users.Post("/", controller.CreateUserHdl)
	users.Delete("/", controller.DeleteUserHdl)
	users.Put("/restore", controller.RestoreUserHdl)
	users.Put("/", controller.UpdateUserHdl)

	userSetting := api.Group("/user/config")
	userSetting.Get("/", controller.GetUserSetting)
	userSetting.Put("/", controller.UpdateAiSetting)

	category := api.Group("/category")
	category.Get("/", controller.ListCategoriesHdl)
	category.Put("/rename", controller.RenameCategoryHdl)
	category.Delete("/", controller.DeleteCategoryHdl)

	feeds := api.Group("/feeds")
	feeds.Post("/", controller.CreateFeedHdl)
	feeds.Get("/", controller.ListFeedsHdl)
	feeds.Post("/refresh", controller.RefreshFeedsHdl)
	feeds.Delete("/", controller.DeleteFeedHdl)
	feeds.Put("/all-read", controller.MarkAllFeedsAsReadHdl)

	items := api.Group("/items")
	items.Get("/", controller.GetItemsHdl)
	items.Put("/read", controller.SetItemReadHdl)

	bookmark := api.Group("/bookmark")
	bookmark.Put("/", controller.ToggleItemBookmarkHdl)

	tools := api.Group("/tools")
	tools.Get("/fetch-remote-content", controller.FetchRemoteContent)

	utils.SetupStatic(app)

	app.Listen(":11299")
}
