package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmcdole/gofeed"
	"ray-d-song.com/echo-rss/model"
)

func ListFeedsHdl(c *fiber.Ctx) error {
	userID := c.Get("user")
	feeds, err := (&model.Feed{}).List(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(feeds)
}

type CreateFeedReq struct {
	FeedUrl string `json:"url"`
}

func CreateFeedHdl(c *fiber.Ctx) error {
	req := new(CreateFeedReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	fd := gofeed.NewParser()
	fdRes, err := fd.ParseURL(req.FeedUrl)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	feed := model.Feed{
		UserID:        c.Get("user"),
		Title:         fdRes.Title,
		Link:          fdRes.Link,
		Favicon:       "",
		Description:   "",
		LastBuildDate: fdRes.Published,
	}
	if fdRes.Image != nil {
		feed.Favicon = fdRes.Image.URL
	}
	if fdRes.Description != "" {
		feed.Description = fdRes.Description
	}
	err = feed.Create()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(feed)
}

func RefreshFeedsHdl(c *fiber.Ctx) error {
	userID := c.Get("user")
	feeds, err := (&model.Feed{}).List(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	for _, feed := range feeds {
		fd := gofeed.NewParser()
		fdRes, err := fd.ParseURL(feed.Link)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		for _, item := range fdRes.Items {
			// check if item already exists
			exists, err := (&model.Item{}).Exists(feed.ID, item.Link, userID)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			if exists {
				continue
			}
			newItem := model.Item{
				FeedID:      feed.ID,
				Title:       item.Title,
				Link:        item.Link,
				Description: "",
				Content:     item.Content,
				PubDate:     item.Published,
			}
			if item.Description != "" {
				newItem.Description = item.Description
			}
			err = newItem.Create()
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
		}
	}
	return c.JSON(feeds)
}
