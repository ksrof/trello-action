package github

import (
	"net/http"
	"testing"
)

func TestService_GetIssueByID(t *testing.T) {
	type args struct {
		opts []string
	}

	tests := []struct {
		name string
		args args
		want int
		err  error
	}{
		{
			name: "Test get issue by ID",
			args: args{
				opts: []string{"opt1", "opt2", "opt3"},
			},
			want: http.StatusOK,
			err:  nil,
		},
		{
			name: "Test get issue by ID not found",
			args: args{
				opts: []string{"opt1", "opt2", "opt3"},
			},
			want: http.StatusNotFound,
			err:  errResourceNotFound,
		},
		{
			name: "Test get issue by ID gone",
			args: args{
				opts: []string{"opt1", "opt2", "opt3"},
			},
			want: http.StatusGone,
			err:  errGone,
		},
	}

	for _, test := range tests {
		client, _ := NewAPIClient(
			WithHost("https://api.github.com/"),
			WithUser("ksrof"),
			WithRepo("trello-action-test"),
			WithEvent("issues"),
			WithID("12"),
		)

		got, err := client.GetIssueByID()
		if err != nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.want, got.StatusCode)
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
			assert.Equal(t, test.want, got.StatusCode)
			assert.Equal(t, test.err, err)
			return
		}

		if err == nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.want, got.StatusCode)
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
			assert.Equal(t, test.want, got.StatusCode)
			assert.Equal(t, test.err, err)
			return
		}

		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.want, got.StatusCode)
		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
	}
}

func TestService_GetPullByID(t *testing.T) {
	type args struct {
		opts []string
	}

	tests := []struct {
		name string
		args args
		want int
		err  error
	}{
		{
			name: "Test get pull by ID",
			args: args{
				opts: []string{"opt1", "opt2", "opt3"},
			},
			want: http.StatusOK,
			err:  nil,
		},
		{
			name: "Test get pull by ID not found",
			args: args{
				opts: []string{"opt1", "opt2", "opt3"},
			},
			want: http.StatusNotFound,
			err:  errResourceNotFound,
		},
		{
			name: "Test get pull by ID internal error",
			args: args{
				opts: []string{"opt1", "opt2", "opt3"},
			},
			want: http.StatusInternalServerError,
			err:  errInternalError,
		},
	}

	for _, test := range tests {
		client, _ := NewAPIClient(
			WithHost("https://api.github.com/"),
			WithUser("ksrof"),
			WithRepo("trello-action-test"),
			WithEvent("issues"),
			WithID("12"),
		)

		got, err := client.GetPullByID()
		if err != nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.want, got.StatusCode)
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
			assert.Equal(t, test.want, got.StatusCode)
			assert.Equal(t, test.err, err)
			return
		}

		if err == nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.want, got.StatusCode)
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
			assert.Equal(t, test.want, got.StatusCode)
			assert.Equal(t, test.err, err)
			return
		}

		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.want, got.StatusCode)
		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
	}
}

func TestService_GetLabelsFromIssue(t *testing.T) {
	type args struct {
		opts []string
	}

	tests := []struct {
		name string
		args args
		want int
		err  error
	}{
		{
			name: "Test get labels from issue",
			args: args{
				opts: []string{"opt1", "opt2", "opt3"},
			},
			want: http.StatusOK,
			err:  nil,
		},
		{
			name: "Test get labels from issue gone",
			args: args{
				opts: []string{"opt1", "opt2", "opt3"},
			},
			want: http.StatusGone,
			err:  errGone,
		},
	}

	for _, test := range tests {
		client, _ := NewAPIClient(
			WithHost("https://api.github.com/"),
			WithUser("ksrof"),
			WithRepo("trello-action-test"),
			WithEvent("issues"),
			WithID("12"),
		)

		got, err := client.GetPullByID()
		if err != nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.want, got.StatusCode)
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
			assert.Equal(t, test.want, got.StatusCode)
			assert.Equal(t, test.err, err)
			return
		}

		if err == nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.want, got.StatusCode)
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
			assert.Equal(t, test.want, got.StatusCode)
			assert.Equal(t, test.err, err)
			return
		}

		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.want, got.StatusCode)
		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
	}
}
