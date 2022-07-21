/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ksrof/trello-action/external/github"
	"github.com/ksrof/trello-action/utils"
)

func main() {
	auth, err := github.NewAuth(
		github.WithToken("ghp_F41fR24d9Tvn3YRzC7GdOPAfhgjBzP5MLP7c"),
	)
	if err != nil {
		log.Printf(
			"failed to create a new instance of *github.auth, error: %s\n",
			err.Error(),
		)
		return
	}

	issues := new(github.IssuesResponse)
	issue, err := issues.Get(
		context.TODO(),
		[]utils.Field{
			utils.WithMap(map[string]string{
				"username":    "ksrof",         // TODO: Use os.Getenv
				"repository":  "trello-action", // TODO: Use os.Getenv
				"issue_id":    "30",            // TODO: Use os.Getenv
				"request_url": "https://api.github.com/repos/%s/%s/issues/%s",
				"token":       auth.Basic(),
			}),
		})
	if err != nil {
		log.Printf(
			"failed to use method github.Get(), error: %s\n",
			err.Error(),
		)
		return
	}

	issueLabels, err := issues.GetLabels(
		context.TODO(),
		[]utils.Field{
			utils.WithMap(map[string]string{
				"username":    "ksrof",         // TODO: Use os.Getenv
				"repository":  "trello-action", // TODO: Use os.Getenv
				"issue_id":    "30",            // TODO: Use os.Getenv
				"request_url": "https://api.github.com/repos/%s/%s/issues/%s/labels",
				"token":       auth.Basic(),
			}),
		})
	if err != nil {
		log.Printf(
			"failed to use method github.GetLabels(), error: %s\n",
			err.Error(),
		)
		return
	}

	fmt.Println(issue)
	fmt.Println("---------------------------")
	fmt.Println(issueLabels)
}
