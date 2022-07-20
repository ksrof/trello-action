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
		validations []utils.Validation
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "utils.ValidateNotEmpty() return nil if string is not empty",
			args: args{
				validations: []utils.Validation{
					utils.ValidateNotEmpty("string"),
				},
			},
			wantErr: nil,
		},
		{
			name: "utils.ValidateRegexp() return nil if string matches regexp",
			args: args{
				validations: []utils.Validation{
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
				validations: []utils.Validation{
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
				validations: []utils.Validation{
					utils.ValidateNotEmpty(float64(3.14)),
				},
			},
			wantErr: utils.ErrInvalidType,
		},
		{
			name: "utils.ValidateRegexp() return error if type is invalid",
			args: args{
				validations: []utils.Validation{
					utils.ValidateRegexp(
						*regexp.MustCompile("^[0-9]{6}"),
						float64(3.14),
					),
				},
			},
			wantErr: utils.ErrInvalidType,
		},
		{
			name: "utils.ValidateNotEmpty() return error if string is empty",
			args: args{
				validations: []utils.Validation{
					utils.ValidateNotEmpty(""),
				},
			},
			wantErr: utils.ErrEmptyValue,
		},
		{
			name: "utils.ValidateRegexp() return error if string does not match regexp",
			args: args{
				validations: []utils.Validation{
					utils.ValidateRegexp(
						*regexp.MustCompile("^[a-z]+$"),
						"STRING",
					),
				},
			},
			wantErr: utils.ErrInvalidMatch,
		},
		{
			name: "utils.ValidateRegexp() return error if int does not match regexp",
			args: args{
				validations: []utils.Validation{
					utils.ValidateRegexp(
						*regexp.MustCompile("^[0-9]{6}"),
						123,
					),
				},
			},
			wantErr: utils.ErrInvalidMatch,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := utils.Validations(tc.args.validations...)
			assert.ErrorIs(t, err, tc.wantErr)
		})
	}
}
