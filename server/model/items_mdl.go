package model

import (
	"ray-d-song.com/echo-rss/db"
)

type Item struct {
	ID          string `json:"id" db:"id"`
	FeedID      string `json:"feedId" db:"feed_id"`
	UserID      string `json:"-" db:"user_id"`
	Title       string `json:"title" db:"title"`
	Link        string `json:"link" db:"link"`
	Description string `json:"description" db:"description"`
	Content     string `json:"content" db:"content"`
	Read        int    `json:"read" db:"read"`
	PubDate     string `json:"pubDate" db:"pub_date"`
	CreatedAt   string `json:"createdAt" db:"created_at"`
}

func (i *Item) CountUnread(feedID string) (int, error) {
	var count int
	err := db.Bind.Get(&count, "SELECT COUNT(*) FROM items WHERE feed_id = ? AND read = 0", feedID)
	if err == nil {
		// update unread count
		db.Bind.Exec("UPDATE feeds SET unread_count = ? WHERE id = ?", count, feedID)
	}
	return count, err
}

func (i *Item) Exists(feedID, link, userID string) (bool, error) {
	var count int
	err := db.Bind.Get(&count, "SELECT COUNT(*) FROM items WHERE feed_id = ? AND link = ? AND user_id = ?", feedID, link, userID)
	return count > 0, err
}

func (i *Item) Create() error {
	_, err := db.Bind.NamedExec("INSERT INTO items (feed_id, user_id, title, link, description, pub_date, content) VALUES (:feed_id, :user_id, :title, :link, :description, :pub_date, :content)", i)
	return err
}

func GetItems(userID, feedID string) ([]Item, error) {
	var items []Item
	err := db.Bind.Select(&items, "SELECT * FROM items WHERE feed_id = ? AND user_id = ?", feedID, userID)
	return items, err
}

func SetItemRead(userID, itemID string) error {
	tx, err := db.Bind.Beginx()
	if err != nil {
		return err
	}
	var feedID string
	err = tx.Get(&feedID, "SELECT feed_id FROM items WHERE id = ? AND user_id = ?", itemID, userID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE items SET read = 1 WHERE id = ? AND user_id = ?", itemID, userID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE feeds SET unread_count = unread_count - 1 WHERE id = ?", feedID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
