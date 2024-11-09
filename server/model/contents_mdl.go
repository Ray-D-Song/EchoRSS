package model

import "ray-d-song.com/echo-rss/db"

type Content struct {
	ID        int    `json:"id" db:"id"`
	UserID    int    `json:"userId" db:"user_id"`
	FolderID  int    `json:"folderId" db:"folder_id"`
	Title     string `json:"title" db:"title"`
	Content   string `json:"content" db:"content"`
	CreatedAt string `json:"createdAt" db:"created_at"`
	IsDeleted int    `json:"isDeleted" db:"is_deleted"`
	DeletedAt string `json:"deletedAt" db:"deleted_at"`
}

func (c *Content) Create() error {
	_, err := db.Bind.NamedExec(`
		INSERT INTO contents (title, content)
		VALUES (:title, :content)
	`, c)
	return err
}

func (c *Content) List(userID string) ([]Content, error) {
	contents := []Content{}
	err := db.Bind.Select(&contents, `SELECT id, user_id, folder_id, title, created_at FROM contents WHERE is_deleted = 0 AND user_id = ?`, userID)
	return contents, err
}

func (c *Content) Get(id int) error {
	return db.Bind.Get(c, `SELECT id, title, content, created_at, is_deleted, deleted_at FROM contents WHERE id = ?`, id)
}

func (c *Content) Update() error {
	_, err := db.Bind.NamedExec(`
		UPDATE contents SET title = :title, content = :content WHERE id = :id
	`, c)
	return err
}

func (c *Content) Delete() error {
	_, err := db.Bind.NamedExec(`
		UPDATE contents SET is_deleted = 1 WHERE id = :id
	`, c)
	return err
}

func (c *Content) Restore() error {
	_, err := db.Bind.NamedExec(`
		UPDATE contents SET is_deleted = 0 WHERE id = :id
	`, c)
	return err
}
