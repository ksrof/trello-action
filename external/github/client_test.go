package github

import "testing"

func TestNewClient(t *testing.T) {
	type args struct {
		opts []string
	}

	tests := []struct {
		name string
		args args
		want struct{}
		err  error
	}{
		{
			name: "Test new client",
			args: args{
				opts: []string{"op1", "op2", "op3"},
			},
			want: struct{}{},
			err:  nil,
		},
		{
			name: "Test new client with empty options",
			args: args{
				opts: []string{},
			},
			err: errEmptyOptions,
		},
	}

	for _, test := range tests {
		got, err := NewAPIClient(test.args.opts)
		if err != nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
			assert.Equal(t, test.err, err)
			return
		}

		if err == nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.want, got)
			assert.Equal(t, test.want, got)
			return
		}

		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.want, got)
	}
}

func TestNewClient_WithHost(t *testing.T) {
	type args struct {
		opts []string
	}

	tests := []struct {
		name string
		args args
		want struct{}
		err  error
	}{
		{
			name: "Test new client with host",
			args: args{
				opts: []string{"opt1"},
			},
			want: struct{}{},
			err:  nil,
		},
		{
			name: "Test new client with empty host",
			args: args{
				opts: []string{"opt1"},
			},
			want: struct{}{},
			err:  errEmptyHost,
		},
		{
			name: "Test new client with invalid host",
			args: args{
				opts: []string{"opt1"},
			},
			want: struct{}{},
			err:  errInvalidHost,
		},
	}

	for _, test := range tests {
		got, err := NewAPIClient(test.args.opts)
		if err != nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
			assert.Equal(t, test.err, err)
			return
		}

		if err == nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.want, got)
			assert.Equal(t, test.want, got)
			return
		}

		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.want, got)
	}
}

func TestNewClient_WithUser(t *testing.T) {
	type args struct {
		opts []string
	}

	tests := []struct {
		name string
		args args
		want struct{}
		err  error
	}{
		{
			name: "Test new client with user",
			args: args{
				opts: []string{"opt2"},
			},
			want: struct{}{},
			err:  nil,
		},
		{
			name: "Test new client with empty user",
			args: args{
				opts: []string{"opt2"},
			},
			want: struct{}{},
			err:  errEmptyUser,
		},
		{
			name: "Test new client with invalid user",
			args: args{
				opts: []string{"opt2"},
			},
			want: struct{}{},
			err:  errInvalidUser,
		},
	}

	for _, test := range tests {
		got, err := NewAPIClient(test.args.opts)
		if err != nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
			assert.Equal(t, test.err, err)
			return
		}

		if err == nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.want, got)
			assert.Equal(t, test.want, got)
			return
		}

		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.want, got)
	}
}

func TestNewClient_WithRepo(t *testing.T) {
	type args struct {
		opts []string
	}

	tests := []struct {
		name string
		args args
		want struct{}
		err  error
	}{
		{
			name: "Test new client with repo",
			args: args{
				opts: []string{"opt3"},
			},
			want: struct{}{},
			err:  nil,
		},
		{
			name: "Test new client with empty repo",
			args: args{
				opts: []string{"opt3"},
			},
			want: struct{}{},
			err:  errEmptyRepo,
		},
	}

	for _, test := range tests {
		got, err := NewAPIClient(test.args.opts)
		if err != nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
			assert.Equal(t, test.err, err)
			return
		}

		if err == nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.want, got)
			assert.Equal(t, test.want, got)
			return
		}

		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.want, got)
	}
}

func TestNewClient_WithEvent(t *testing.T) {
	type args struct {
		opts []string
	}

	tests := []struct {
		name string
		args args
		want struct{}
		err  error
	}{
		{
			name: "Test new client with event",
			args: args{
				opts: []string{"opt4"},
			},
			want: struct{}{},
			err:  nil,
		},
		{
			name: "Test new client with empty event",
			args: args{
				opts: []string{"opt4"},
			},
			want: struct{}{},
			err:  errEmptyEvent,
		},
		{
			name: "Test new client with invalid event",
			args: args{
				opts: []string{"opt4"},
			},
			want: struct{}{},
			err:  errInvalidEvent,
		},
	}

	for _, test := range tests {
		got, err := NewAPIClient(test.args.opts)
		if err != nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
			assert.Equal(t, test.err, err)
			return
		}

		if err == nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.want, got)
			assert.Equal(t, test.want, got)
			return
		}

		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.want, got)
	}
}

func TestNewClient_WithID(t *testing.T) {
	type args struct {
		opts []string
	}

	tests := []struct {
		name string
		args args
		want struct{}
		err  error
	}{
		{
			name: "Test new client with ID",
			args: args{
				opts: []string{"opt5"},
			},
			want: struct{}{},
			err:  nil,
		},
		{
			name: "Test new client with empty ID",
			args: args{
				opts: []string{"opt5"},
			},
			want: struct{}{},
			err:  errEmptyID,
		},
		{
			name: "Test new client with invalid ID",
			args: args{
				opts: []string{"opt5"},
			},
			want: struct{}{},
			err:  errInvalidID,
		},
	}

	for _, test := range tests {
		got, err := NewAPIClient(test.args.opts)
		if err != nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
			assert.Equal(t, test.err, err)
			return
		}

		if err == nil {
			t.Logf("Name: %s, want (%v), got (%v)", test.name, test.want, got)
			assert.Equal(t, test.want, got)
			return
		}

		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.err, err)
		t.Errorf("Name: %s, want (%v), got (%v)", test.name, test.want, got)
	}
}
