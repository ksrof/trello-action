package main

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/ksrof/trello-action"
)

func main() {
	config, err := config.NewConfig(
		config.WithMethods(
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
		),
		config.WithPaths(
			"https://api.trello.com/v1",
		),
		config.WithParams(
			map[string]string{
				"key":     "the-key",
				"token":   "the-token",
				"idBoard": "the-id-board",
				"idList":  "the-id-list",
			},
		),
	)
	if err != nil {
		log.Fatalf("failed to initialize a new configuration object: %v\n", err)
	}

	fmt.Println(config)
}
