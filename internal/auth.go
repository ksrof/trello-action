package internal

import (
	"fmt"
	"net/http"
	"regexp"
)

// Auth validates, parses and sets the authorization header for the given request.
// It returns the given request with the authorization header set.
func Auth(token string, req *http.Request) (*http.Request, error) {
	match, err := regexp.MatchString("[A-Za-z0-9_]{40}", token)
	if err != nil {
		return req, fmt.Errorf("failed to run regexp: %v", err)
	}

	if !match {
		return req, fmt.Errorf("failed to check token validity: %v", token)
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	return req, nil
}
