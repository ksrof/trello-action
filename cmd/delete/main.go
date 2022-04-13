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
		log.Fatalf("faled to load environment: %v", err)
	}

	cards, err := trello.GetCards(*env)
	if err != nil {
		log.Printf("failed to get cards: %v", err)
	}

	if env.Github.Event == "issues" {
		issue, err := github.GetIssue(*env)
		if err != nil {
			log.Printf("failed to get issue: %v", err)
		}

		for _, card := range cards {
			issueTitle := fmt.Sprint(issue["title"])
			cardName := fmt.Sprint(card["name"])
			
			if issueTitle == cardName {
				err = trello.DeleteCard(*env, fmt.Sprint(card["id"]))
				if err != nil {
					log.Printf("failed to delete card: %v", err)
				}
			}
		}
	}

	if env.Github.Event == "pull_request" {
		pull, err := github.GetPull(*env)
		if err != nil {
			log.Printf("failed to get pull: %v", err)
		}

		for _, card := range cards {
			pullTitle := fmt.Sprint(pull["title"])
			cardName := fmt.Sprint(card["name"])

			if pullTitle == cardName {
				err = trello.DeleteCard(*env, fmt.Sprint(card["id"]))
				if err != nil {
					log.Printf("failed to delete card: %v", err)
				}
			}
		}
	}
}
