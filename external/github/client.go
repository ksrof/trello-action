package github

import (
	"context"
	"errors"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	minUsernameLength    int   = 4
	ErrInvalidHost       error = errors.New("invalid github host url")
	ErrInvalidToken      error = errors.New("invalid github personal access token")
	ErrInvalidUser       error = errors.New("invalid github username")
	ErrInvalidRepo       error = errors.New("invalid github repository")
	ErrInvalidEvent      error = errors.New("invalid github action event")
	ErrInvalidIdentifier error = errors.New("invalid github action event identifier")
)

type Option func(c *client) error

type Client interface {
	Host() string
	Token() string
	User() string
	Repo() string
	Event() string
	ID() string

	GetIssueByID(ctx context.Context) (*http.Response, error)
	GetPullByID(ctx context.Context) (*http.Response, error)
	GetIssueLabels(ctx context.Context) (*http.Response, error)
}

type client struct {
	host  string
	token string
	user  string
	repo  string
	event string
	id    string
}

func NewClient(opts ...Option) (Client, error) {
	c := &client{}

	for _, opt := range append(defaults(), opts...) {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func defaults() []Option {
	return []Option{
		WithHost("https://api.github.com"),
		WithToken(os.Getenv("GH_TOKEN")),
		WithUser(os.Getenv("GH_USER")),
		WithRepo(os.Getenv("GH_REPO")),
		WithEvent(os.Getenv("GH_EVENT")),
		WithID(os.Getenv("GH_ID")),
	}
}

func (c *client) Host() string {
	return c.host
}

func WithHost(host string) Option {
	return func(c *client) error {
		err := checkHost(host)
		if err != nil {
			return err
		}

		c.host = host
		return nil
	}
}

func checkHost(host string) error {
	if strings.TrimSpace(host) == "" {
		return ErrInvalidHost
	}

	return nil
}

func (c *client) Token() string {
	return c.token
}

func WithToken(token string) Option {
	return func(c *client) error {
		err := checkToken(token)
		if err != nil {
			return err
		}

		c.token = token
		return nil
	}
}

func checkToken(token string) error {
	pattern := regexp.MustCompile("[A-Za-z0-9_]{40}")
	ok := pattern.MatchString(token)
	if !ok {
		return ErrInvalidToken
	}

	return nil
}

func (c *client) User() string {
	return c.user
}

func WithUser(user string) Option {
	return func(c *client) error {
		err := checkUser(user)
		if err != nil {
			return err
		}

		c.user = user
		return nil
	}
}

func checkUser(user string) error {
	if len(strings.TrimSpace(user)) < minUsernameLength {
		return ErrInvalidUser
	}

	return nil
}

func (c *client) Repo() string {
	return c.repo
}

func WithRepo(repo string) Option {
	return func(c *client) error {
		err := checkRepo(repo)
		if err != nil {
			return err
		}

		c.repo = repo
		return nil
	}
}

func checkRepo(repo string) error {
	if strings.TrimSpace(repo) == "" {
		return ErrInvalidRepo
	}

	return nil
}

func (c *client) Event() string {
	return c.event
}

func WithEvent(event string) Option {
	return func(c *client) error {
		err := checkEvent(event)
		if err != nil {
			return err
		}

		c.event = event
		return nil
	}
}

func checkEvent(event string) error {
	events := []string{
		"issues",
		"pulls",
	}

	for _, evt := range events {
		if evt == event {
			return nil
		}
	}

	return ErrInvalidEvent
}

func (c *client) ID() string {
	return c.id
}

func WithID(id string) Option {
	return func(c *client) error {
		err := checkID(id)
		if err != nil {
			return err
		}

		c.id = id
		return nil
	}
}

func checkID(id string) error {
	pattern := regexp.MustCompile("^[0-9]+$")
	ok := pattern.MatchString(id)
	if !ok {
		return ErrInvalidIdentifier
	}

	return nil
}
