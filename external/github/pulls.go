/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ksrof/trello-action/utils"
)

// Pulls wraps methods that handle pull endpoint requests
// of an authenticated user within the Github API.
//go:generate mockgen -destination=mock/pulls.go -package=mock . Pulls
type Pulls interface {
	Get(ctx context.Context, values map[string]string) (*PullsResponse, error)
}

// PullsResponse represents the data returned by the pull endpoint request.
type PullsResponse struct {
	Response interface{}
	Status   string
	Code     int
	Error    string
}

// Get returns a specific pull request by its identifier.
func (r *PullsResponse) Get(ctx context.Context, values map[string]string) (*PullsResponse, error) {
	fields, err := utils.FieldMapper(values)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	reqURL := fmt.Sprintf(fields["request_url"], fields["username"], fields["repository"], fields["pull_id"])
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return nil, utils.LogError(err.Error(), utils.LogPrefixInfo, utils.LogLevelInfo)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", fields["token"]))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, utils.LogError(err.Error(), utils.LogPrefixInfo, utils.LogLevelInfo)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, utils.LogError(err.Error(), utils.LogPrefixInfo, utils.LogLevelInfo)
	}

	var response map[string]any
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, utils.LogError(err.Error(), utils.LogPrefixInfo, utils.LogLevelInfo)
	}

	return &PullsResponse{
		Response: response,
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
	}, nil
}
