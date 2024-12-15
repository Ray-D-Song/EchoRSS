//go:build prod

package utils

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed dist/*
var DistFiles embed.FS

func SetupStatic(app *fiber.App) {
	app.Use("/*", filesystem.New(filesystem.Config{
		Root:       http.FS(DistFiles),
		PathPrefix: "/dist",
	}))
}
