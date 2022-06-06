package github_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/ksrof/trello-action/external/github"
)

func TestGetIssueByID(t *testing.T) {
	setTestEnv(t)

	want := http.StatusOK

	client, err := github.NewClient()
	if err != nil {
		t.Logf("received error (%v)", err)
		t.Fail()
	}

	res, err := client.GetIssueByID(context.TODO())
	if err != nil {
		switch err {
		case github.ErrCreatingNewRequest:
			t.Logf("expected error (%v), received error (%v)", github.ErrCreatingNewRequest, err)
			t.Fail()
		case github.ErrDoingRequest:
			t.Logf("expected error (%v), received error (%v)", github.ErrDoingRequest, err)
			t.Fail()
		default:
			t.Logf("received error (%v)", err)
			t.Fail()
		}
	}

	if res != nil {
		switch res.StatusCode {
		case http.StatusOK:
			t.Logf("expected code (%v), received code (%v)", want, res.StatusCode)
		case http.StatusNotFound:
			t.Logf("expected code (%v), received code (%v)", want, res.StatusCode)
			t.Fail()
		case http.StatusGone:
			t.Logf("expected code (%v), received code (%v)", want, res.StatusCode)
			t.Fail()
		}
	}
}

func TestGetPullByID(t *testing.T) {
	setTestEnv(t)

	want := http.StatusOK

	client, err := github.NewClient(
		github.WithEvent("pulls"),
		github.WithID("7"),
	)
	if err != nil {
		t.Logf("received error (%v)", err)
		t.Fail()
	}

	res, err := client.GetIssueByID(context.TODO())
	if err != nil {
		switch err {
		case github.ErrCreatingNewRequest:
			t.Logf("expected error (%v), received error (%v)", github.ErrCreatingNewRequest, err)
			t.Fail()
		case github.ErrDoingRequest:
			t.Logf("expected error (%v), received error (%v)", github.ErrDoingRequest, err)
			t.Fail()
		default:
			t.Logf("received error (%v)", err)
			t.Fail()
		}
	}

	if res != nil {
		switch res.StatusCode {
		case http.StatusOK:
			t.Logf("expected code (%v), received code (%v)", want, res.StatusCode)
		case http.StatusNotFound:
			t.Logf("expected code (%v), received code (%v)", want, res.StatusCode)
			t.Fail()
		case http.StatusInternalServerError:
			t.Logf("expected code (%v), received code (%v)", want, res.StatusCode)
			t.Fail()
		}
	}
}

func TestGetLabelsFromIssue(t *testing.T) {
	setTestEnv(t)

	want := http.StatusOK

	client, err := github.NewClient()
	if err != nil {
		t.Logf("received error (%v)", err)
		t.Fail()
	}

	res, err := client.GetIssueByID(context.TODO())
	if err != nil {
		switch err {
		case github.ErrCreatingNewRequest:
			t.Logf("expected error (%v), received error (%v)", github.ErrCreatingNewRequest, err)
			t.Fail()
		case github.ErrDoingRequest:
			t.Logf("expected error (%v), received error (%v)", github.ErrDoingRequest, err)
			t.Fail()
		default:
			t.Logf("received error (%v)", err)
			t.Fail()
		}
	}

	if res != nil {
		switch res.StatusCode {
		case http.StatusOK:
			t.Logf("expected code (%v), received code (%v)", want, res.StatusCode)
		case http.StatusGone:
			t.Logf("expected code (%v), received code (%v)", want, res.StatusCode)
			t.Fail()
		}
	}
}
