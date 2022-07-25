package utils_test

import (
	"testing"

	"github.com/ksrof/trello-action/utils"
	"github.com/stretchr/testify/assert"
)

func TestFieldMapper(t *testing.T) {
	tests := []struct {
		name    string
		values  map[string]string
		want    map[string]string
		wantErr error
	}{
		{
			name:    "returns error if value length is zero",
			values:  map[string]string{},
			wantErr: utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
		},
		{
			name: "returns error if map key length is zero",
			values: map[string]string{
				"": "ksrof",
			},
			wantErr: utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
		},
		{
			name: "returns error if map value length is zero",
			values: map[string]string{
				"username": "",
			},
			wantErr: utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
		},
		{
			name: "returns the fields",
			values: map[string]string{
				"username": "ksrof",
			},
			want: map[string]string{
				"username": "ksrof",
			},
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fields, err := utils.FieldMapper(tc.values)
			if err != nil {
				assert.ErrorContains(t, err, tc.wantErr.Error())
				return
			}

			assert.ErrorIs(t, err, tc.wantErr)
			assert.Equal(t, tc.values, fields)
		})
	}
}
