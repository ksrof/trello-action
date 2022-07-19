package main

import (
	"fmt"
	"log"

	"github.com/ksrof/trello-action/external/github"
)

func main() {
	auth, err := github.NewAuth(
		github.WithUser("johndoe"),
		github.WithRepo("floppyunicorn"),
		github.WithToken("token"),
	)
	if err != nil {
		log.Printf(
			"failed to create a new instance of *github.auth, error: %s",
			err.Error(),
		)
		return
	}

	if auth.Basic() == "" {
		log.Println("failed to return token")
		return
	}

	fmt.Println(auth.Basic())
}
