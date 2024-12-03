package controller

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"ray-d-song.com/echo-rss/model"
	"ray-d-song.com/echo-rss/utils"
)

func TranslateHdl(c *fiber.Ctx) error {
	userId := c.Locals("user").(string)
	if userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.LogError("user id is required"))
	}

	content := c.FormValue("content")
	url := c.FormValue("url")
	if url == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError("url is required"))
	}
	if content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError("content is required"))
	}
	// query cache first
	cache := model.TranslateCache{
		Url: url,
	}
	translated, err := cache.Get()
	if err != nil && err != sql.ErrNoRows {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	c.Set("Content-Type", "text/plain")
	if translated != "" {
		utils.Logger.Info("translate cache hit", zap.String("url", url))
		return c.SendString(translated)
	}
	userSetting := model.UserSetting{
		UserId: userId,
	}
	if err := userSetting.GetUserSetting(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	// translate and cache
	translated, err = utils.Translate(content, userSetting.OPENAI_API_KEY, userSetting.API_ENDPOINT, userSetting.TARGET_LANGUAGE)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	cache.Content = translated
	if err := cache.Set(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.SendString(translated)
}
