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

		title := fmt.Sprint(issue["title"])
		url := fmt.Sprint(issue["html_url"])

		err = trello.CreateCard(*env, title, url)
		if err != nil {
			log.Printf("failed to create card: %v", err)
		}
	}

	if env.Github.Event == "pull_request" {
		pull, err := github.GetPull(*env)
		if err != nil {
			log.Printf("failed to get pull: %v", err)
		}

		title := fmt.Sprint(pull["title"])
		url := fmt.Sprint(pull["html_url"])

		err = trello.CreateCard(*env, title, url)
		if err != nil {
			log.Printf("failed to create card: %v", err)
		}
	}
}
