package model

import (
	"errors"
	"time"

	"go.uber.org/zap"
	"ray-d-song.com/echo-rss/db"
	"ray-d-song.com/echo-rss/utils"
)

type Feed struct {
	ID                string `json:"id" db:"id"`
	UserID            string `json:"-" db:"user_id"`
	Title             string `json:"title" db:"title"`
	Link              string `json:"link" db:"link"`
	Favicon           string `json:"favicon" db:"favicon"`
	Description       string `json:"description" db:"description"`
	LastBuildDate     string `json:"lastBuildDate" db:"last_build_date"`
	CategoryID        int    `json:"categoryID" db:"category_id"`
	UnreadCount       int    `json:"unreadCount" db:"unread_count"`
	TotalCount        int    `json:"totalCount" db:"total_count"`
	RecentUpdateCount int    `json:"recentUpdateCount" db:"recent_update_count"`
	CreatedAt         string `json:"createdAt" db:"created_at"`
}

func (f *Feed) List(userID string) ([]Feed, error) {
	feeds := []Feed{}
	err := db.Bind.Select(&feeds, "SELECT id, title, link, favicon, description, last_build_date, category_id, unread_count, total_count, recent_update_count, created_at FROM feeds WHERE user_id = ?", userID)
	return feeds, err
}

func (f *Feed) Create() error {
	if f.UserID == "" {
		return errors.New("userID is required")
	}
	if f.Title == "" {
		return errors.New("title is required")
	}
	if f.Link == "" {
		return errors.New("link is required")
	}
	if f.Description == "" {
		return errors.New("description is required")
	}
	if f.LastBuildDate == "" {
		f.LastBuildDate = time.Now().Format(time.RFC3339)
	}
	if f.CategoryID == 0 {
		return errors.New("categoryID is required")
	}
	_, err := db.Bind.NamedExec("INSERT INTO feeds (user_id, title, link, favicon, description, last_build_date, category_id) VALUES (:user_id, :title, :link, :favicon, :description, :last_build_date, :category_id)", f)
	return err
}

func (f *Feed) UpdateCount() error {
	if f.ID == "" {
		return errors.New("id is required")
	}
	_, err := db.Bind.NamedExec("UPDATE feeds SET unread_count = :unread_count, total_count = :total_count, recent_update_count = :recent_update_count WHERE id = :id", f)
	if err != nil {
		utils.Logger.Error("update feed count", zap.String("feedID", f.ID), zap.Int("unread_count", f.UnreadCount), zap.Int("total_count", f.TotalCount), zap.Int("recent_update_count", f.RecentUpdateCount), zap.Error(err))
	}
	return err
}

func (f *Feed) Delete(feedID, userID string) error {
	if feedID == "" {
		return errors.New("feedID is required")
	}
	if userID == "" {
		return errors.New("userID is required")
	}
	// start transaction
	tx, err := db.Bind.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM feeds WHERE id = ? AND user_id = ?", feedID, userID)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("DELETE FROM items WHERE feed_id = ? AND user_id = ?", feedID, userID)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
