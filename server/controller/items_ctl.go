package controller

import (
	"github.com/gofiber/fiber/v2"
	"ray-d-song.com/echo-rss/model"
)

func GetItemsHdl(c *fiber.Ctx) error {
	userID := c.Locals("user").(string)
	feedID := c.Query("feedId")
	items, err := model.GetItems(userID, feedID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(items)
}
