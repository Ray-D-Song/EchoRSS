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
	Deleted   int    `json:"deleted" db:"deleted"`
	CreatedAt string `json:"createdAt" db:"created_at"`
}

func NewUser() *User {
	return &User{
		ID:   uuid.New().String(),
		Role: "user",
	}
}

func (u *User) Create() error {
	u.ID = uuid.New().String()
	if u.Role == "" {
		u.Role = "user"
	}
	_, err := db.Bind.NamedExec(`
		INSERT INTO users (id, username, password, role)
		VALUES (:id, :username, :password, :role)
	`, u)
	return err
}

func (u *User) GetRoleByID(id string) error {
	return db.Bind.Get(u, `SELECT role FROM users WHERE id = ?`, id)
}

func (u *User) GetByUsername(username string) error {
	return db.Bind.Get(u, `SELECT id, username, password, role, created_at FROM users WHERE username = ? AND deleted = 0`, username)
}

func ListUsers(users *[]User) error {
	return db.Bind.Select(users, `SELECT id, username, password, role, deleted FROM users`)
}

func DeleteUser(id string) error {
	_, err := db.Bind.Exec(`UPDATE users SET deleted = 1 WHERE id = ?`, id)
	return err
}

func RestoreUser(id string) error {
	_, err := db.Bind.Exec(`UPDATE users SET deleted = 0 WHERE id = ?`, id)
	return err
}
