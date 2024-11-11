package controller

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mmcdole/gofeed"
	"ray-d-song.com/echo-rss/model"
)

func ListFeedsHdl(c *fiber.Ctx) error {
	userID := c.Locals("user").(string)
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
		UserID:        c.Locals("user").(string),
		Title:         fdRes.Title,
		Link:          req.FeedUrl,
		Favicon:       "",
		Description:   "",
		LastBuildDate: fdRes.Updated,
	}
	if fdRes.Image != nil {
		// get favicon from image url
		resp, err := http.Get(fdRes.Image.URL)
		if err == nil {
			defer resp.Body.Close()
			// Read response body into bytes
			body, err := io.ReadAll(resp.Body)
			if err == nil {
				feed.Favicon = base64.StdEncoding.EncodeToString(body)
			}
		}
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
	userID := c.Locals("user").(string)
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
				UserID:      userID,
				Title:       item.Title,
				Link:        item.Link,
				Description: "",
				Content:     "",
				PubDate:     item.Updated,
			}
			if item.Description != "" {
				newItem.Description = item.Description
			}
			if item.Content != "" {
				newItem.Content = item.Content
			}
			err = newItem.Create()
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
		}
	}
	return c.JSON(feeds)
}
