package main

import (
	"context"
	"log"

	"github.com/ksrof/trello-action/external/github"
)

func main() {
	ctx := context.Background()

	ghClient, err := github.NewClient()
	if err != nil {
		log.Fatalln(err)
	}

	ghService := github.NewService(ghClient)

	_, err = ghService.GetIssueByID(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	// _, _ = ghService.GetPullByID(ctx)
	// _, _ = ghService.GetLabelsFromIssue(ctx)
}
