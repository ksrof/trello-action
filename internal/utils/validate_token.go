package utils

import (
	"fmt"
	"regexp"
)

// ValidateToken matches the github personal access token against a regexp.
func ValidateToken(token string) error {
	match, err := regexp.MatchString("[A-Za-z0-9_]{40}", token)
	if err != nil {
		return fmt.Errorf("failed to run regexp: %v", err)
	}

	if !match {
		return fmt.Errorf("token is invalid: %v", err)
	}

	return nil
}
