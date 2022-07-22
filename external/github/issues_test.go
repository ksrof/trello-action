/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

package github_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ksrof/trello-action/external/github"
	"github.com/ksrof/trello-action/external/github/mock"
	"github.com/ksrof/trello-action/utils"
	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	issues := mock.NewMockIssues(ctrl)

	type args struct {
		ctx  context.Context
		opts []utils.Field
	}

	tests := []struct {
		name    string
		args    args
		want    *github.IssuesResponse
		wantErr error
		mocks   func(issues *mock.MockIssues)
	}{
		{
			name: "issues.Get() returns a successful *github.IssuesResponse struct",
			args: args{
				ctx: ctx,
				opts: []utils.Field{
					utils.WithMap(
						map[string]string{
							"username":    "ksrof",
							"repository":  "trello-action",
							"issue_id":    "30",
							"request_url": "https://api.github.com",
							"token":       "ghp_F41fR24d9Tvn3YRzC7GdOPAfhgjBzP5MLP7c",
						},
					),
				},
			},
			want: &github.IssuesResponse{
				Response: map[string]any{
					"id":             30,
					"node_id":        "MDU6SXNzdWUx",
					"url":            "https://api.github.com/",
					"repository_url": "https://api.github.com/",
					"labels_url":     "https://api.github.com/",
					"comments_url":   "https://api.github.com/",
					"events_url":     "https://api.github.com/",
					"html_url":       "https://github.com/",
					"number":         30,
					"state":          "open",
					"title":          "Found a bug",
					"body":           "I'm having a problem with this.",
				},
				Status: http.StatusText(http.StatusOK),
				Code:   http.StatusOK,
			},
			mocks: func(issues *mock.MockIssues) {
				issues.EXPECT().Get(ctx, gomock.Any()).
					Return(
						&github.IssuesResponse{
							Response: map[string]any{
								"id":             30,
								"node_id":        "MDU6SXNzdWUx",
								"url":            "https://api.github.com/",
								"repository_url": "https://api.github.com/",
								"labels_url":     "https://api.github.com/",
								"comments_url":   "https://api.github.com/",
								"events_url":     "https://api.github.com/",
								"html_url":       "https://github.com/",
								"number":         30,
								"state":          "open",
								"title":          "Found a bug",
								"body":           "I'm having a problem with this.",
							},
							Status: http.StatusText(http.StatusOK),
							Code:   http.StatusOK,
						}, nil).MaxTimes(1)
			},
		},
		{
			name: "issues.Get() returns a failed *github.IssuesResponse struct if map value is empty",
			args: args{
				ctx: ctx,
				opts: []utils.Field{
					utils.WithMap(
						map[string]string{
							"username": "",
						},
					),
				},
			},
			want: &github.IssuesResponse{
				Status: http.StatusText(http.StatusBadRequest),
				Code:   http.StatusBadRequest,
				Error:  utils.ErrEmptyValue.Error(),
			},
			wantErr: utils.ErrEmptyValue,
			mocks: func(issues *mock.MockIssues) {
				issues.EXPECT().Get(ctx, gomock.Any()).
					Return(
						&github.IssuesResponse{
							Status: http.StatusText(http.StatusBadRequest),
							Code:   http.StatusBadRequest,
							Error:  utils.ErrEmptyValue.Error(),
						}, nil).MaxTimes(1)
			},
		},
		// TODO: Add more test cases
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mocks(issues)
			issue, err := issues.Get(tc.args.ctx, tc.args.opts)
			if err != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				assert.Equal(t, tc.want, issue)
				return
			}

			assert.Equal(t, tc.want, issue)
		})
	}
}

func TestGetLabels(t *testing.T) {
	ctrl := gomock.NewController(t)
	issues := mock.NewMockIssues(ctrl)

	type args struct {
		ctx  context.Context
		opts []utils.Field
	}

	tests := []struct {
		name    string
		args    args
		want    *github.IssuesResponse
		wantErr error
		mocks   func(issues *mock.MockIssues)
	}{
		{
			name: "issues.GetLabels() returns a successful *github.IssuesResponse struct",
			args: args{
				ctx: ctx,
				opts: []utils.Field{
					utils.WithMap(
						map[string]string{
							"username":    "ksrof",
							"repository":  "trello-action",
							"issue_id":    "30",
							"request_url": "https://api.github.com",
							"token":       "ghp_F41fR24d9Tvn3YRzC7GdOPAfhgjBzP5MLP7c",
						},
					),
				},
			},
			want: &github.IssuesResponse{
				Response: []map[string]any{
					{
						"id":          208045946,
						"node_id":     "MDU6TGFiZWwyMDgwNDU5NDY=",
						"url":         "https://api.github.com/",
						"name":        "bug",
						"description": "Something isn't working",
						"color":       "f29513",
						"default":     true,
					},
					{

						"id":          208045947,
						"node_id":     "MDU6TGFiZWwyMDgwNDU5NDc=",
						"url":         "https://api.github.com/",
						"name":        "enhancement",
						"description": "New feature or request",
						"color":       "a2eeef",
						"default":     false,
					},
				},
				Status: http.StatusText(http.StatusOK),
				Code:   http.StatusOK,
			},
			mocks: func(issues *mock.MockIssues) {
				issues.EXPECT().GetLabels(ctx, gomock.Any()).
					Return(
						&github.IssuesResponse{
							Response: []map[string]any{
								{
									"id":          208045946,
									"node_id":     "MDU6TGFiZWwyMDgwNDU5NDY=",
									"url":         "https://api.github.com/",
									"name":        "bug",
									"description": "Something isn't working",
									"color":       "f29513",
									"default":     true,
								},
								{

									"id":          208045947,
									"node_id":     "MDU6TGFiZWwyMDgwNDU5NDc=",
									"url":         "https://api.github.com/",
									"name":        "enhancement",
									"description": "New feature or request",
									"color":       "a2eeef",
									"default":     false,
								},
							},
							Status: http.StatusText(http.StatusOK),
							Code:   http.StatusOK,
						}, nil).MaxTimes(1)
			},
		},
		{
			name: "issues.GetLabels() returns a failed *github.IssuesResponse struct if map value is empty",
			args: args{
				ctx: ctx,
				opts: []utils.Field{
					utils.WithMap(
						map[string]string{
							"username": "",
						},
					),
				},
			},
			want: &github.IssuesResponse{
				Status: http.StatusText(http.StatusBadRequest),
				Code:   http.StatusBadRequest,
				Error:  utils.ErrEmptyValue.Error(),
			},
			wantErr: utils.ErrEmptyValue,
			mocks: func(issues *mock.MockIssues) {
				issues.EXPECT().GetLabels(ctx, gomock.Any()).
					Return(
						&github.IssuesResponse{
							Status: http.StatusText(http.StatusBadRequest),
							Code:   http.StatusBadRequest,
							Error:  utils.ErrEmptyValue.Error(),
						}, nil).MaxTimes(1)
			},
		},
		// TODO: Add more test cases
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mocks(issues)
			issueLabels, err := issues.GetLabels(tc.args.ctx, tc.args.opts)
			if err != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				assert.Equal(t, tc.want, issueLabels)
				return
			}

			assert.Equal(t, tc.want, issueLabels)
		})
	}
}
