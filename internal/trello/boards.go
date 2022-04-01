package trello

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ksrof/trello-action/internal/models"
	"github.com/ksrof/trello-action/internal/utils"
)

// GetCards performs a GET request to api.trello.com/1/boards/id/cards.
// It returns the response object unmarshaled inside a map.
func GetCards(env models.Env) ([]map[string]interface{}, error) {
	reqURL, err := utils.ParseURL("https://api.trello.com/1", fmt.Sprintf("/boards/%s/cards", env.Trello.IDBoard))
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %v", err)
	}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build new http request: %v", err)
	}

	// Set query parameters.
	query := req.URL.Query()
	query.Add("key", env.Trello.Key)
	query.Add("token", env.Trello.Token)

	// Encode values into the URL.
	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send GET request to the server: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body response: %v", err)
	}

	// response holds all data required to represent the cards.
	var response []map[string]interface{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body response: %v", err)
	}

	return response, nil
}

// GetLists performs a GET request to api.trello.com/1/boards/id/lists.
// It returns the response object unmarshaled inside a map.
func GetLists(env models.Env) ([]map[string]interface{}, error) {
	reqURL, err := utils.ParseURL("https://api.trello.com/1", fmt.Sprintf("/boards/%s/lists", env.Trello.IDBoard))
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %v", err)
	}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build new http request: %v", err)
	}

	// Set query parameters.
	query := req.URL.Query()
	query.Add("key", env.Trello.Key)
	query.Add("token", env.Trello.Token)

	// Encode values into the URL.
	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send GET request to the server: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body response: %v", err)
	}

	// response holds all data required to represent the lists.
	var response []map[string]interface{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body response: %v", err)
	}

	return response, nil
}
