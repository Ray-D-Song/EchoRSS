package controller

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mmcdole/gofeed"
	"go.uber.org/zap"
	"ray-d-song.com/echo-rss/model"
	"ray-d-song.com/echo-rss/utils"
)

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
		feed.UnreadCount += newItemsCount
		feed.TotalCount += newItemsCount
		feed.RecentUpdateCount = newItemsCount
		utils.Logger.Info("refresh feed", zap.String("feedID", feed.ID), zap.Int("unread_count", feed.UnreadCount), zap.Int("total_count", feed.TotalCount), zap.Int("recent_update_count", feed.RecentUpdateCount))
		err = feed.UpdateCount()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.LogError(err.Error()))
		}
	}
	return c.JSON(feeds)
}

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
