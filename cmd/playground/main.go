/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"context"
	"fmt"

	"github.com/ksrof/trello-action/external/github"
)

func main() {
	token, err := github.Basic("ghp_F41fR24d9Tvn3YRzC7GdOPAfhgjBzP5MLP7c") // TODO: Use os.Getenv
	if err != nil {
		return
	}

	issues := new(github.IssuesResponse)
	issue, err := issues.Get(
		context.TODO(),
		map[string]string{
			"username":    "ksrof",         // TODO: Use os.Getenv
			"repository":  "trello-action", // TODO: Use os.Getenv
			"issue_id":    "30",            // TODO: Use os.Getenv
			"request_url": "https://api.github.com/repos/%s/%s/issues/%s",
			"token":       token,
		},
	)
	if err != nil {
		return
	}

	labels, err := issues.GetLabels(
		context.TODO(),
		map[string]string{
			"username":    "ksrof",         // TODO: Use os.Getenv
			"repository":  "trello-action", // TODO: Use os.Getenv
			"issue_id":    "30",            // TODO: Use os.Getenv
			"request_url": "https://api.github.com/repos/%s/%s/issues/%s/labels",
			"token":       token,
		},
	)
	if err != nil {
		return
	}

	pulls := new(github.PullsResponse)
	pull, err := pulls.Get(
		context.TODO(),
		map[string]string{
			"username":    "ksrof",         // TODO: Use os.Getenv
			"repository":  "trello-action", // TODO: Use os.Getenv
			"pull_id":     "30",            // TODO: Use os.Getenv
			"request_url": "https://api.github.com/repos/%s/%s/pulls/%s",
			"token":       token,
		},
	)
	if err != nil {
		return
	}

	fmt.Println(issue)
	fmt.Println("---------------------------")
	fmt.Println(labels)
	fmt.Println("---------------------------")
	fmt.Println(pull)
}
