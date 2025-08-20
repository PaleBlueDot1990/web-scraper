package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	path := strings.TrimSuffix(u.Path, "/") 
	return u.Host + path, nil
}

