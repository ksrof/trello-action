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
		log.Fatalf("failed to get load environment: %v", err)
	}

	if env.Github.Event == "issues" {
		response, err := github.GetIssue(*env)
		if err != nil {
			log.Printf("failed to get issue: %v", err)
		}

		title := fmt.Sprint(response["title"])
		url := fmt.Sprint(response["html_url"])

		err = trello.CreateCard(*env, title, url)
		if err != nil {
			log.Printf("failed to create card: %v", err)
		}
	}

	if env.Github.Event == "pull_request" {
		response, err := github.GetPull(*env)
		if err != nil {
			log.Printf("failed to get pull: %v", err)
		}

		title := fmt.Sprint(response["title"])
		url := fmt.Sprint(response["html_url"])

		err = trello.CreateCard(*env, title, url)
		if err != nil {
			log.Printf("failed to create card: %v", err)
		}
	}
}
