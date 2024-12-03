package model

import (
	"errors"

	"go.uber.org/zap"
	"ray-d-song.com/echo-rss/db"
	"ray-d-song.com/echo-rss/utils"
)

type TranslateCache struct {
	Url     string `json:"url" db:"url"`
	Content string `json:"content" db:"content"`
}

func (m *TranslateCache) Get() (string, error) {
	if m.Url == "" {
		return "", errors.New("url is required")
	}
	var content string
	err := db.Bind.Get(&content, "SELECT content FROM translate_cache WHERE url = ?", m.Url)
	if content != "" {
		utils.Logger.Info("cache hit", zap.String("url", m.Url))
	}
	return content, err
}

func (m *TranslateCache) Set() error {
	if m.Url == "" {
		return errors.New("url is required")
	}
	if m.Content == "" {
		return errors.New("content is required")
	}
	_, err := db.Bind.Exec("INSERT INTO translate_cache (url, content) VALUES (?, ?) ON CONFLICT(url) DO UPDATE SET content = ?", m.Url, m.Content, m.Content)
	if err != nil {
		utils.Logger.Error("failed to set cache", zap.String("url", m.Url), zap.Error(err))
	}
	return err
}
