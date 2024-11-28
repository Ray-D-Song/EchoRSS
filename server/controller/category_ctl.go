package controller

import (
	"github.com/gofiber/fiber/v2"
	"ray-d-song.com/echo-rss/model"
	"ray-d-song.com/echo-rss/utils"
)

// ListCategoriesHdl godoc
// @Summary List all categories
// @Description Get a list of all categories
// @Tags categories
// @Produce json
// @Success 200 {array} model.Category
// @Failure 500 {object} utils.ErrRes
func ListCategoriesHdl(c *fiber.Ctx) error {
	userID := c.Locals("user").(string)
	categories, err := (&model.Category{}).List(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.JSON(categories)
}

// RenameCategoryHdl godoc
// @Summary Rename a category
// @Description Rename a category by original name and new name
// @Tags categories
// @Param originalName query string true "Original name"
// @Param newName query string true "New name"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} utils.ErrRes
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

// DeleteCategoryHdl godoc
// @Summary Delete a category
// @Description Delete a category by name
// @Tags categories
// @Param name query string true "Category name"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} utils.ErrRes
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
