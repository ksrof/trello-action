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
			name: "utils.WithMap() return error if key is empty",
			args: args{
				opts: []utils.Field{
					utils.WithMap(map[string]string{
						"": "ksrof",
					}),
				},
			},
			wantErr: utils.ErrEmptyValue,
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
			wantErr: utils.ErrEmptyValue,
		},
		{
			name: "utils.NewFieldsMapper() return error if there are no options",
			args: args{
				[]utils.Field{},
			},
			wantErr: utils.ErrEmptyOptions,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fields, err := utils.NewFieldsMapper(tc.args.opts...)
			if err != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}

			assert.Equal(t, tc.want, fields)
		})
	}
}
