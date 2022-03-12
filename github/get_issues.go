package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ksrof/gha-trello/internal"
	"github.com/ksrof/gha-trello/models"
)

// Github API repo issues URL
var issuesURL string = "https://api.github.com/repos/%s/%s/issues"

// GetIssues performs a GET request to api.github.com/repos/user/repo/issues.
// It returns the response body parsed in a struct.
func GetIssues(token, user, repo string) (issue models.Issue, err error) {
	newReq, err := internal.NewReq("GET", fmt.Sprintf(issuesURL, user, repo), nil)
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
