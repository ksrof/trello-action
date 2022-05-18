package main

import (
	"log"
	"os"

	"github.com/ksrof/trello-action/external/github"
)

func main() {
	github, err := github.New(
		github.WithToken(os.Getenv("GH_TOKEN")),
		github.WithUser(os.Getenv("GH_USER")),
		github.WithRepo(os.Getenv("GH_REPO")),
		github.WithEvent(os.Getenv("GH_EVENT")),
		github.WithID(os.Getenv("GITHUB_ID")),
	)
	if err != nil {
		log.Fatalln(err)
	}

	_, _ = github.GetIssueByID()
	_, _ = github.GetPullByID()
	_, _ = github.GetLabelsFromIssue()
}
