package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ksrof/gha-trello/github"
	"github.com/ksrof/gha-trello/internal"
	"github.com/ksrof/gha-trello/models"
)

// Trello API card URL
var cardURL string = "https://api.trello.com/1/cards"

// create performs a POST request to api.trello.com/1/cards.
// It creates a new card on a trello board with the data provided by a new Issue or PR got from Github.
func create(env models.Env) error {
	newReq, err := internal.NewReq("POST", cardURL, nil)
	if err != nil {
		return fmt.Errorf("failed to build new http request: %v", err)
	}

	switch strings.ToLower(env.Action) {
	case "issue":
		issue, err := github.GetIssues(env.GithubToken, env.GithubUser, env.GithubRepo)
		if err != nil {
			return fmt.Errorf("failed to get issues: %v", err)
		}

		params := models.Params{
			IDList: env.TrelloIDList,
			Key:    env.TrelloKey,
			Token:  env.TrelloToken,
			Title:  issue[0].Title,
			Number: strconv.Itoa(issue[0].Number),
			URL:    issue[0].HTMLURL,
		}

		req, err := internal.Params(params, newReq)
		if err != nil {
			return fmt.Errorf("failed to set query parameters: %v", err)
		}

		_, err = http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("failed to send POST request to the server: %v", err)
		}
	case "pull":
		pull, err := github.GetPulls(env.GithubToken, env.GithubUser, env.GithubRepo)
		if err != nil {
			return fmt.Errorf("failed to get pulls: %v", err)
		}

		params := models.Params{
			IDList: env.TrelloIDList,
			Key:    env.TrelloKey,
			Token:  env.TrelloToken,
			Title:  pull[0].Title,
			Number: strconv.Itoa(pull[0].Number),
			URL:    pull[0].HTMLURL,
		}

		req, err := internal.Params(params, newReq)
		if err != nil {
			return fmt.Errorf("failed to set query parameters: %v", err)
		}

		_, err = http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("failed to send POST request to the server: %v", err)
		}

		_, err = http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("failed to send POST request to the server: %v", err)
		}
	}

	return nil
}

func main() {
	env, err := internal.Env()
	if err != nil {
		log.Fatalf("failed to get environment variables: %v", err)
	}

	err = create(env)
	if err != nil {
		log.Fatalf("failed to create new trello card: %v", err)
	}
}
