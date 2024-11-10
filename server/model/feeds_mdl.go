package model

import (
	"ray-d-song.com/echo-rss/db"
)

type Feed struct {
	ID        string `json:"id" db:"id"`
	UserID    string `json:"user_id" db:"user_id"`
	Title     string `json:"title" db:"title"`
	Link      string `json:"link" db:"link"`
	Favicon   string `json:"favicon" db:"favicon"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

func (f *Feed) List(userID string) ([]Feed, error) {
	feeds := []Feed{}
	db.Bind.Select(&feeds, "SELECT * FROM feeds WHERE user_id = ?", userID)
	return feeds, nil
}

func (f *Feed) Create() error {
	_, err := db.Bind.NamedExec("INSERT INTO feeds (user_id, title, link, favicon) VALUES (:user_id, :title, :link, :favicon)", f)
	return err
}
