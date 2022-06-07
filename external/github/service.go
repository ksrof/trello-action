package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

var (
	ErrCreatingNewRequest error = errors.New("failed to create new request")
	ErrDoingRequest       error = errors.New("failed to do request")
	ErrParsingURL         error = errors.New("failed to parse request url")
)

type Service interface {
	GetIssueByID(ctx context.Context) (*http.Response, error)
	GetPullByID(ctx context.Context) (*http.Response, error)
	GetLabelsFromIssue(ctx context.Context) (*http.Response, error)
}

func (c *client) GetIssueByID(ctx context.Context) (*http.Response, error) {
	reqURL, err := buildURL([]string{
		c.user,
		c.repo,
		c.event,
		c.id,
	})
	if err != nil {
		return nil, err
	}

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
	reqURL, err := buildURL([]string{
		c.user,
		c.repo,
		c.event,
		c.id,
	})
	if err != nil {
		return nil, err
	}

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
	reqURL, err := buildURL([]string{
		c.user,
		c.repo,
		c.event,
		c.id,
		"labels",
	})
	if err != nil {
		return nil, err
	}

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

func buildURL(parts []string) (string, error) {
	u, err := url.Parse("https://api.github.com/")
	if err != nil {
		return "", ErrParsingURL
	}

	reqParts := path.Join(parts...)
	reqURL := fmt.Sprint(u.String(), reqParts)

	return reqURL, nil
}

// func decodeResponse(body []byte) (interface{}, error) {
// 	var response map[string]interface{}

// 	err := json.Unmarshal(body, &response)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return response, nil
// }
