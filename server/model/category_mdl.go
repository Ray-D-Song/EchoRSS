package model

import (
	"database/sql"
	"errors"

	"ray-d-song.com/echo-rss/db"
)

type Category struct {
	ID        int    `json:"id" db:"id"`
	UserID    string `json:"-" db:"user_id"`
	Name      string `json:"name" db:"name"`
	CreatedAt string `json:"createdAt" db:"created_at"`
}

func (c *Category) List(userID string) ([]Category, error) {
	categories := []Category{}
	err := db.Bind.Select(&categories, "SELECT id, name FROM categories WHERE user_id = ?", userID)
	return categories, err
}

func (c *Category) Create() (int, error) {
	if c.UserID == "" {
		return 0, errors.New("userID is required")
	}
	if c.Name == "" {
		return 0, errors.New("name is required")
	}
	result, err := db.Bind.NamedExec("INSERT INTO categories (user_id, name) VALUES (:user_id, :name)", c)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func (c *Category) GetIDByName(userID, name string) (int, error) {
	if userID == "" {
		return 0, errors.New("userID is required")
	}
	if name == "" {
		return 0, errors.New("name is required")
	}
	var id int
	err := db.Bind.Get(&id, "SELECT id FROM categories WHERE user_id = ? AND name = ?", userID, name)
	if err != nil && err == sql.ErrNoRows {
		return 0, nil
	}
	return id, err
}

func (c *Category) Rename(userID, originalName, newName string) error {
	if userID == "" {
		return errors.New("userID is required")
	}
	if originalName == "" || newName == "" {
		return errors.New("originalName and newName are required")
	}
	_, err := db.Bind.Exec("UPDATE categories SET name = ? WHERE user_id = ? AND name = ?", newName, userID, originalName)
	return err
}

func (c *Category) Delete() error {
	_, err := db.Bind.NamedExec("DELETE FROM categories WHERE user_id = :user_id AND name = :name", c)
	return err
}
