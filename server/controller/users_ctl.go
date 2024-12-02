package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"ray-d-song.com/echo-rss/model"
	"ray-d-song.com/echo-rss/utils"
)

// ListUsersHdl godoc
// @Summary List all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} model.User
// @Failure 500 {object} utils.ErrRes
// @Router /users [get]
func ListUsersHdl(c *fiber.Ctx) error {
	users := []model.User{}
	if err := model.ListUsers(&users); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.JSON(users)
}

// CreateUserHdl godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User details"
// @Success 200 {object} model.User
// @Failure 400 {object} utils.ErrRes
// @Failure 500 {object} utils.ErrRes
// @Router /users [post]
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

// DeleteUserHdl godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Param id query string true "User ID"
// @Success 204
// @Failure 403 {object} utils.ErrRes
// @Failure 400 {object} utils.ErrRes
// @Failure 500 {object} utils.ErrRes
// @Router /users [delete]
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

// RestoreUserHdl godoc
// @Summary Restore a user
// @Description Restore a user by ID
// @Tags users
// @Param id query string true "User ID"
// @Success 204
// @Failure 403 {object} utils.ErrRes
// @Failure 400 {object} utils.ErrRes
// @Failure 500 {object} utils.ErrRes
// @Router /users/restore [post]
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

// UpdateUserHdl godoc
// @Summary Update a user
// @Description Update a user's username and password
// @Tags users
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Param user body model.User true "User details"
// @Success 204
// @Failure 403 {object} utils.ErrRes
// @Failure 400 {object} utils.ErrRes
// @Failure 500 {object} utils.ErrRes
// @Router /users [put]
func UpdateUserHdl(c *fiber.Ctx) error {
	operatorID := c.Locals("user").(string)

	if operatorID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError("id is required"))
	}
	if !model.IsAdmin(operatorID) {
		return c.Status(fiber.StatusForbidden).JSON(utils.LogError("unauthorized"))
	}

	user := model.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError(err.Error()))
	}

	utils.Logger.Info("update user", zap.Any("user", user))
	if err := user.Update(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}

	return c.SendStatus(fiber.StatusNoContent)
}
