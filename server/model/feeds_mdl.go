package model

import (
	"ray-d-song.com/echo-rss/db"
)

type Feed struct {
	ID            string `json:"id" db:"id"`
	UserID        string `json:"-" db:"user_id"`
	Title         string `json:"title" db:"title"`
	Link          string `json:"link" db:"link"`
	Favicon       string `json:"favicon" db:"favicon"`
	Description   string `json:"description" db:"description"`
	LastBuildDate string `json:"lastBuildDate" db:"last_build_date"`
	CreatedAt     string `json:"createdAt" db:"created_at"`
}

func (f *Feed) List(userID string) ([]Feed, error) {
	feeds := []Feed{}
	db.Bind.Select(&feeds, "SELECT id, title, link, favicon, description, last_build_date, created_at FROM feeds WHERE user_id = ?", userID)
	return feeds, nil
}

func (f *Feed) Create() error {
	_, err := db.Bind.NamedExec("INSERT INTO feeds (user_id, title, link, favicon, description, last_build_date) VALUES (:user_id, :title, :link, :favicon, :description, :last_build_date)", f)
	return err
}
