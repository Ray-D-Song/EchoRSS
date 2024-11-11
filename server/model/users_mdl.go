package model

import (
	"github.com/google/uuid"
	"ray-d-song.com/echo-rss/db"
)

type User struct {
	ID        string `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	Role      string `json:"role" db:"role"`
	CreatedAt string `json:"createdAt" db:"created_at"`
}

func NewUser() *User {
	return &User{
		ID:   uuid.New().String(),
		Role: "user",
	}
}

func (u *User) Create() error {
	_, err := db.Bind.NamedExec(`
		INSERT INTO users (id, username, password, role)
		VALUES (:id, :username, :password, :role)
	`, u)
	return err
}

func (u *User) Get(username string) error {
	return db.Bind.Get(u, `SELECT id, username, password, role, created_at FROM users WHERE username = ?`, username)
}

func (u *User) GetByUsername(username string) error {
	return db.Bind.Get(u, `SELECT id, username, password, role, created_at FROM users WHERE username = ?`, username)
}
