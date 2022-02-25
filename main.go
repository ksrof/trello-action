package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/ksrof/gha-trello/models"
	"github.com/ksrof/gha-trello/utils"
)

var (
	// Trello base url
	trURL string = "https://api.trello.com/1"
	// Github base url
	ghURL string = "https://api.github.com"
)

// getIssues performs a GET requests to api.github.com/repos/user/repo/issues
// with an Authorization token and returns the response as a struct.
func getIssues(token, user, repo string) (models.Issue, error) {
	// Parse the ghURL with the env values
	reqURL, err := url.Parse(fmt.Sprintf("%s/repos/%s/%s/issues", ghURL, user, repo))
	if err != nil {
		return models.Issue{}, fmt.Errorf("unable to parse url: %v", err)
	}

	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return models.Issue{}, fmt.Errorf("unable to create new request: %v", err)
	}

	// Set the Authorization header using the
	// Github Personal Access Token
	authToken, err := utils.SetAuth(token)
	if err != nil {
		return models.Issue{}, fmt.Errorf("unable to get token: %v", err)
	}
	req.Header.Set("Authorization", authToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.Issue{}, fmt.Errorf("unable to perform GET request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Issue{}, fmt.Errorf("unable to read body response: %v", err)
	}

	// Unmarshal the body response values into
	// the models.Issue data structure
	var issue models.Issue
	err = json.Unmarshal(body, &issue)
	if err != nil {
		return models.Issue{}, fmt.Errorf("unable to unmarshal body response: %v", err)
	}

	return issue, nil
}

// getPulls performs a GET requests to api.github.com/repos/user/repo/pulls
// with an Authorization token and returns the response as a struct.
func getPulls(token, user, repo string) (models.Pull, error) {
	// Parse the ghURL with the env values
	reqURL, err := url.Parse(fmt.Sprintf("%s/repos/%s/%s/pulls", ghURL, user, repo))
	if err != nil {
		return models.Pull{}, fmt.Errorf("unable to parse url: %v", err)
	}

	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return models.Pull{}, fmt.Errorf("unable to create new request: %v", err)
	}

	// Set the Authorization header using the
	// Github Personal Access Token
	authToken, err := utils.SetAuth(token)
	if err != nil {
		return models.Pull{}, fmt.Errorf("unable to get token: %v", err)
	}
	req.Header.Set("Authorization", authToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.Pull{}, fmt.Errorf("unable to perform GET request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Pull{}, fmt.Errorf("unable to read body response: %v", err)
	}

	// Unmarshal the body response values into
	// the models.Pull data structure
	var pull models.Pull
	err = json.Unmarshal(body, &pull)
	if err != nil {
		return models.Pull{}, fmt.Errorf("unable to unmarshal body response: %v", err)
	}

	return pull, nil
}

// createCard performs a POST request to api.trello.com/1/cards
// creating a new card on the trello board with the data
// provided by the new Issue or PR got from Github
func createCard(config utils.Config) error {
	reqURL, err := url.Parse(fmt.Sprintf("%s/cards", trURL))
	if err != nil {
		return fmt.Errorf("unable to parse url: %v", err)
	}

	req, err := http.NewRequest("POST", reqURL.String(), nil)
	if err != nil {
		return fmt.Errorf("unable to create new request: %v", err)
	}

	// Set the default parameters needed by the trello API
	query := req.URL.Query()
	query.Add("idList", config.TrelloIDList)
	query.Add("key", config.TrelloKey)
	query.Add("token", config.TrelloToken)

	// Determine which data to get based on
	// the given action type
	switch strings.ToLower(config.Action) {
	case "pull":
		pull, err := getPulls(config.GithubToken, config.GithubUser, config.GithubRepo)
		if err != nil {
			return fmt.Errorf("unable to get pull request: %v", err)
		}

		// Set the parameters related to the data shown by the card
		query.Add("name", fmt.Sprintf("%s %d", pull[0].Title, pull[0].Number))
		query.Add("urlSource", pull[0].URL)
		req.URL.RawQuery = query.Encode()

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("unable to perform POST request: %v", err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("unable to read body response: %v", err)
		}

		fmt.Printf("URL: %s", req.URL.RawQuery)
		fmt.Printf("Body: %s", string(body))
	case "issue":
		issue, err := getIssues(config.GithubToken, config.GithubUser, config.GithubRepo)
		if err != nil {
			return fmt.Errorf("unable to get issues: %v", err)
		}

		// Set the parameters related to the data shown by the card
		query.Add("name", fmt.Sprintf("%s %d", issue[0].Title, issue[0].Number))
		query.Add("urlSource", issue[0].URL)
		req.URL.RawQuery = query.Encode()

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("unable to perform POST request: %v", err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("unable to read body response: %v", err)
		}

		fmt.Printf("URL: %s", req.URL.RawQuery)
		fmt.Printf("Body: %s", string(body))
	}

	return nil
}

func main() {
	config, err := utils.SetEnv()
	if err != nil {
		log.Fatalf("unable to set environment: %v", err)
	}

	err = createCard(config)
	if err != nil {
		log.Fatalf("unable to create card: %v", err)
	}
}
