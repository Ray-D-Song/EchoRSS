package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Translate(content string, openaiApiKey string, apiEndpoint string, targetLanguage string) (string, error) {
	formData := url.Values{}
	formData.Set("MD_CONTENT", content)
	formData.Set("TARGET_LANGUAGE", targetLanguage)

	resp, err := http.PostForm(apiEndpoint, formData)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to translate: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
