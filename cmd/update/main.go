package main

import (
	"fmt"
	"log"

	"github.com/ksrof/trello-action/internal/configs"
	"github.com/ksrof/trello-action/internal/github"
	"github.com/ksrof/trello-action/internal/trello"
)

var (
	cardID    string
	labelName string
	listID    string
)

func main() {
	env, err := configs.Environment()
	if err != nil {
		log.Fatalf("failed to load environment: %v", err)
	}

	cards, err := trello.GetCards(*env)
	if err != nil {
		log.Printf("failed to get cards: %v", err)
	}

	lists, err := trello.GetLists(*env)
	if err != nil {
		log.Printf("failed to get lists: %v", err)
	}

	labels, err := github.GetLabels(*env)
	if err != nil {
		log.Printf("failed to get labels: %v", err)
	}

	// Loop through the labels of the Issue or Pull Request.
	for _, label := range labels {
		labelName = fmt.Sprint(label["name"])
	}

	// Loop through all the lists of the board.
	// Check if label name and list name match.
	// If they match return the ID of the list.
	for _, list := range lists {
		listName := fmt.Sprint(list["name"])

		if labelName == listName {
			listID = fmt.Sprint(list["id"])
		}
	}

	if env.Github.Event == "issues" {
		issue, err := github.GetIssue(*env)
		if err != nil {
			log.Printf("failed to get issue: %v", err)
		}

		issueTitle := fmt.Sprint(issue["title"])

		// Loop through all the cards of the board.
		// Check if issue title and card name match.
		// If they match return the ID of the card.
		for _, card := range cards {
			cardName := fmt.Sprint(card["name"])

			if issueTitle == cardName {
				cardID = fmt.Sprint(card["id"])
			}
		}

		err = trello.UpdateCard(*env, cardID, listID)
		if err != nil {
			log.Printf("failed to update card")
		}
	}

	if env.Github.Event == "pull_request" {
		pull, err := github.GetPull(*env)
		if err != nil {
			log.Printf("failed to get pull: %v", err)
		}

		pullTitle := fmt.Sprint(pull["title"])

		// Loop through all the cards of the board.
		// Check if pull title and card name match.
		// If they match return the ID of the card.
		for _, card := range cards {
			cardName := fmt.Sprint(card["name"])

			if pullTitle == cardName {
				cardID = fmt.Sprint(card["id"])
			}
		}

		err = trello.UpdateCard(*env, cardID, listID)
		if err != nil {
			log.Printf("failed to update card")
		}
	}
}
