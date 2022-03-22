package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ksrof/gha-trello/internal"
	"github.com/ksrof/gha-trello/models"
)

// We should be getting the ID of the Issue or Pull from the github action itself
// that way we don't need to return the first element of the array, instead we just return one
// and we can perform actions on that one.

// It's dangerous to go alone! take this.
// https://docs.github.com/en/developers/webhooks-and-events/events/github-event-types#event-object-common-properties
// https://docs.github.com/en/rest/reference/issues#get-an-issue
// https://docs.github.com/en/rest/reference/pulls#get-a-pull-request

// Github API repo issues URL
var issuesURL string = "https://api.github.com/repos/%s/%s/issues/%s"

// GetIssue performs a GET request to api.github.com/repos/user/repo/issues/id.
// It returns the response body parsed in a struct.
func GetIssue(token, user, repo, id string) (issue models.Issue, err error) {
	newReq, err := internal.NewReq("GET", fmt.Sprintf(issuesURL, user, repo, id), nil)
	if err != nil {
		return issue, fmt.Errorf("failed to build new http request: %v", err)
	}

	req, err := internal.Auth(token, newReq)
	if err != nil {
		return issue, fmt.Errorf("failed to set authorization header: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return issue, fmt.Errorf("failed to send GET request to the server: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return issue, fmt.Errorf("failed to read body response: %v", err)
	}

	err = json.Unmarshal(body, &issue)
	if err != nil {
		return issue, fmt.Errorf("failed to unmarshal body response: %v", err)
	}

	return issue, nil
}
