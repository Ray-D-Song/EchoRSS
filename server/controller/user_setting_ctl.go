package controller

import (
	"github.com/gofiber/fiber/v2"
	"ray-d-song.com/echo-rss/model"
	"ray-d-song.com/echo-rss/utils"
)

func GetUserSetting(c *fiber.Ctx) error {
	userId := c.Locals("user").(string)
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError("user_id is required"))
	}
	userSetting := &model.UserSetting{UserId: userId}
	err := userSetting.GetUserSetting()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.JSON(userSetting)
}

func UpdateAiSetting(c *fiber.Ctx) error {
	userId := c.Locals("user").(string)
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError("user_id is required"))
	}
	userSetting := &model.UserSetting{UserId: userId}
	err := c.BodyParser(userSetting)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError(err.Error()))
	}
	err = userSetting.UpdateAiSetting()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.JSON(userSetting)
}
