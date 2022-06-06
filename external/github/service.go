package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrCreatingNewRequest error = errors.New("failed to create new request")
	ErrDoingRequest       error = errors.New("failed to do request")
)

type Service interface {
	GetIssueByID(ctx context.Context) (*http.Response, error)
	GetPullByID(ctx context.Context) (*http.Response, error)
	GetLabelsFromIssue(ctx context.Context) (*http.Response, error)
}

func (c *client) GetIssueByID(ctx context.Context) (*http.Response, error) {
	reqURL := fmt.Sprintf(
		"%s/repos/%s/%s/%s/%s",
		"https://api.github.com",
		c.user,
		c.repo,
		c.event,
		c.id,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return nil, ErrCreatingNewRequest
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", c.token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, ErrDoingRequest
	}

	return res, nil
}

func (c *client) GetPullByID(ctx context.Context) (*http.Response, error) {
	reqURL := fmt.Sprintf(
		"%s/repos/%s/%s/%s/%s",
		"https://api.github.com",
		c.user,
		c.repo,
		c.event,
		c.id,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return nil, ErrCreatingNewRequest
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", c.token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, ErrDoingRequest
	}

	return res, nil
}

func (c *client) GetLabelsFromIssue(ctx context.Context) (*http.Response, error) {
	reqURL := fmt.Sprintf(
		"%s/repos/%s/%s/%s/%s/%s",
		"https://api.github.com",
		c.user,
		c.repo,
		c.event,
		c.id,
		"labels",
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return nil, ErrCreatingNewRequest
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", c.token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, ErrDoingRequest
	}

	return res, nil
}
