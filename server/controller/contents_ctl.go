package controller

import (
	"github.com/gofiber/fiber/v2"
	"ray-d-song.com/echo-rss/model"
)

type NewContent struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func ListContentsHdl(c *fiber.Ctx) error {
	userID := c.Get("userId")
	contents, err := (&model.Content{}).List(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(contents)
}

func CreateContentHdl(c *fiber.Ctx) error {
	var newContent NewContent
	if err := c.BodyParser(&newContent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	content := model.Content{
		Title:   newContent.Title,
		Content: newContent.Content,
	}
	err := content.Create()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Content created successfully"})
}
