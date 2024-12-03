package model

import "ray-d-song.com/echo-rss/db"

func ToggleItemBookmark(userID, itemID string) error {
	_, err := db.Bind.Exec("UPDATE items SET bookmark = CASE WHEN bookmark = 0 THEN 1 ELSE 0 END WHERE id = ? AND user_id = ?", itemID, userID)
	return err
}
