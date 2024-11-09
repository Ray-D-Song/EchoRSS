package model

import "ray-d-song.com/echo-rss/db"

type Folder struct {
	ID        int    `json:"id" db:"id"`
	UserID    string `json:"userId" db:"user_id"`
	Name      string `json:"name" db:"name"`
	ParentID  int    `json:"parentId" db:"parent_id"`
	CreatedAt string `json:"createdAt" db:"created_at"`
}

func (f *Folder) List(userID string) ([]Folder, error) {
	folders := []Folder{}
	err := db.Bind.Select(&folders, `SELECT id, user_id, name, parent_id, created_at FROM folders WHERE is_deleted = 0 AND user_id = ?`, userID)
	return folders, err
}

func (f *Folder) Create() error {
	_, err := db.Bind.NamedExec(`INSERT INTO folders (user_id, name, parent_id) VALUES (:user_id, :name, :parent_id)`, f)
	return err
}

func (f *Folder) Delete() error {
	_, err := db.Bind.NamedExec(`UPDATE folders SET is_deleted = 1 WHERE id = :id`, f)
	return err
}

func (f *Folder) Update() error {
	_, err := db.Bind.NamedExec(`UPDATE folders SET name = :name WHERE id = :id`, f)
	return err
}
