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

// Issues wraps methods that handle issue endpoint requests
// of an authenticated user within the Github API.
//go:generate mockgen -destination=mock/issues.go -package=mock . Issues
type Issues interface {
	Get(ctx context.Context, values map[string]string) (*IssuesResponse, error)
	GetLabels(ctx context.Context, values map[string]string) (*IssuesResponse, error)
}

// IssueResponse represents the data returned by the issue endpoint request.
type IssuesResponse struct {
	Response interface{}
	Status   string
	Code     int
	Error    string
}

// Get returns a specific issue by its identifier.
func (r *IssuesResponse) Get(ctx context.Context, values map[string]string) (*IssuesResponse, error) {
	fields, err := utils.FieldMapper(values)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	reqURL := fmt.Sprintf(fields["request_url"], fields["username"], fields["repository"], fields["issue_id"])
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

	return &IssuesResponse{
		Response: response,
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
	}, nil
}

// GetLabels returns the labels from a specific issue by its identifier.
func (r *IssuesResponse) GetLabels(ctx context.Context, values map[string]string) (*IssuesResponse, error) {
	fields, err := utils.FieldMapper(values)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	reqURL := fmt.Sprintf(fields["request_url"], fields["username"], fields["repository"], fields["issue_id"])
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

	var response []map[string]any
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, utils.LogError(err.Error(), utils.LogPrefixInfo, utils.LogLevelInfo)
	}

	return &IssuesResponse{
		Response: response,
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
	}, nil
}
