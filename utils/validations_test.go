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
		errStr  string
		wantErr error
	}{
		{
			name: "utils.ValidateNotEmpty() return nil if string is not empty",
			args: args{
				opts: []utils.Validation{
					utils.ValidateNotEmpty("string"),
				},
			},
			wantErr: nil,
		},
		{
			name: "utils.ValidateRegexp() return nil if string matches regexp",
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
			name: "utils.ValidateRegexp() return nil if int matches regexp",
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
			name: "utils.ValidateNotEmpty() return error if type is invalid",
			args: args{
				opts: []utils.Validation{
					utils.ValidateNotEmpty(float64(3.14)),
				},
			},
			errStr: utils.ErrInvalidType.Error(),
		},
		{
			name: "utils.ValidateRegexp() return error if type is invalid",
			args: args{
				opts: []utils.Validation{
					utils.ValidateRegexp(
						*regexp.MustCompile("^[0-9]{6}"),
						float64(3.14),
					),
				},
			},
			errStr: utils.ErrInvalidType.Error(),
		},
		{
			name: "utils.ValidateNotEmpty() return error if string is empty",
			args: args{
				opts: []utils.Validation{
					utils.ValidateNotEmpty(""),
				},
			},
			errStr: utils.ErrZeroLength.Error(),
		},
		{
			name: "utils.ValidateRegexp() return error if string does not match regexp",
			args: args{
				opts: []utils.Validation{
					utils.ValidateRegexp(
						*regexp.MustCompile("^[a-z]+$"),
						"STRING",
					),
				},
			},
			errStr: utils.ErrInvalidMatch.Error(),
		},
		{
			name: "utils.ValidateRegexp() return error if int does not match regexp",
			args: args{
				opts: []utils.Validation{
					utils.ValidateRegexp(
						*regexp.MustCompile("^[0-9]{6}"),
						123,
					),
				},
			},
			errStr: utils.ErrInvalidMatch.Error(),
		},
		{
			name: "utils.Validations() return error if there are no options",
			args: args{
				opts: []utils.Validation{},
			},
			errStr: utils.ErrZeroLength.Error(),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := utils.Validations(tc.args.opts...)
			if err != nil {
				assert.EqualError(t, err, tc.errStr)
				return
			}

			assert.ErrorIs(t, err, tc.wantErr)
		})
	}
}
