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

func TestIssues_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	issues := mock.NewMockIssues(ctrl)

	type args struct {
		ctx    context.Context
		values map[string]string
	}

	tests := []struct {
		name    string
		args    args
		want    *github.IssuesResponse
		wantErr error
		mocks   func(issues *mock.MockIssues)
	}{
		{
			name: "returns a successful response",
			args: args{
				ctx: ctx,
				values: map[string]string{
					"username":    "ksrof",
					"repository":  "trello-action",
					"issue_id":    "30",
					"request_url": "https://api.github.com/repos/%s/%s/issues/%s",
					"token":       "ghp_F41fR24d9Tvn3YRzC7GdOPAfhgjBzP5MLP7c",
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
			wantErr: nil,
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
			name: "returns error if map is empty",
			args: args{
				ctx:    ctx,
				values: map[string]string{},
			},
			wantErr: utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
			mocks: func(issues *mock.MockIssues) {
				issues.EXPECT().Get(ctx, gomock.Any()).
					Return(nil,
						utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo)).
					MaxTimes(1)
			},
		},
		{
			name: "returns error if map value length is zero",
			args: args{
				ctx: ctx,
				values: map[string]string{
					"": "ksrof",
				},
			},
			wantErr: utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
			mocks: func(issues *mock.MockIssues) {
				issues.EXPECT().Get(ctx, gomock.Any()).
					Return(nil,
						utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo)).
					MaxTimes(1)
			},
		},
		{
			name: "returns error if map key is empty",
			args: args{
				ctx:    ctx,
				values: map[string]string{},
			},
			wantErr: utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
			mocks: func(issues *mock.MockIssues) {
				issues.EXPECT().Get(ctx, gomock.Any()).
					Return(nil,
						utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo)).
					MaxTimes(1)
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mocks(issues)
			issue, err := issues.Get(tc.args.ctx, tc.args.values)
			if err != nil {
				assert.ErrorContains(t, err, tc.wantErr.Error())
				return
			}

			assert.ErrorIs(t, err, tc.wantErr)
			assert.Equal(t, tc.want, issue)
		})
	}
}

func TestIssues_GetLabels(t *testing.T) {
	ctrl := gomock.NewController(t)
	issues := mock.NewMockIssues(ctrl)

	type args struct {
		ctx    context.Context
		values map[string]string
	}

	tests := []struct {
		name    string
		args    args
		want    *github.IssuesResponse
		wantErr error
		mocks   func(issues *mock.MockIssues)
	}{
		{
			name: "returns a successful response",
			args: args{
				ctx: ctx,
				values: map[string]string{
					"username":    "ksrof",
					"repository":  "trello-action",
					"issue_id":    "30",
					"request_url": "https://api.github.com/repos/%s/%s/issues/%s/labels",
					"token":       "ghp_F41fR24d9Tvn3YRzC7GdOPAfhgjBzP5MLP7c",
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
			wantErr: nil,
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
			name: "returns an error if map is empty",
			args: args{
				ctx:    ctx,
				values: map[string]string{},
			},
			wantErr: utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
			mocks: func(issues *mock.MockIssues) {
				issues.EXPECT().GetLabels(ctx, gomock.Any()).
					Return(nil,
						utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo)).
					MaxTimes(1)
			},
		},
		{
			name: "returns an error if map value is empty",
			args: args{
				ctx: ctx,
				values: map[string]string{
					"username": "",
				},
			},
			wantErr: utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
			mocks: func(issues *mock.MockIssues) {
				issues.EXPECT().GetLabels(ctx, gomock.Any()).
					Return(nil,
						utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo)).
					MaxTimes(1)
			},
		},
		{
			name: "returns an error if map key is empty",
			args: args{
				ctx: ctx,
				values: map[string]string{
					"": "ksrof",
				},
			},
			wantErr: utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
			mocks: func(issues *mock.MockIssues) {
				issues.EXPECT().GetLabels(ctx, gomock.Any()).
					Return(nil,
						utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo)).
					MaxTimes(1)
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mocks(issues)
			issueLabels, err := issues.GetLabels(tc.args.ctx, tc.args.values)
			if err != nil {
				assert.ErrorContains(t, err, tc.wantErr.Error())
				return
			}

			assert.ErrorIs(t, err, tc.wantErr)
			assert.Equal(t, tc.want, issueLabels)
		})
	}
}
