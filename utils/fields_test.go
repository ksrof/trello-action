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

func TestNewFieldsMapper(t *testing.T) {
	type args struct {
		opts []utils.Field
	}

	tests := []struct {
		name    string
		args    args
		want    map[string]string
		errStr  string
		wantErr error
	}{
		{
			name: "utils.WithMap() return nil if key and value are not empty",
			args: args{
				opts: []utils.Field{
					utils.WithMap(map[string]string{
						"username": "ksrof",
					}),
				},
			},
			want: map[string]string{
				"username": "ksrof",
			},
			wantErr: nil,
		},
		{
			name: "utils.WithMap() return error if map is empty",
			args: args{
				opts: []utils.Field{
					utils.WithMap(map[string]string{}),
				},
			},
			errStr: utils.ErrZeroLength.Error(),
		},
		{
			name: "utils.WithMap() return error if key is empty",
			args: args{
				opts: []utils.Field{
					utils.WithMap(map[string]string{
						"": "ksrof",
					}),
				},
			},
			errStr: utils.ErrZeroLength.Error(),
		},
		{
			name: "utils.WithMap() return error if value is empty",
			args: args{
				opts: []utils.Field{
					utils.WithMap(map[string]string{
						"username": "",
					}),
				},
			},
			errStr: utils.ErrZeroLength.Error(),
		},
		{
			name: "utils.NewFieldsMapper() return error if there are no options",
			args: args{
				[]utils.Field{},
			},
			errStr: utils.ErrZeroLength.Error(),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fields, err := utils.NewFieldsMapper(tc.args.opts...)
			if err != nil {
				assert.EqualError(t, err, tc.errStr)
				return
			}

			assert.ErrorIs(t, err, tc.wantErr)
			assert.Equal(t, tc.want, fields)
		})
	}
}
