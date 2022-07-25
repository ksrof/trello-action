/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

package utils_test

import (
	"regexp"
	"testing"

	"github.com/ksrof/trello-action/utils"
	"github.com/stretchr/testify/assert"
)

func TestValidations(t *testing.T) {
	type args struct {
		opts []utils.Validation
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "returns nil if string is not empty",
			args: args{
				opts: []utils.Validation{
					utils.ValidateNotEmpty("string"),
				},
			},
			wantErr: nil,
		},
		{
			name: "returns nil if string matches regexp",
			args: args{
				opts: []utils.Validation{
					utils.ValidateRegexp(
						*regexp.MustCompile("^[a-z]+$"),
						"string",
					),
				},
			},
			wantErr: nil,
		},
		{
			name: "returns nil if int matches regexp",
			args: args{
				opts: []utils.Validation{
					utils.ValidateRegexp(
						*regexp.MustCompile("^[0-9]{6}"),
						123456,
					),
				},
			},
			wantErr: nil,
		},
		{
			name: "returns error if type is invalid",
			args: args{
				opts: []utils.Validation{
					utils.ValidateNotEmpty(float64(3.14)),
				},
			},
			wantErr: utils.LogError(utils.ErrInvalidType.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
		},
		{
			name: "returns error if type is invalid",
			args: args{
				opts: []utils.Validation{
					utils.ValidateRegexp(
						*regexp.MustCompile("^[0-9]{6}"),
						float64(3.14),
					),
				},
			},
			wantErr: utils.LogError(utils.ErrInvalidType.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
		},
		{
			name: "returns error if string is empty",
			args: args{
				opts: []utils.Validation{
					utils.ValidateNotEmpty(""),
				},
			},
			wantErr: utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
		},
		{
			name: "returns error if string does not match regexp",
			args: args{
				opts: []utils.Validation{
					utils.ValidateRegexp(
						*regexp.MustCompile("^[a-z]+$"),
						"STRING",
					),
				},
			},
			wantErr: utils.LogError(utils.ErrInvalidMatch.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
		},
		{
			name: "returns error if int does not match regexp",
			args: args{
				opts: []utils.Validation{
					utils.ValidateRegexp(
						*regexp.MustCompile("^[0-9]{6}"),
						123,
					),
				},
			},
			wantErr: utils.LogError(utils.ErrInvalidMatch.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
		},
		{
			name: "returns error if there are no options",
			args: args{
				opts: []utils.Validation{},
			},
			wantErr: utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := utils.Validations(tc.args.opts...)
			if err != nil {
				assert.ErrorContains(t, err, tc.wantErr.Error())
				return
			}

			assert.ErrorIs(t, err, tc.wantErr)
		})
	}
}
