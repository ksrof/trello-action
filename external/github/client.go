package github

import (
	"errors"
	"regexp"
	"strings"
)

var (
	minUserLength   int   = 4
	ErrEmptyOptions error = errors.New("no options provided")
	ErrEmptyHost    error = errors.New("no host provided")
	ErrEmptyUser    error = errors.New("no user provided")
	ErrEmptyRepo    error = errors.New("no repo provided")
	ErrEmptyEvent   error = errors.New("no event provided")
	ErrEmptyID      error = errors.New("no id provided")
	ErrInvalidHost  error = errors.New("invalid host provided")
	ErrInvalidUser  error = errors.New("invalid user provided")
	ErrInvalidEvent error = errors.New("invalid event provided")
	ErrInvalidID    error = errors.New("invalid id provided")
)

type Options func(c *Client) error

type Client struct {
	Host  string
	User  string
	Repo  string
	Event string
	ID    string
}

func NewClient(opts ...Options) (*Client, error) {
	c := &Client{}

	if len(opts) == 0 {
		return nil, ErrEmptyOptions
	}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func WithHost(host string) Options {
	return func(c *Client) error {
		err := checkHost(host)
		if err != nil {
			return err
		}

		c.Host = host
		return nil
	}
}

func WithUser(user string) Options {
	return func(c *Client) error {
		err := checkUser(user)
		if err != nil {
			return err
		}

		c.User = user
		return nil
	}
}

func WithRepo(repo string) Options {
	return func(c *Client) error {
		err := checkRepo(repo)
		if err != nil {
			return err
		}

		c.Repo = repo
		return nil
	}
}

func WithEvent(event string) Options {
	return func(c *Client) error {
		err := checkEvent(event)
		if err != nil {
			return err
		}

		c.Event = event
		return nil
	}
}

func WithID(id string) Options {
	return func(c *Client) error {
		err := checkID(id)
		if err != nil {
			return err
		}

		c.ID = id
		return nil
	}
}

func checkHost(host string) error {
	if strings.TrimSpace(host) == "" {
		return ErrEmptyHost
	}

	if !strings.Contains(host, "api.github.com") {
		return ErrInvalidHost
	}

	return nil
}

func checkUser(user string) error {
	if strings.TrimSpace(user) == "" {
		return ErrEmptyUser
	}

	if len(strings.TrimSpace(user)) < minUserLength {
		return ErrInvalidUser
	}

	return nil
}

func checkRepo(repo string) error {
	if strings.TrimSpace(repo) == "" {
		return ErrEmptyRepo
	}

	return nil
}

func checkEvent(event string) error {
	events := []string{
		"issues",
		"pull_request",
	}

	if strings.TrimSpace(event) == "" {
		return ErrEmptyEvent
	}

	for _, evt := range events {
		if evt != event {
			return ErrInvalidEvent
		}
	}

	return nil
}

func checkID(id string) error {
	if strings.TrimSpace(id) == "" {
		return ErrEmptyID
	}

	pattern := regexp.MustCompile("^[0-9]+$")
	ok := pattern.MatchString(id)
	if !ok {
		return ErrInvalidID
	}

	return nil
}
