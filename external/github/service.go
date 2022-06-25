package github

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func (c *client) GetIssueByID(ctx context.Context) (*http.Response, error) {
	reqURL := fmt.Sprintf(
		"%s/%s/%s/%s/%s/%s",
		c.Host(),
		"repos",
		c.User(),
		c.Repo(),
		c.Event(),
		c.ID(),
	)

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", c.Token()))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) GetPullByID(ctx context.Context) (*http.Response, error) {
	reqURL := fmt.Sprintf(
		"%s/%s/%s/%s/%s/%s",
		c.Host(),
		"repos",
		c.User(),
		c.Repo(),
		c.Event(),
		c.ID(),
	)

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", c.Token()))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) GetIssueLabels(ctx context.Context) (*http.Response, error) {
	reqURL := fmt.Sprintf(
		"%s/%s/%s/%s/%s/%s/%s",
		c.Host(),
		"repos",
		c.User(),
		c.Repo(),
		c.Event(),
		c.ID(),
		"labels",
	)

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", c.Token()))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
