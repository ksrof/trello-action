/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

package github_test

import (
	"testing"

	"github.com/ksrof/trello-action/external/github"
	"github.com/ksrof/trello-action/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewAuth(t *testing.T) {
	type args struct {
		opts []github.Option
	}

	tests := []struct {
		name    string
		args    args
		errStr  string
		wantErr error
	}{
		{
			name: "github.WithToken() return nil if token is valid",
			args: args{
				opts: []github.Option{
					github.WithToken("ghp_F41fR24d9Tvn3YRzC7GdOPAfhgjBzP5MLP7c"),
				},
			},
			wantErr: nil,
		},
		{
			name: "github.WithToken() return error if token is empty",
			args: args{
				opts: []github.Option{
					github.WithToken(""),
				},
			},
			errStr: utils.ErrZeroLength.Error(),
		},
		{
			name: "github.WithToken() return error if token is invalid",
			args: args{
				opts: []github.Option{
					github.WithToken("invalid"),
				},
			},
			errStr: utils.ErrInvalidMatch.Error(),
		},
		{
			name: "github.NewAuth() return error if there are no options",
			args: args{
				opts: []github.Option{},
			},
			errStr: utils.ErrZeroLength.Error(),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := github.NewAuth(tc.args.opts...)
			if err != nil {
				assert.EqualError(t, err, tc.errStr)
				return
			}

			assert.ErrorIs(t, err, tc.wantErr)
		})
	}
}

func TestAuth_Basic(t *testing.T) {
	type args struct {
		opts []github.Option
	}

	tests := []struct {
		name    string
		args    args
		want    string
		errStr  string
		wantErr error
	}{
		{
			name: "github.Basic() return string if token is valid",
			args: args{
				opts: []github.Option{
					github.WithToken("ghp_F41fR24d9Tvn3YRzC7GdOPAfhgjBzP5MLP7c"),
				},
			},
			want:    "ghp_F41fR24d9Tvn3YRzC7GdOPAfhgjBzP5MLP7c",
			wantErr: nil,
		},
		{
			name: "github.Basic() return error if token is empty",
			args: args{
				opts: []github.Option{
					github.WithToken(""),
				},
			},
			errStr: utils.ErrZeroLength.Error(),
		},
		{
			name: "github.Basic() return error if token is invalid",
			args: args{
				opts: []github.Option{
					github.WithToken("invalid"),
				},
			},
			errStr: utils.ErrInvalidMatch.Error(),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			auth, err := github.NewAuth(tc.args.opts...)
			if err != nil {
				assert.EqualError(t, err, tc.errStr)
				return
			}

			assert.ErrorIs(t, err, tc.wantErr)

			result := auth.Basic()
			assert.Equal(t, tc.want, result)
		})
	}
}
