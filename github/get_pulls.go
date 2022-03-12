package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ksrof/gha-trello/internal"
	"github.com/ksrof/gha-trello/models"
)

// Github API repo pulls URL
var pullsURL string = "https://api.github.com/repos/%s/%s/pulls"

// GetPulls performs a GET request to api.github.com/repos/user/repo/pulls.
// It returns the response body parsed in a struct.
func GetPulls(token, user, repo string) (pull models.Pull, err error) {
	newReq, err := internal.NewReq("GET", fmt.Sprintf(pullsURL, user, repo), nil)
	if err != nil {
		return pull, fmt.Errorf("failed to build new http request: %v", err)
	}

	req, err := internal.Auth(token, newReq)
	if err != nil {
		return pull, fmt.Errorf("failed to set authorization header: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return pull, fmt.Errorf("failed to send GET request to the server: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return pull, fmt.Errorf("failed to read body response: %v", err)
	}

	err = json.Unmarshal(body, &pull)
	if err != nil {
		return pull, fmt.Errorf("failed to unmarshal body response: %v", err)
	}

	return pull, nil
}
