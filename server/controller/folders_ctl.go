package controller

import (
	"github.com/gofiber/fiber/v2"
	"ray-d-song.com/echo-rss/model"
)

func ListFoldersHdl(c *fiber.Ctx) error {
	userID := c.Get("userId")
	folders, err := (&model.Folder{}).List(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(folders)
}

func CreateFolderHdl(c *fiber.Ctx) error {
	userID := c.Get("userId")
	folder := new(model.Folder)
	folder.UserID = userID
	folder.Name = c.Query("name")
	folder.ParentID = c.QueryInt("parentId", 0)
	if err := folder.Create(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(folder)
}

func DeleteFolderHdl(c *fiber.Ctx) error {
	folderID := c.QueryInt("id")
	if folderID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Folder ID is required"})
	}
	folder := new(model.Folder)
	folder.ID = folderID
	if err := folder.Delete(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Folder deleted"})
}

func RenameFolderHdl(c *fiber.Ctx) error {
	folderID := c.QueryInt("id")
	if folderID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Folder ID is required"})
	}
	folder := new(model.Folder)
	folder.ID = folderID
	folder.Name = c.Query("name")
	if err := folder.Update(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(folder)
}
