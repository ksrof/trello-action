package github_test

import (
	"testing"

	"github.com/ksrof/trello-action/external/github"
)

func setTestEnv(t *testing.T) {
	t.Setenv("GH_TOKEN", "xxxx")
	t.Setenv("GH_USER", "ksrof")
	t.Setenv("GH_REPO", "trello-action-test")
	t.Setenv("GH_EVENT", "issues")
	t.Setenv("GH_ID", "2")
}

func TestNewClient(t *testing.T) {
	setTestEnv(t)

	client, err := github.NewClient()
	if err != nil {
		switch err {
		case github.ErrEmptyString:
			t.Logf("expected error (%v), received error (%v)", github.ErrEmptyString, err)
			t.Fail()
		case github.ErrInvalidToken:
			t.Logf("expected error (%v), received error (%v)", github.ErrInvalidToken, err)
			t.Fail()
		case github.ErrInvalidUser:
			t.Logf("expected error (%v), received error (%v)", github.ErrInvalidUser, err)
			t.Fail()
		case github.ErrInvalidEvent:
			t.Logf("expected error (%v), received error (%v)", github.ErrInvalidEvent, err)
			t.Fail()
		case github.ErrInvalidID:
			t.Logf("expected error (%v), received error (%v)", github.ErrInvalidID, err)
			t.Fail()
		default:
			t.Logf("received error (%v)", err)
			t.Fail()
		}
	}

	if err == nil {
		t.Logf("received client (%v)", client)
	}
}

func TestWithInvalidToken(t *testing.T) {
	setTestEnv(t)

	tests := []struct {
		token   string
		wantErr error
	}{
		{token: "", wantErr: github.ErrEmptyString},
		{token: "abc123", wantErr: github.ErrInvalidToken},
	}

	for _, test := range tests {
		_, err := github.NewClient(
			github.WithToken(test.token),
		)

		if err != nil {
			switch err {
			case github.ErrEmptyString:
				t.Logf("expected error (%v), received error (%v)", test.wantErr, err)
			case github.ErrInvalidToken:
				t.Logf("expected error (%v), received error (%v)", test.wantErr, err)
			default:
				t.Logf("received error (%v)", err)
				t.Fail()
			}
		}

		if err == nil {
			t.Logf("expected error (%v), received error (%v)", test.wantErr, err)
			t.Fail()
		}
	}
}

func TestWithInvalidUser(t *testing.T) {
	setTestEnv(t)

	tests := []struct {
		user    string
		wantErr error
	}{
		{user: "", wantErr: github.ErrEmptyString},
		{user: "ksr", wantErr: github.ErrInvalidUser},
	}

	for _, test := range tests {
		_, err := github.NewClient(
			github.WithUser(test.user),
		)

		if err != nil {
			switch err {
			case github.ErrEmptyString:
				t.Logf("expected error (%v), received error (%v)", test.wantErr, err)
			case github.ErrInvalidUser:
				t.Logf("expected error (%v), received error (%v)", test.wantErr, err)
			default:
				t.Logf("received error (%v)", err)
				t.Fail()
			}
		}

		if err == nil {
			t.Logf("expected error (%v) received error (%v)", test.wantErr, err)
			t.Fail()
		}
	}
}

func TestWithInvalidRepo(t *testing.T) {
	setTestEnv(t)

	tests := []struct {
		repo    string
		wantErr error
	}{
		{repo: "", wantErr: github.ErrEmptyString},
	}

	for _, test := range tests {
		_, err := github.NewClient(
			github.WithRepo(test.repo),
		)

		if err != nil {
			t.Logf("expected error (%v), received error (%v)", test.wantErr, err)
		}

		if err == nil {
			t.Logf("expected error (%v), received error (%v)", test.wantErr, err)
			t.Fail()
		}
	}
}

func TestWithInvalidEvent(t *testing.T) {
	setTestEnv(t)

	tests := []struct {
		event   string
		wantErr error
	}{
		{event: "", wantErr: github.ErrEmptyString},
		{event: "invalid", wantErr: github.ErrInvalidEvent},
	}

	for _, test := range tests {
		_, err := github.NewClient(
			github.WithEvent(test.event),
		)

		if err != nil {
			switch err {
			case github.ErrEmptyString:
				t.Logf("expected error (%v), received error (%v)", test.wantErr, err)
			case github.ErrInvalidEvent:
				t.Logf("expected error (%v), received error (%v)", test.wantErr, err)
			default:
				t.Logf("received error (%v)", err)
				t.Fail()
			}
		}

		if err == nil {
			t.Logf("expected error (%v), received error (%v)", test.wantErr, err)
			t.Fail()
		}
	}
}

func TestWithInvalidID(t *testing.T) {
	setTestEnv(t)

	tests := []struct {
		id      string
		wantErr error
	}{
		{id: "", wantErr: github.ErrEmptyString},
		{id: "12a", wantErr: github.ErrInvalidID},
	}

	for _, test := range tests {
		_, err := github.NewClient(
			github.WithID(test.id),
		)

		if err != nil {
			switch err {
			case github.ErrEmptyString:
				t.Logf("expected error (%v), received error (%v)", test.wantErr, err)
			case github.ErrInvalidID:
				t.Logf("expected error (%v), received error (%v)", test.wantErr, err)
			default:
				t.Logf("received error (%v)", err)
				t.Fail()
			}
		}

		if err == nil {
			t.Logf("expected error (%v), received error (%v)", test.wantErr, err)
			t.Fail()
		}
	}
}
