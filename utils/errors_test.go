/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

package utils_test

import (
	"testing"

	"github.com/ksrof/trello-action/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	type args struct {
		opts []utils.Errors
	}

	tests := []struct {
		name   string
		args   args
		errStr string
	}{
		{
			name: "utils.NewError() return error if there are no options",
			args: args{
				opts: []utils.Errors{},
			},
			errStr: utils.ErrZeroLength.Error(),
		},
		{
			name: "utils.WithLogger() return error if string is empty",
			args: args{
				opts: []utils.Errors{
					utils.WithLogger("", "", 1),
				},
			},
			errStr: utils.ErrZeroLength.Error(),
		},
		{
			name: "utils.WithLogger() return error if string is not empty and print log level one",
			args: args{
				opts: []utils.Errors{
					utils.WithLogger("level one", "", 1),
				},
			},
			errStr: "level one",
		},
		{
			name: "utils.WithLogger() return error if string is not empty and print default log",
			args: args{
				opts: []utils.Errors{
					utils.WithLogger("default", "", 5),
				},
			},
			errStr: "default",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := utils.NewError(tc.args.opts...)
			assert.EqualError(t, err, tc.errStr)
		})
	}
}
