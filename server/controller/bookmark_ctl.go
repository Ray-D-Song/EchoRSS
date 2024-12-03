package controller

import (
	"github.com/gofiber/fiber/v2"
	"ray-d-song.com/echo-rss/model"
	"ray-d-song.com/echo-rss/utils"
)

func ToggleItemBookmarkHdl(c *fiber.Ctx) error {
	userID := c.Locals("user").(string)
	itemID := c.Query("itemId")
	err := model.ToggleItemBookmark(userID, itemID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ok"})
}
