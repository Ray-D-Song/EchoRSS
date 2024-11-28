package controller

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mmcdole/gofeed"
	"ray-d-song.com/echo-rss/model"
	"ray-d-song.com/echo-rss/utils"
)

// ListFeedsHdl godoc
// @Summary List all feeds
// @Description Get a list of all feeds
// @Tags feeds
// @Produce json
// @Success 200 {array} model.Feed
// @Failure 500 {object} utils.ErrRes
func ListFeedsHdl(c *fiber.Ctx) error {
	userID := c.Locals("user").(string)
	feeds, err := (&model.Feed{}).List(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.JSON(feeds)
}

type CreateFeedReq struct {
	FeedUrl      string `json:"url"`
	CategoryName string `json:"category"`
}

// CreateFeedHdl godoc
// @Summary Create a new feed
// @Description Create a new feed with the provided details
// @Tags feeds
// @Param feed body CreateFeedReq true "Feed details"
// @Success 200 {object} model.Feed
// @Failure 400 {object} utils.ErrRes
// @Failure 500 {object} utils.ErrRes
// @Router /feeds [post]
func CreateFeedHdl(c *fiber.Ctx) error {
	req := new(CreateFeedReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError(err.Error()))
	}
	userID := c.Locals("user").(string)
	categoryID, err := (&model.Category{}).GetIDByName(userID, req.CategoryName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	if categoryID == 0 {
		newCategory := model.Category{
			UserID: userID,
			Name:   req.CategoryName,
		}
		newCategoryID, err := newCategory.Create()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
		}
		categoryID = newCategoryID
	}

	fd := gofeed.NewParser()
	fdRes, err := fd.ParseURL(req.FeedUrl)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError(err.Error()))
	}
	feed := model.Feed{
		UserID:        c.Locals("user").(string),
		Title:         fdRes.Title,
		Link:          req.FeedUrl,
		Favicon:       "",
		Description:   "",
		LastBuildDate: fdRes.Updated,
		CategoryID:    categoryID,
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
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.JSON(feed)
}

// RefreshFeedsHdl godoc
// @Summary Refresh feeds
// @Description Refresh feeds for the user
// @Tags feeds
// @Success 200 {array} model.Feed
// @Failure 500 {object} utils.ErrRes
func RefreshFeedsHdl(c *fiber.Ctx) error {
	userID := c.Locals("user").(string)
	feeds, err := (&model.Feed{}).List(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	for _, feed := range feeds {
		fd := gofeed.NewParser()
		fdRes, err := fd.ParseURL(feed.Link)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
		}
		newItemsCount := 0
		for _, item := range fdRes.Items {
			// check if item already exists
			exists, err := (&model.Item{}).Exists(feed.ID, item.Link, userID)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
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
				Read:        0,
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
				return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
			}
			newItemsCount++
		}
		feed.RecentUpdateCount = newItemsCount
		err = feed.UpdateCount()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
		}
	}
	return c.JSON(feeds)
}

// DeleteFeedHdl godoc
// @Summary Delete a feed
// @Description Delete a feed by ID
// @Tags feeds
// @Param feedID query string true "Feed ID"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} utils.ErrRes
// @Failure 500 {object} utils.ErrRes
func DeleteFeedHdl(c *fiber.Ctx) error {
	feedID := c.Query("feedID")
	if feedID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError("feedID is required"))
	}
	userID := c.Locals("user").(string)
	err := (&model.Feed{}).Delete(feedID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.JSON(fiber.Map{"success": true})
}

// MarkAllFeedsAsReadHdl godoc
// @Summary Mark all feeds as read
// @Description Mark all feeds as read for the user
// @Tags feeds
// @Success 200 {object} fiber.Map
// @Failure 500 {object} utils.ErrRes
func MarkAllFeedsAsReadHdl(c *fiber.Ctx) error {
	userID := c.Locals("user").(string)
	feedID := c.Query("feedID")
	if feedID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.LogError("feedID is required"))
	}
	err := (&model.Feed{}).MarkAllAsRead(userID, feedID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
	}
	return c.JSON(fiber.Map{"success": true})
}
