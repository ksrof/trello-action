package github_test

import (
	"testing"

	"github.com/ksrof/trello-action/external/github"
	"github.com/ksrof/trello-action/utils"
	"github.com/stretchr/testify/assert"
)

func TestOAuth(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    string
		wantErr error
	}{
		{
			name:    "returns an error if the value length is zero",
			value:   "",
			want:    "",
			wantErr: utils.LogError(utils.ErrZeroLength.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
		},
		{
			name:    "returns an error if the value does not match the regexp",
			value:   "abc123",
			want:    "",
			wantErr: utils.LogError(utils.ErrInvalidMatch.Error(), utils.LogPrefixInfo, utils.LogLevelInfo),
		},
		{
			name:    "returns a token",
			value:   "ghp_F41fR24d9Tvn3YRzC7GdOPAfhgjBzP5MLP7c",
			want:    "ghp_F41fR24d9Tvn3YRzC7GdOPAfhgjBzP5MLP7c",
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			token, err := github.Basic(tc.value)
			if err != nil {
				assert.ErrorContains(t, err, tc.wantErr.Error())
				return
			}

			assert.ErrorIs(t, err, tc.wantErr)
			assert.Equal(t, tc.want, token)
		})
	}
}
