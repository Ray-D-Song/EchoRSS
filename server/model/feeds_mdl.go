package model

import (
	"errors"
	"time"

	"ray-d-song.com/echo-rss/db"
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

// read items count from db and update to feed
func (f *Feed) UpdateCount() error {
	if f.ID == "" {
		return errors.New("id is required")
	}
	count, err := (&Item{}).CountByFeedID(f.ID)
	if err != nil {
		return err
	}
	unreadCount, err := (&Item{}).CountUnreadByFeedID(f.ID)
	if err != nil {
		return err
	}
	f.TotalCount = count
	f.UnreadCount = unreadCount
	if f.RecentUpdateCount > 0 {
		_, err = db.Bind.NamedExec("UPDATE feeds SET recent_update_count = :recent_update_count, total_count = :total_count, unread_count = :unread_count WHERE id = :id", f)
	} else {
		_, err = db.Bind.NamedExec("UPDATE feeds SET total_count = :total_count, unread_count = :unread_count WHERE id = :id", f)
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

func (f *Feed) MarkAllAsRead(userID string, feedID string) error {
	// start transaction
	tx, err := db.Bind.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE items SET read = 1 WHERE user_id = ? AND feed_id = ?", userID, feedID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE feeds SET unread_count = 0 WHERE id = ?", feedID)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
