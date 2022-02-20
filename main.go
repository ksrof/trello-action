package main

// TODO's
// get the issues from a specific repository
// get the pull request from a specific repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/ksrof/gha-trello/models"
	"github.com/ksrof/gha-trello/utils"
)

var (
	baseURL string = "https://api.github.com"
	token   string = os.Getenv("GH_TOKEN")
	user    string = os.Getenv("GH_USER")
	repo    string = os.Getenv("GH_REPO")
)

// getIssues performs a GET requests to api.github.com/repos/user/repo/issues
// with an Authorization token and returns the response as a struct.
func getIssues(ghToken, ghUser, ghRepo string) (models.Issue, error) {
	// Parse the baseURL with the env values
	reqURL, err := url.Parse(fmt.Sprintf("%s/repos/%s/%s/issues", baseURL, ghUser, ghRepo))
	if err != nil {
		return models.Issue{}, fmt.Errorf("unable to parse url: %v", err)
	}

	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return models.Issue{}, fmt.Errorf("unable to create new request: %v", err)
	}

	// Set the Authorization header using the
	// Github Personal Access Token
	token, err := utils.SetAuth(ghToken)
	if err != nil {
		return models.Issue{}, fmt.Errorf("unable to get token: %v", err)
	}
	req.Header.Set("Authorization", token)

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
func getPulls(ghToken, ghUser, ghRepo string) (models.Pull, error) {
	// Parse the baseURL with the env values
	reqURL, err := url.Parse(fmt.Sprintf("%s/repos/%s/%s/pulls", baseURL, ghUser, ghRepo))
	if err != nil {
		return models.Pull{}, fmt.Errorf("unable to parse url: %v", err)
	}

	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return models.Pull{}, fmt.Errorf("unable to create new request: %v", err)
	}

	// Set the Authorization header using the
	// Github Personal Access Token
	token, err := utils.SetAuth(ghToken)
	if err != nil {
		return models.Pull{}, fmt.Errorf("unable to get token: %v", err)
	}
	req.Header.Set("Authorization", token)

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

func main() {
	issue, err := getIssues(token, user, repo)
	if err != nil {
		log.Fatalf("unable to get issues from repository: %v", err)
	}

	fmt.Printf("Issue Title: %s", issue[0].Title)

	pull, err := getPulls(token, user, repo)
	if err != nil {
		log.Fatalf("unable to get pulls from repository: %v", err)
	}

	fmt.Printf("Pull Title: %s", pull[0].Title)
}
