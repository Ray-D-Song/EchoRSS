package controller

import (
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"ray-d-song.com/echo-rss/utils"
)

func FetchRemoteContent(c *fiber.Ctx) error {
	url := c.Query("url")
	if url == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError("url is required"))
	}
	resp, err := http.Get(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	defer resp.Body.Close()
	// return the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	c.Set("Content-Type", "text/html")
	return c.Send(body)
}
