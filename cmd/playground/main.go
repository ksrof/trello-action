package main

import (
	"log"

	"github.com/ksrof/trello-action/external/github"
)

func main() {
	client, err := github.New(
		github.WithToken("dGhpc2lzYXN1cGVyc2VjcmV0dG9rZW4="),
		github.WithUser("ksrof"),
		github.WithRepo("trello-action"),
		github.WithEvent("issue"),
		github.WithID("13"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	_ = client.GetIssueByID()
}
