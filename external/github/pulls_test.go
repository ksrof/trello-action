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
		ctx    context.Context
		values map[string]string
	}

	tests := []struct {
		name    string
		args    args
		want    *github.PullsResponse
		wantErr error
		mocks   func(pulls *mock.MockPulls)
	}{
		{
			name: "returns a successful response",
			args: args{
				ctx: ctx,
				values: map[string]string{
					"username":    "ksrof",
					"repository":  "trello-action",
					"pull_id":     "30",
					"request_url": "https://api.github.com",
					"token":       "ghp_F41fR24d9Tvn3YRzC7GdOPAfhgjBzP5MLP7c",
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
			name: "returns error if map is empty",
			args: args{
				ctx:    ctx,
				values: map[string]string{},
			},
			wantErr: utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
			mocks: func(issues *mock.MockPulls) {
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
			mocks: func(issues *mock.MockPulls) {
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
			mocks: func(issues *mock.MockPulls) {
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
			tc.mocks(pulls)
			pull, err := pulls.Get(tc.args.ctx, tc.args.values)
			if err != nil {
				assert.ErrorContains(t, err, tc.wantErr.Error())
				return
			}

			assert.ErrorIs(t, err, tc.wantErr)
			assert.Equal(t, tc.want, pull)
		})
	}
}
