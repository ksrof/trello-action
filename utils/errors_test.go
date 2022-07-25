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

func TestLogError(t *testing.T) {
	type args struct {
		errStr string
		prefix string
		level  int
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "returns an error if the error string length is zero",
			args: args{
				errStr: "",
				prefix: utils.LogPrefixInfo,
				level:  utils.LogLevelInfo,
			},
			wantErr: utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
		},
		{
			name: "returns an error",
			args: args{
				errStr: "an error",
				prefix: utils.LogPrefixInfo,
				level:  utils.LogLevelInfo,
			},
			wantErr: utils.LogError("an error", utils.LogPrefixInfo, utils.LogLevelInfo),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := utils.LogError(tc.args.errStr, tc.args.prefix, tc.args.level)
			assert.ErrorContains(t, err, tc.wantErr.Error())
		})
	}
}
