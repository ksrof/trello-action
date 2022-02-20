package utils

import (
	"errors"
	"fmt"
)

// SetAuth returns the Github Personal Access Token
// provided by the environment variables.
func SetAuth(token string) (string, error) {
	if len(token) <= 0 {
		return "", fmt.Errorf("token needs to be provided: %v", errors.New("no token provided"))
	}

	return fmt.Sprintf("token %s", token), nil
}
