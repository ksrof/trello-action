package main

import (
	"fmt"
	"log"

	"github.com/ksrof/trello-action/internal/configs"
	"github.com/ksrof/trello-action/internal/github"
	"github.com/ksrof/trello-action/internal/trello"
)

func main() {
	env, err := configs.Environment()
	if err != nil {
		log.Fatalf("failed to load environment: %v", err)
	}

	if env.Github.Event == "issues" {
		issue, err := github.GetIssue(*env)
		if err != nil {
			log.Printf("failed to get issue: %v", err)
		}

		issueTitle := fmt.Sprint(issue["title"])
		issueURL := fmt.Sprint(issue["html_url"])

		err = trello.CreateCard(*env, issueTitle, issueURL)
		if err != nil {
			log.Printf("failed to create card: %v", err)
		}
	}

	if env.Github.Event == "pull_request" {
		pull, err := github.GetPull(*env)
		if err != nil {
			log.Printf("failed to get pull: %v", err)
		}

		pullTitle := fmt.Sprint(pull["title"])
		pullURL := fmt.Sprint(pull["html_url"])

		err = trello.CreateCard(*env, pullTitle, pullURL)
		if err != nil {
			log.Printf("failed to create card: %v", err)
		}
	}
}
