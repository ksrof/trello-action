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

func TestPulls_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	pulls := mock.NewMockPulls(ctrl)

	type args struct {
		ctx  context.Context
		opts []utils.Field
	}

	tests := []struct {
		name    string
		args    args
		want    *github.PullsResponse
		errStr  string
		wantErr error
		mocks   func(pulls *mock.MockPulls)
	}{
		{
			name: "pulls.Get() returns a successful *github.PullsResponse struct",
			args: args{
				ctx: ctx,
				opts: []utils.Field{
					utils.WithMap(
						map[string]string{
							"username":    "ksrof",
							"repository":  "trello-action",
							"pull_id":     "30",
							"request_url": "https://api.github.com",
							"token":       "ghp_F41fR24d9Tvn3YRzC7GdOPAfhgjBzP5MLP7c",
						},
					),
				},
			},
			want: &github.PullsResponse{
				Response: map[string]any{
					"url":                 "https://api.github.com/repos/octocat/Hello-World/pulls/1347",
					"id":                  30,
					"node_id":             "MDExOlB1bGxSZXF1ZXN0MQ==",
					"html_url":            "https://github.com/",
					"diff_url":            "https://github.com/",
					"patch_url":           "https://github.com/",
					"issue_url":           "https://api.github.com/",
					"commits_url":         "https://api.github.com/",
					"review_comments_url": "https://api.github.com/",
					"review_comment_url":  "https://api.github.com/",
					"comments_url":        "https://api.github.com/",
					"statuses_url":        "https://api.github.com/",
					"number":              30,
					"state":               "open",
					"locked":              true,
					"title":               "Amazing new feature",
				},
				Status: http.StatusText(http.StatusOK),
				Code:   http.StatusOK,
			},
			wantErr: nil,
			mocks: func(pulls *mock.MockPulls) {
				pulls.EXPECT().Get(ctx, gomock.Any()).
					Return(
						&github.PullsResponse{
							Response: map[string]any{
								"url":                 "https://api.github.com/repos/octocat/Hello-World/pulls/1347",
								"id":                  30,
								"node_id":             "MDExOlB1bGxSZXF1ZXN0MQ==",
								"html_url":            "https://github.com/",
								"diff_url":            "https://github.com/",
								"patch_url":           "https://github.com/",
								"issue_url":           "https://api.github.com/",
								"commits_url":         "https://api.github.com/",
								"review_comments_url": "https://api.github.com/",
								"review_comment_url":  "https://api.github.com/",
								"comments_url":        "https://api.github.com/",
								"statuses_url":        "https://api.github.com/",
								"number":              30,
								"state":               "open",
								"locked":              true,
								"title":               "Amazing new feature",
							},
							Status: http.StatusText(http.StatusOK),
							Code:   http.StatusOK,
						}, nil).MaxTimes(1)
			},
		},
		{
			name: "pulls.Get() returns a failed *github.PullsResponse struct if map is empty",
			args: args{
				ctx: ctx,
				opts: []utils.Field{
					utils.WithMap(map[string]string{}),
				},
			},
			want: &github.PullsResponse{
				Status: http.StatusText(http.StatusBadRequest),
				Code:   http.StatusBadRequest,
				Error:  utils.ErrZeroLength.Error(),
			},
			errStr: utils.ErrZeroLength.Error(),
			mocks: func(pulls *mock.MockPulls) {
				pulls.EXPECT().Get(ctx, gomock.Any()).
					Return(
						&github.PullsResponse{
							Status: http.StatusText(http.StatusBadRequest),
							Code:   http.StatusBadRequest,
							Error:  utils.ErrZeroLength.Error(),
						}, utils.NewError(
							utils.WithLogger(
								utils.ErrZeroLength.Error(),
								utils.LogPrefixInfo,
								utils.LogLevelInfo,
							),
						)).MaxTimes(1)
			},
		},
		{
			name: "pulls.Get() returns a failed *github.PullsResponse struct if map value is empty",
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
			want: &github.PullsResponse{
				Status: http.StatusText(http.StatusBadRequest),
				Code:   http.StatusBadRequest,
				Error:  utils.ErrZeroLength.Error(),
			},
			errStr: utils.ErrZeroLength.Error(),
			mocks: func(pulls *mock.MockPulls) {
				pulls.EXPECT().Get(ctx, gomock.Any()).
					Return(
						&github.PullsResponse{
							Status: http.StatusText(http.StatusBadRequest),
							Code:   http.StatusBadRequest,
							Error:  utils.ErrZeroLength.Error(),
						}, utils.NewError(
							utils.WithLogger(
								utils.ErrZeroLength.Error(),
								utils.LogPrefixInfo,
								utils.LogLevelInfo,
							),
						)).MaxTimes(1)
			},
		},
		{
			name: "pulls.Get() returns a failed *github.PullsResponse struct if map key is empty",
			args: args{
				ctx: ctx,
				opts: []utils.Field{
					utils.WithMap(map[string]string{
						"": "ksrof",
					}),
				},
			},
			want: &github.PullsResponse{
				Status: http.StatusText(http.StatusBadRequest),
				Code:   http.StatusBadRequest,
				Error:  utils.ErrZeroLength.Error(),
			},
			errStr: utils.ErrZeroLength.Error(),
			mocks: func(pulls *mock.MockPulls) {
				pulls.EXPECT().Get(ctx, gomock.Any()).
					Return(
						&github.PullsResponse{
							Status: http.StatusText(http.StatusBadRequest),
							Code:   http.StatusBadRequest,
							Error:  utils.ErrZeroLength.Error(),
						}, utils.NewError(
							utils.WithLogger(
								utils.ErrZeroLength.Error(),
								utils.LogPrefixInfo,
								utils.LogLevelInfo,
							),
						)).MaxTimes(1)
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mocks(pulls)
			issue, err := pulls.Get(tc.args.ctx, tc.args.opts)
			if err != nil {
				assert.EqualError(t, err, tc.errStr)
				assert.Equal(t, tc.want, issue)
				return
			}

			assert.ErrorIs(t, err, tc.wantErr)
			assert.Equal(t, tc.want, issue)
		})
	}
}
