/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

package github

import (
	"log"
	"regexp"

	"github.com/ksrof/trello-action/utils"
)

// Auth wraps methods that handle the authorization
// of an user within the Github API.
type Auth interface {
	Basic() (token string)
}

type Option func(a *auth) error

type auth struct {
	token string
}

// NewAuth takes a set of options and returns
// an instance of *auth.
func NewAuth(opts ...Option) (Auth, error) {
	if len(opts) == 0 {
		return nil, utils.ErrEmptyOptions
	}

	a := &auth{}

	for _, opt := range opts {
		err := opt(a)
		if err != nil {
			log.Printf(
				"failed to add options, error: %s\n",
				err.Error(),
			)
			return nil, err
		}
	}

	return a, nil
}

// WithToken validates and adds the given token
// to the *auth.token struct field.
func WithToken(token string) Option {
	return func(a *auth) error {
		err := utils.Validations(
			utils.ValidateNotEmpty(token),
			utils.ValidateRegexp(
				*regexp.MustCompile("[A-Za-z0-9_]{40}"),
				token,
			),
		)
		if err != nil {
			log.Printf(
				"failed to validate *auth.token, error: %s\n",
				err.Error(),
			)
			return err
		}

		a.token = token
		return nil
	}
}

// Basic returns a Github Personal Access Token.
func (a *auth) Basic() (token string) {
	token = a.token
	return token
}
