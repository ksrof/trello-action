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
		want    string
		wantErr error
	}{
		{
			name: "github.WithUser() return nil if user is valid",
			args: args{
				opts: []github.Option{
					github.WithUser("ksrof"),
				},
			},
			wantErr: nil,
		},
		{
			name: "github.WithRepo() return nil if repo is valid",
			args: args{
				opts: []github.Option{
					github.WithRepo("trello-action"),
				},
			},
			wantErr: nil,
		},
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
			name: "github.WithUser() return error if user is empty",
			args: args{
				opts: []github.Option{
					github.WithUser(""),
				},
			},
			wantErr: utils.ErrEmptyValue,
		},
		{
			name: "github.WithRepo() return error if repo is empty",
			args: args{
				opts: []github.Option{
					github.WithRepo(""),
				},
			},
			wantErr: utils.ErrEmptyValue,
		},
		{
			name: "github.WithToken() return error if token is empty",
			args: args{
				opts: []github.Option{
					github.WithToken(""),
				},
			},
			wantErr: utils.ErrEmptyValue,
		},
		{
			name: "github.WithToken() return error if token is invalid",
			args: args{
				opts: []github.Option{
					github.WithToken("invalid"),
				},
			},
			wantErr: utils.ErrInvalidMatch,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, err := github.NewAuth(tc.args.opts...)
			assert.ErrorIs(t, err, tc.wantErr)
		})
	}
}

func TestBasic(t *testing.T) {
	type args struct {
		opts []github.Option
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "github.Basic() return string if token is valid",
			args: args{
				opts: []github.Option{
					github.WithToken("ghp_F41fR24d9Tvn3YRzC7GdOPAfhgjBzP5MLP7c"),
				},
			},
			want: "ghp_F41fR24d9Tvn3YRzC7GdOPAfhgjBzP5MLP7c",
		},
		{
			name: "github.Basic() return error if token is empty",
			args: args{
				opts: []github.Option{
					github.WithToken(""),
				},
			},
			wantErr: utils.ErrEmptyValue,
		},
		{
			name: "github.Basic() return error if token is invalid",
			args: args{
				opts: []github.Option{
					github.WithToken("invalid"),
				},
			},
			wantErr: utils.ErrInvalidMatch,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			auth, err := github.NewAuth(tc.args.opts...)
			if err != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}

			result := auth.Basic()
			assert.Equal(t, tc.want, result)
		})
	}
}
