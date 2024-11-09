package utils

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func GetFavicon(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}

	var faviconURL string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "link" {
			for _, attr := range n.Attr {
				if attr.Key == "rel" && (attr.Val == "icon" || attr.Val == "shortcut icon") {
					for _, attr := range n.Attr {
						if attr.Key == "href" {
							faviconURL = attr.Val
							return
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if faviconURL == "" {
		return "", fmt.Errorf("favicon not found")
	}

	resp, err = http.Get(faviconURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	faviconData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	encodedFavicon := base64.StdEncoding.EncodeToString(faviconData)
	return encodedFavicon, nil
}
