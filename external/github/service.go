package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Service interface {
	GetIssueByID(ctx context.Context) (interface{}, error)
	GetPullByID(ctx context.Context) (interface{}, error)
	GetLabelsFromIssue(ctx context.Context) (interface{}, error)
}

type SVC struct {
	client *client
}

func NewService(cli *client) SVC {
	return SVC{
		client: cli,
	}
}

func decodeResponse() (interface{}, error) {
	// TODO: Decode http.Response in a separated function.
	return nil, nil
}

func (s *SVC) GetIssueByID(ctx context.Context) (interface{}, error) { // TODO: Do request
	reqURL := fmt.Sprintf(
		"%s/repos/%s/%s/%s/%s", // TODO: use url.Parse and path.Join
		s.client.host,
		s.client.user,
		s.client.repo,
		s.client.event,
		s.client.id,
	)

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf( // TODO: Make an auth with token function
		"token %s",
		s.client.token,
	))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err // TODO: Handle API errors (docs)
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%w", errors.New(res.Status))
	}

	fmt.Println(res.Status)

	_, _ = decodeResponse() // TODO: Decode request response

	return nil, nil
}

func (s *SVC) GetPullByID(ctx context.Context) (interface{}, error) { // TODO: Do request
	_, _ = decodeResponse()
	return nil, nil
}

func (s *SVC) GetLabelsFromIssue(ctx context.Context) (interface{}, error) { // TODO: Do request
	_, _ = decodeResponse()
	return nil, nil
}
