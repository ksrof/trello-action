/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

package github

import (
	"log"

	"github.com/ksrof/trello-action/utils"
)

// Auth wraps methods that handle the authorization
// of an user within the Github API.
type Auth interface {
	Basic() string
}

type Option func(a *auth) error

type auth struct {
	user  string
	repo  string
	token string
}

// NewAuth takes a set of options and returns
// an instance of *auth.
func NewAuth(opts ...Option) (Auth, error) {
	a := &auth{}

	for _, opt := range opts {
		err := opt(a)
		if err != nil {
			log.Printf(
				"failed to add options, error: %s",
				err.Error(),
			)
			return nil, err
		}
	}

	return a, nil
}

// WithUser validates and adds the given user
// to the *auth.user struct field.
func WithUser(user string) Option {
	return func(a *auth) error {
		err := utils.Validations(
			utils.ValidateType(user),
			utils.ValidateNotZero(user),
		)
		if err != nil {
			log.Printf(
				"failed to validate *auth.user, error: %s",
				err.Error(),
			)
			return err
		}

		a.user = user
		return nil
	}
}

// WithRepo validates and adds the given user
// to the *auth.repo struct field.
func WithRepo(repo string) Option {
	return func(a *auth) error {
		err := utils.Validations(
			utils.ValidateType(repo),
			utils.ValidateNotZero(repo),
		)
		if err != nil {
			log.Printf(
				"failed to validate *auth.repo, error: %s",
				err.Error(),
			)
			return err
		}

		a.repo = repo
		return nil
	}
}

// WithToken validates and adds the given token
// to the *auth.token struct field.
func WithToken(token string) Option {
	return func(a *auth) error {
		err := utils.Validations(
			utils.ValidateType(token),
			utils.ValidateNotZero(token),
			// TODO: Add ValidateRegexp validation.
		)
		if err != nil {
			log.Printf(
				"failed to validate *auth.token, error: %s",
				err.Error(),
			)
			return err
		}

		a.token = token
		return nil
	}
}

// Basic returns a Github Personal Access Token,
// or an error string in case of failure.
func (a *auth) Basic() string {
	err := utils.Validations(
		utils.ValidateNotZero(a.token),
	)
	if err != nil {
		log.Printf(
			"failed to validate *auth.token, error: %s",
			err.Error(),
		)
		return err.Error()
	}

	return a.token
}
