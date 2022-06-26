package github_test

import (
	"testing"

	"github.com/ksrof/trello-action/external/github"
	"github.com/stretchr/testify/assert"
)

func setTestEnv(t *testing.T) {
	t.Setenv("GH_TOKEN", "ghp_Abc123Abc123Abc123Abc123Abc123Abc123")
	t.Setenv("GH_USER", "ksrof")
	t.Setenv("GH_REPO", "trello-action")
	t.Setenv("GH_EVENT", "issues")
	t.Setenv("GH_ID", "1")
}

func TestNewClient(t *testing.T) {
	setTestEnv(t)

	tests := []struct {
		name string
		opts []github.Option
		want error
	}{
		{
			name: "New client",
			opts: nil,
			want: nil,
		},
		{
			name: "Invalid host",
			opts: []github.Option{
				github.WithHost(""),
			},
			want: github.ErrInvalidHost,
		},
		{
			name: "Invalid token",
			opts: []github.Option{
				github.WithToken("xxx"),
			},
			want: github.ErrInvalidToken,
		},
		{
			name: "Invalid user",
			opts: []github.Option{
				github.WithUser("bad"),
			},
			want: github.ErrInvalidUser,
		},
		{
			name: "Invalid repo",
			opts: []github.Option{
				github.WithRepo(""),
			},
			want: github.ErrInvalidRepo,
		},
		{
			name: "Invalid event",
			opts: []github.Option{
				github.WithEvent("banana"),
			},
			want: github.ErrInvalidEvent,
		},
		{
			name: "Invalid identifier",
			opts: []github.Option{
				github.WithID("abc"),
			},
			want: github.ErrInvalidIdentifier,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := github.NewClient(tt.opts...)
			if err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.want.Error())
				return
			}

			assert.NotNil(t, client)
			assert.NotNil(t, client.Host())
			assert.NotNil(t, client.Token())
			assert.NotNil(t, client.User())
			assert.NotNil(t, client.Repo())
			assert.NotNil(t, client.Event())
			assert.NotNil(t, client.ID())
		})
	}
}
