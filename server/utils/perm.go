package utils

import (
	"github.com/gofiber/fiber/v2"
	"ray-d-song.com/echo-rss/model"
)

func IsAdmin(c *fiber.Ctx) bool {
	userID := c.Locals("user").(string)
	var user model.User
	if err := user.GetRoleByID(userID); err != nil {
		return false
	}
	return user.Role == "admin"
}
