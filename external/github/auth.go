package github

import (
	"regexp"

	"github.com/ksrof/trello-action/utils"
)

// Basic returns a Github Personal Access Token.
func Basic(value string) (token string, err error) {
	err = utils.Validations(
		utils.ValidateNotEmpty(value),
		utils.ValidateRegexp(*regexp.MustCompile("[A-Za-z0-9_]{40}"), value),
	)
	if err != nil {
		return "", err
	}

	return value, nil
}
