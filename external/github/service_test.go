package github_test

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ksrof/trello-action/external/github"
	"github.com/stretchr/testify/assert"
)

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func TestGetIssueByID(t *testing.T) {
	setTestEnv(t)

	mux := http.NewServeMux()
	mux.HandleFunc("/repos/ksrof/trello-action/issues/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			w.WriteHeader(http.StatusNotFound)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("issue.json"))
	})

	s := httptest.NewServer(mux)
	defer s.Close()

	c, _ := github.NewClient(
		github.WithHost(s.URL),
	)

	res, err := c.GetIssueByID(context.TODO())
	if err != nil {
		assert.Error(t, err)
	}

	data, _ := io.ReadAll(res.Body)

	assert.NotNil(t, res)
	assert.JSONEq(t, fixture("issue.json"), string(data))
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestGetPullByID(t *testing.T) {
	setTestEnv(t)
	t.Setenv("GH_EVENT", "pulls") // Override `issues` event

	mux := http.NewServeMux()
	mux.HandleFunc("/repos/ksrof/trello-action/pulls/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			w.WriteHeader(http.StatusNotFound)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("pull.json"))
	})

	s := httptest.NewServer(mux)
	defer s.Close()

	c, _ := github.NewClient(
		github.WithHost(s.URL),
	)

	res, err := c.GetPullByID(context.TODO())
	if err != nil {
		assert.Error(t, err)
	}

	data, _ := io.ReadAll(res.Body)

	assert.NotNil(t, res)
	assert.JSONEq(t, fixture("pull.json"), string(data))
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestGetIssueLabels(t *testing.T) {
	setTestEnv(t)

	mux := http.NewServeMux()
	mux.HandleFunc("/repos/ksrof/trello-action/issues/1/labels", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			w.WriteHeader(http.StatusNotFound)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("issue_labels.json"))
	})

	s := httptest.NewServer(mux)
	defer s.Close()

	c, _ := github.NewClient(
		github.WithHost(s.URL),
	)

	res, err := c.GetIssueLabels(context.TODO())
	if err != nil {
		assert.Error(t, err)
	}

	data, _ := io.ReadAll(res.Body)

	assert.NotNil(t, res)
	assert.JSONEq(t, fixture("issue_labels.json"), string(data))
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
