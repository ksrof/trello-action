package trello

import (
	"fmt"
	"net/http"

	"github.com/ksrof/trello-action/internal/models"
	"github.com/ksrof/trello-action/internal/utils"
)

// CreateCard performs a POST request to api.trello.com/1/cards.
func CreateCard(env models.Env, title, url string) error {
	reqURL, err := utils.ParseURL("https://api.trello.com/1", "/cards")
	if err != nil {
		return fmt.Errorf("failed to parse url: %v", err)
	}

	req, err := http.NewRequest("POST", reqURL, nil)
	if err != nil {
		return fmt.Errorf("failed to build new http request: %v", err)
	}

	// Set query parameters.
	query := req.URL.Query()
	query.Add("idList", env.Trello.IDList)
	query.Add("name", title)
	query.Add("urlSource", url)
	query.Add("key", env.Trello.Key)
	query.Add("token", env.Trello.Token)

	// Encode values into the URL.
	req.URL.RawQuery = query.Encode()

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send POST request to the server: %v", err)
	}

	return nil
}

// UpdateCard performs a PUT request to api.trello.com/1/cards/id.
func UpdateCard(env models.Env, id, idList string) error {
	reqURL, err := utils.ParseURL("https://api.trello.com/1", fmt.Sprintf("/cards/%s", id))
	if err != nil {
		return fmt.Errorf("failed to parse url: %v", err)
	}

	req, err := http.NewRequest("PUT", reqURL, nil)
	if err != nil {
		return fmt.Errorf("failed to build new http request: %v", err)
	}

	// Set query parameters.
	query := req.URL.Query()
	query.Add("idList", idList)
	query.Add("key", env.Trello.Key)
	query.Add("token", env.Trello.Token)

	// Encode values into the URL.
	req.URL.RawQuery = query.Encode()

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send PUT request to the server: %v", err)
	}

	return nil
}

// DeleteCard performs a DELETE request to api.trello.com/1/cards/id.
func DeleteCard(env models.Env, id string) error {
	reqURL, err := utils.ParseURL("https://api.trello.com/1", fmt.Sprintf("/cards/%s", id))
	if err != nil {
		return fmt.Errorf("failed to parse url: %v", err)
	}

	req, err := http.NewRequest("DELETE", reqURL, nil)
	if err != nil {
		return fmt.Errorf("failed to build new http request: %v", err)
	}

	// Set query parameters.
	query := req.URL.Query()
	query.Add("key", env.Trello.Key)
	query.Add("token", env.Trello.Token)

	// Encode values into the URL.
	req.URL.RawQuery = query.Encode()

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send DELETE request to the server: %v", err)
	}

	return nil
}
