package model

import "ray-d-song.com/echo-rss/db"

type Item struct {
	ID          string `json:"id" db:"id"`
	FeedID      string `json:"feed_id" db:"feed_id"`
	Title       string `json:"title" db:"title"`
	Link        string `json:"link" db:"link"`
	Description string `json:"description" db:"description"`
	Content     string `json:"content" db:"content"`
	PubDate     string `json:"pub_date" db:"pub_date"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}

func (i *Item) Exists(feedID, link, userID string) (bool, error) {
	var count int
	err := db.Bind.Get(&count, "SELECT COUNT(*) FROM items WHERE feed_id = ? AND link = ? AND user_id = ?", feedID, link, userID)
	return count > 0, err
}

func (i *Item) Create() error {
	_, err := db.Bind.NamedExec("INSERT INTO items (feed_id, title, link, description, pub_date) VALUES (:feed_id, :title, :link, :description, :pub_date)", i)
	return err
}
