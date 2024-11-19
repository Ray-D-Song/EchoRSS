package controller

import (
	"github.com/gofiber/fiber/v2"
	"ray-d-song.com/echo-rss/model"
	"ray-d-song.com/echo-rss/utils"
)

func ListUsersHdl(c *fiber.Ctx) error {
	users := []model.User{}
	if err := model.ListUsers(&users); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.JSON(users)
}

func CreateUserHdl(c *fiber.Ctx) error {
	if !model.IsAdmin(c.Get("user_id")) {
		return c.Status(fiber.StatusForbidden).JSON(utils.LogError("unauthorized"))
	}
	user := model.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError(err.Error()))
	}
	if err := user.Create(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.JSON(user)
}

func DeleteUserHdl(c *fiber.Ctx) error {
	if !model.IsAdmin(c.Get("user_id")) {
		return c.Status(fiber.StatusForbidden).JSON(utils.LogError("unauthorized"))
	}
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError("id is required"))
	}
	if err := model.DeleteUser(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func RestoreUserHdl(c *fiber.Ctx) error {
	if !model.IsAdmin(c.Get("user_id")) {
		return c.Status(fiber.StatusForbidden).JSON(utils.LogError("unauthorized"))
	}
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError("id is required"))
	}
	if err := model.RestoreUser(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.SendStatus(fiber.StatusNoContent)
}
