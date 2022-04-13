package utils

import (
	"fmt"
	"net/url"
	"path"
)

// ParseURL joins both the basePath and the newPath.
// It returns a string that represents the new path.
func ParseURL(basePath, newPath string) (string, error) {
	baseURL, err := url.Parse(basePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse base path: %v", err)
	}

	baseURL.Path = path.Join(baseURL.Path, newPath)

	return baseURL.String(), nil
}
