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
	_, err := db.Bind.Exec("UPDATE items SET read = 1 WHERE id = ? AND user_id = ?", itemID, userID)
	return err
}
