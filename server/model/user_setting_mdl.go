package model

import (
	"database/sql"

	"ray-d-song.com/echo-rss/db"
)

type UserSetting struct {
	// key will be user's id
	UserId          string `json:"-" db:"user_id"`
	OPENAI_API_KEY  string `json:"OPENAI_API_KEY" db:"OPENAI_API_KEY"`
	API_ENDPOINT    string `json:"API_ENDPOINT" db:"API_ENDPOINT"`
	TARGET_LANGUAGE string `json:"TARGET_LANGUAGE" db:"TARGET_LANGUAGE"`
}

func (u *UserSetting) CreateUserSetting() error {
	_, err := db.Bind.NamedExec("INSERT INTO user_setting (user_id) VALUES (:user_id)", u)
	return err
}

func (u *UserSetting) GetUserSetting() error {
	err := db.Bind.Get(u, "SELECT * FROM user_setting WHERE user_id = :user_id", u.UserId)
	if err == nil {
		return nil
	}
	if err == sql.ErrNoRows {
		// create config
		err = u.CreateUserSetting()
		if err != nil {
			return err
		}
		return u.GetUserSetting()
	}
	return err
}

func (u *UserSetting) UpdateAiSetting() error {
	_, err := db.Bind.NamedExec("UPDATE user_setting SET OPENAI_API_KEY = :OPENAI_API_KEY, API_ENDPOINT = :API_ENDPOINT, TARGET_LANGUAGE = :TARGET_LANGUAGE WHERE user_id = :user_id", u)
	return err
}
