package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ksrof/gha-trello/internal/models"
	"github.com/ksrof/gha-trello/internal/utils"
)

// GetPull performs a GET request to api.github.com/repos/user/repo/pulls/id.
// It returns the response object unmarshaled inside a map.
func GetPull(env models.Env) (map[string]interface{}, error) {
	reqURL, err := utils.ParseURL("https://api.github.com", fmt.Sprintf("/repos/%s/%s/pulls/%s", env.Github.User, env.Github.Repo, env.Github.ID))
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %v", err)
	}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build new http request: %v", err)
	}

	err = utils.ValidateToken(env.Github.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %v", err)
	}

	// Set the Token in the Authorization Header field.
	req.Header.Set("Authorization", fmt.Sprintf("token %s", env.Github.Token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send GET request to the server: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body response: %v", err)
	}

	// response holds all data required to represent a pull_request.
	var response map[string]interface{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body response: %v", err)
	}

	return response, nil
}
