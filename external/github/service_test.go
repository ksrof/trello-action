package github_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ksrof/trello-action/external/github"
	"github.com/stretchr/testify/assert"
)

func TestGetIssueByID(t *testing.T) {
	setTestEnv(t)

	mux := http.NewServeMux()

	mux.HandleFunc("/repos/ksrof/trello-action/issues/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(`{"hello": "world"}`)
		if err != nil {
			panic(err)
		}
	})

	s := httptest.NewServer(mux)
	defer s.Close()

	c, _ := github.NewClient(
		github.WithHost(s.URL),
	)

	_, err := c.GetIssueByID(context.TODO())
	if err != nil {
		assert.Error(t, err)
		return
	}
}

func TestGetPullByID(t *testing.T) {
	setTestEnv(t)
	t.Setenv("GH_EVENT", "pulls") // Override `issues` event

	mux := http.NewServeMux()

	mux.HandleFunc("/repos/ksrof/trello-action/pulls/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	s := httptest.NewServer(mux)
	defer s.Close()

	c, _ := github.NewClient(
		github.WithHost(s.URL),
	)

	_, err := c.GetPullByID(context.TODO())
	if err != nil {
		assert.Error(t, err)
		return
	}
}

func TestGetIssueLabels(t *testing.T) {
	setTestEnv(t)

	mux := http.NewServeMux()

	mux.HandleFunc("/repos/ksrof/trello-action/issues/1/labels", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	s := httptest.NewServer(mux)
	defer s.Close()

	c, _ := github.NewClient(
		github.WithHost(s.URL),
	)

	_, err := c.GetIssueLabels(context.TODO())
	if err != nil {
		assert.Error(t, err)
		return
	}
}
