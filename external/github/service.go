package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *client) GetIssueByID(ctx context.Context) (data map[string]string, err error) {
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

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *client) GetPullByID(ctx context.Context) (data map[string]string, err error) {
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

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *client) GetIssueLabels(ctx context.Context) (data []map[string]string, err error) {
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

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
