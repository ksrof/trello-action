package github

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type options func(c *client) error

type client struct {
	host  string
	token string
	user  string
	repo  string
	event string
	id    string
}

func NewClient(opts ...options) (*client, error) {
	c := &client{}

	for _, opt := range append(defaultOpts(), opts...) {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func defaultOpts() []options {
	return []options{
		WithHost("https://api.github.com/"),
		WithToken(os.Getenv("GH_TOKEN")),
		WithUser(os.Getenv("GH_USERNAME")),
		WithRepo(os.Getenv("GH_REPOSITORY")),
		WithEvent(os.Getenv("GH_EVENT_TYPE")),
		WithID(os.Getenv("GH_EVENT_ID")),
	}
}

func checkTokenValidity(token string) error {
	ok, err := regexp.MatchString("[A-Za-z0-9_]{40}", token) // gh personal access token regexp
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("%w", errors.New("invalid token"))
	}

	return nil
}

func checkEmptyString(str string) error {
	if strings.TrimSpace(str) == "" {
		return fmt.Errorf("%w", errors.New("invalid repo"))
	}

	return nil
}

func WithHost(host string) options {
	return func(c *client) error {
		c.host = host
		return nil
	}
}

func WithToken(token string) options {
	return func(c *client) error {
		err := checkTokenValidity(token)
		if err != nil {
			return err
		}

		c.token = token
		return nil
	}
}

func WithUser(user string) options {
	return func(c *client) error {
		err := checkEmptyString(user)
		if err != nil {
			return err
		}

		c.user = user
		return nil
	}
}

func WithRepo(repo string) options {
	return func(c *client) error {
		err := checkEmptyString(repo)
		if err != nil {
			return err
		}

		c.repo = repo
		return nil
	}
}

func WithEvent(event string) options {
	return func(c *client) error {
		err := checkEmptyString(event)
		if err != nil {
			return err
		}

		c.event = event
		return nil
	}
}

func WithID(id string) options {
	return func(c *client) error {
		err := checkEmptyString(id)
		if err != nil {
			return err
		}

		c.id = id
		return nil
	}
}
