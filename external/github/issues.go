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

// TODO: Generate mocks for the Issues interface.
// TODO: Test that every method works as expected.

// Issues wraps methods that handle issue requests
// of an authenticated user within the Github API.
//go:generate mockgen -destination=mock/issues.go -package=mock . Issues
type Issues interface {
	Get(ctx context.Context, opts []utils.Field) (*IssuesResponse, error)
	GetLabels(ctx context.Context, opts []utils.Field) (*IssuesResponse, error)
}

// IssueResponse represents the data returned by issue requests.
type IssuesResponse struct {
	Response interface{}
	Status   string
	Code     int
	Error    string
}

// Get returns a specific issue by its identifier.
func (r *IssuesResponse) Get(ctx context.Context, opts []utils.Field) (*IssuesResponse, error) {
	fields, err := utils.NewFieldsMapper(opts...)
	if err != nil {
		return &IssuesResponse{
			Status: http.StatusText(http.StatusBadRequest),
			Code:   http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	field := fields.(map[string]string)
	reqURL := fmt.Sprintf(field["request_url"], field["username"], field["repository"], field["issue_id"])

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return &IssuesResponse{
				Status: http.StatusText(http.StatusInternalServerError),
				Code:   http.StatusInternalServerError,
				Error:  err.Error(),
			}, utils.NewError(
				utils.WithLogger(
					err.Error(),
					utils.LogPrefixInfo,
					utils.LogLevelInfo,
				),
			)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", field["token"]))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &IssuesResponse{
				Status: http.StatusText(http.StatusInternalServerError),
				Code:   http.StatusInternalServerError,
				Error:  err.Error(),
			}, utils.NewError(
				utils.WithLogger(
					err.Error(),
					utils.LogPrefixInfo,
					utils.LogLevelInfo,
				),
			)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &IssuesResponse{
				Status: http.StatusText(http.StatusInternalServerError),
				Code:   http.StatusInternalServerError,
				Error:  err.Error(),
			}, utils.NewError(
				utils.WithLogger(
					err.Error(),
					utils.LogPrefixInfo,
					utils.LogLevelInfo,
				),
			)
	}

	var response map[string]any
	err = json.Unmarshal(data, &response)
	if err != nil {
		return &IssuesResponse{
				Status: http.StatusText(http.StatusInternalServerError),
				Code:   http.StatusInternalServerError,
				Error:  err.Error(),
			}, utils.NewError(
				utils.WithLogger(
					err.Error(),
					utils.LogPrefixInfo,
					utils.LogLevelInfo,
				),
			)
	}

	return &IssuesResponse{
		Response: response,
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
	}, nil
}

// GetLabels returns the labels from a specific issue by its identifier.
func (r *IssuesResponse) GetLabels(ctx context.Context, opts []utils.Field) (*IssuesResponse, error) {
	fields, err := utils.NewFieldsMapper(opts...)
	if err != nil {
		return &IssuesResponse{
			Status: http.StatusText(http.StatusBadRequest),
			Code:   http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	field := fields.(map[string]string)
	reqURL := fmt.Sprintf(field["request_url"], field["username"], field["repository"], field["issue_id"])

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return &IssuesResponse{
				Status: http.StatusText(http.StatusInternalServerError),
				Code:   http.StatusInternalServerError,
				Error:  err.Error(),
			}, utils.NewError(
				utils.WithLogger(
					err.Error(),
					utils.LogPrefixInfo,
					utils.LogLevelInfo,
				),
			)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", field["token"]))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &IssuesResponse{
				Status: http.StatusText(http.StatusInternalServerError),
				Code:   http.StatusInternalServerError,
				Error:  err.Error(),
			}, utils.NewError(
				utils.WithLogger(
					err.Error(),
					utils.LogPrefixInfo,
					utils.LogLevelInfo,
				),
			)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &IssuesResponse{
				Status: http.StatusText(http.StatusInternalServerError),
				Code:   http.StatusInternalServerError,
				Error:  err.Error(),
			}, utils.NewError(
				utils.WithLogger(
					err.Error(),
					utils.LogPrefixInfo,
					utils.LogLevelInfo,
				),
			)
	}

	var response []map[string]any
	err = json.Unmarshal(data, &response)
	if err != nil {
		return &IssuesResponse{
				Status: http.StatusText(http.StatusInternalServerError),
				Code:   http.StatusInternalServerError,
				Error:  err.Error(),
			}, utils.NewError(
				utils.WithLogger(
					err.Error(),
					utils.LogPrefixInfo,
					utils.LogLevelInfo,
				),
			)
	}

	return &IssuesResponse{
		Response: response,
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
	}, nil
}
