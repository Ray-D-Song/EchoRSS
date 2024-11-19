package controller

import (
	"github.com/gofiber/fiber/v2"
	"ray-d-song.com/echo-rss/model"
	"ray-d-song.com/echo-rss/utils"
)

func ListCategoriesHdl(c *fiber.Ctx) error {
	userID := c.Locals("user").(string)
	categories, err := (&model.Category{}).List(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.JSON(categories)
}

func RenameCategoryHdl(c *fiber.Ctx) error {
	userID := c.Locals("user").(string)
	originalName := c.Query("originalName")
	newName := c.Query("newName")
	if originalName == "" || newName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError("originalName and newName are required"))
	}
	err := (&model.Category{}).Rename(userID, originalName, newName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ok"})
}

func DeleteCategoryHdl(c *fiber.Ctx) error {
	userID := c.Locals("user").(string)
	name := c.Query("name")
	if name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError("name is required"))
	}

	category := &model.Category{
		UserID: userID,
		Name:   name,
	}

	err := category.Delete()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ok"})
}
