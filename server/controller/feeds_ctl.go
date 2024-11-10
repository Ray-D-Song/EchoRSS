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

func CreateFeedHdl(c *fiber.Ctx) error {
	feedUrl := c.FormValue("feed_url")
	fd := gofeed.NewParser()
	feed, err := fd.ParseURL(feedUrl)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	feed := new(model.Feed)
	feed.UserID = c.Get("user")
	feed.Title = feed.Title
	feed.Link = feed.Link
	feed.Favicon = feed.Favicon
	if err := c.BodyParser(feed); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	feed.Create()
	return c.JSON(feed)
}
