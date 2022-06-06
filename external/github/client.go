package github

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

var (
	minUserLength   int   = 4
	ErrEmptyString  error = errors.New("string is empty")
	ErrInvalidToken error = errors.New("invalid github token")
	ErrInvalidUser  error = errors.New("invalid github user")
	ErrInvalidEvent error = errors.New("invalid github event")
	ErrInvalidID    error = errors.New("invalid github issue/pull id")
)

type clientOption func(c *client) error

type client struct {
	token string
	user  string
	repo  string
	event string
	id    string
}

func NewClient(opts ...clientOption) (*client, error) {
	c := &client{}

	for _, opt := range append(defaults(), opts...) {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func defaults() []clientOption {
	return []clientOption{
		WithToken(os.Getenv("GH_TOKEN")),
		WithUser(os.Getenv("GH_USER")),
		WithRepo(os.Getenv("GH_REPO")),
		WithEvent(os.Getenv("GH_EVENT")),
		WithID(os.Getenv("GH_ID")),
	}
}

func WithToken(token string) clientOption {
	return func(c *client) error {
		err := checkToken(token)
		if err != nil {
			return err
		}

		c.token = token
		return nil
	}
}

func WithUser(user string) clientOption {
	return func(c *client) error {
		err := checkUser(user)
		if err != nil {
			return err
		}

		c.user = user
		return nil
	}
}

func WithRepo(repo string) clientOption {
	return func(c *client) error {
		err := checkRepo(repo)
		if err != nil {
			return err
		}

		c.repo = repo
		return nil
	}
}

func WithEvent(event string) clientOption {
	return func(c *client) error {
		err := checkEvent(event)
		if err != nil {
			return err
		}

		c.event = event
		return nil
	}
}

func WithID(id string) clientOption {
	return func(c *client) error {
		err := checkID(id)
		if err != nil {
			return err
		}

		c.id = id
		return nil
	}
}

func checkToken(token string) error {
	err := checkEmpty(token)
	if err != nil {
		return err
	}

	pattern := regexp.MustCompile("[A-Za-z0-9_]{40}")
	ok := pattern.MatchString(token)

	if !ok {
		return ErrInvalidToken
	}

	return nil
}

func checkUser(user string) error {
	err := checkEmpty(user)
	if err != nil {
		return err
	}

	if len(strings.TrimSpace(user)) < minUserLength {
		return ErrInvalidUser
	}

	return nil
}

func checkRepo(repo string) error {
	err := checkEmpty(repo)
	if err != nil {
		return err
	}

	return nil
}

func checkEvent(event string) error {
	err := checkEmpty(event)
	if err != nil {
		return err
	}

	switch event {
	case "issues":
		return nil
	case "pulls":
		return nil
	default:
		return ErrInvalidEvent
	}
}

func checkID(id string) error {
	err := checkEmpty(id)
	if err != nil {
		return err
	}

	pattern := regexp.MustCompile("^[0-9]+$")
	ok := pattern.MatchString(id)

	if !ok {
		return ErrInvalidID
	}

	return nil
}

func checkEmpty(str string) error {
	if strings.TrimSpace(str) == "" {
		return ErrEmptyString
	}

	return nil
}
