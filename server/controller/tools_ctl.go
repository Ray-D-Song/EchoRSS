package controller

import (
	"database/sql"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"ray-d-song.com/echo-rss/model"
	"ray-d-song.com/echo-rss/utils"
)

func FetchRemoteContent(c *fiber.Ctx) error {
	url := c.Query("url")
	if url == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError("url is required"))
	}
	c.Set("Content-Type", "text/html")
	cache := model.WebPageCache{Url: url}
	content, err := cache.Get()
	if err != nil && err != sql.ErrNoRows {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	if content != "" {
		return c.SendString(content)
	}
	resp, err := http.Get(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	defer resp.Body.Close()
	// return the body
	body, err := io.ReadAll(resp.Body)
	cache.Content = string(body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	_ = cache.Set()
	return c.Send(body)
}
