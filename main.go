// Copyright 2022 Kevin Suñer
// SPDX-License-Identifier: Apache-2.0

// This program follows a top-down approach as defined by Procedural Programming.
// It's main objective is to solve the problem at hand, rather than going into
// unnecesary abstractions.

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	TrelloKey     string = os.Getenv("TRELLO_KEY")
	TrelloToken   string = os.Getenv("TRELLO_TOKEN")
	TrelloBoardID string = os.Getenv("TRELLO_BOARD_ID")
	TrelloListID  string = os.Getenv("TRELLO_LIST_ID")

	GithubToken  string = os.Getenv("GH_TOKEN")
	GithubUser   string = os.Getenv("GH_USER")
	GithubRepo   string = os.Getenv("GH_REPO")
	GithubID     string = os.Getenv("GH_ID")
	GithubAction string = os.Getenv("GH_ACTION")
	GithubType   string = os.Getenv("GH_TYPE")
)

func main() {
	var ctx = context.Background()

	err := validateStrings(
		TrelloKey, TrelloToken, TrelloBoardID, TrelloListID,
		GithubToken, GithubUser, GithubRepo, GithubID, GithubAction, GithubType,
	)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	// Change regexp
	err = validateRegexp(TrelloKey, *regexp.MustCompile("[A-Za-z0-9]{32}"))
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	// Change regexp
	err = validateRegexp(TrelloToken, *regexp.MustCompile("[A-Za-z0-9]{64}"))
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	err = validateRegexp(GithubToken, *regexp.MustCompile("[A-Za-z0-9_]{40}"))
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	if GithubAction == "create" {
		err := create(ctx)
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
	}

	if GithubAction == "update" {
		err := update(ctx)
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
	}

	if GithubAction == "delete" {
		err := delete(ctx)
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
	}
}

func create(ctx context.Context) error {
	if GithubType == "issues" {
		issue, err := getIssue(
			ctx,
			"https://api.github.com/repos/%s/%s/issues/%s",
			GithubUser,
			GithubRepo,
			GithubID,
		)
		if err != nil {
			return err
		}

		err = createCard(
			ctx,
			"https://api.trello.com/1/cards",
			issue["name"].(string),
			issue["html_url"].(string),
			TrelloListID,
		)
		if err != nil {
			return err
		}
	}

	if GithubType == "pulls" {
		pull, err := getPull(
			ctx,
			"https://api.github.com/repos/%s/%s/pulls/%s",
			GithubUser,
			GithubRepo,
			GithubID,
		)
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}

		err = createCard(
			ctx,
			"https://api.trello.com/1/cards",
			pull["name"].(string),
			pull["html_url"].(string),
			TrelloListID,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func update(ctx context.Context) error {
	issue, err := getIssue(
		ctx,
		"https://api.github.com/repos/%s/%s/issues/%s",
		GithubUser,
		GithubRepo,
		GithubID,
	)
	if err != nil {
		return err
	}

	labels, err := getLabels(
		ctx,
		"https://api.github.com/repos/%s/%s/issues/%s/labels/",
		GithubUser,
		GithubRepo,
		GithubID,
	)
	if err != nil {
		return err
	}

	lists, err := getLists(
		ctx,
		"https://api.trello.com/v1/boards/%s/lists",
		TrelloBoardID,
	)
	if err != nil {
		return err
	}

	cards, err := getCards(
		ctx,
		"https://api.trello.com/v1/boards/%s/cards",
		TrelloBoardID,
	)
	if err != nil {
		return err
	}

	err = validateNotEmpty(issue, labels, lists, cards)
	if err != nil {
		return err
	}

	var listID string
	for _, label := range labels {
		for _, list := range lists {
			if label["name"] == list["name"] {
				listID = list["id"].(string)
			}

			if label["name"] != list["name"] {
				errStr := "label name does not match list name"
				return errors.New(errStr)
			}
		}
	}

	var cardID string
	for _, card := range cards {
		if card["name"] == issue["name"] {
			cardID = card["id"].(string)
		}

		if card["name"] != issue["name"] {
			errStr := "card name does not match issue name"
			return errors.New(errStr)
		}
	}

	err = updateCard(
		ctx,
		"https://api.trello.com/v1/cards/%s",
		cardID,
		listID,
	)
	if err != nil {
		return err
	}

	return nil
}

// func delete
func delete(ctx context.Context) error {
	issue, err := getIssue(
		ctx,
		"https://api.github.com/repos/%s/%s/issues/%s",
		GithubUser,
		GithubRepo,
		GithubID,
	)
	if err != nil {
		return err
	}

	cards, err := getCards(
		ctx,
		"https://api.trello.com/v1/boards/%s/cards",
		TrelloBoardID,
	)
	if err != nil {
		return err
	}

	err = validateNotEmpty(issue, cards)
	if err != nil {
		return err
	}

	var cardID string
	for _, card := range cards {
		if card["name"] == issue["name"] {
			cardID = card["id"].(string)
		}

		if card["name"] != issue["name"] {
			errStr := "card name does not match issue name"
			return errors.New(errStr)
		}
	}

	err = deleteCard(
		ctx,
		"https://api.trello.com/v1/cards/%s",
		cardID,
	)
	if err != nil {
		return err
	}

	return nil
}

func getIssue(ctx context.Context, url, user, repo, id string) (map[string]any, error) {
	// Use globals
	var token string = os.Getenv("GH_TOKEN")

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	reqURL := fmt.Sprintf(url, user, repo, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	err = validateRegexp(token, *regexp.MustCompile("[A-Za-z0-9_]{40}"))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var issue map[string]any
	err = json.Unmarshal(data, &issue)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

func getPull(ctx context.Context, url, user, repo, id string) (map[string]any, error) {
	// Use globals
	var token string = os.Getenv("GH_TOKEN")

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	reqURL := fmt.Sprintf(url, user, repo, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	err = validateRegexp(token, *regexp.MustCompile("[A-Za-z0-9_]{40}"))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var pull map[string]any
	err = json.Unmarshal(data, &pull)
	if err != nil {
		return nil, err
	}

	return pull, nil
}

func getLabels(ctx context.Context, url, user, repo, id string) ([]map[string]any, error) {
	// Use globals
	var token string = os.Getenv("GH_TOKEN")

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	reqURL := fmt.Sprintf(url, user, repo, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	err = validateRegexp(token, *regexp.MustCompile("[A-Za-z0-9_]{40}"))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var labels []map[string]any
	err = json.Unmarshal(data, &labels)
	if err != nil {
		return nil, err
	}

	return labels, nil
}

func getLists(ctx context.Context, url, id string) ([]map[string]any, error) {
	// Use globals
	var (
		key   string = os.Getenv("TRELLO_KEY")
		token string = os.Getenv("TRELLO_TOKEN")
	)

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	reqURL := fmt.Sprintf(url, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	err = validateRegexp(key, *regexp.MustCompile("[A-Za-z0-9]{32}"))
	if err != nil {
		return nil, err
	}

	err = validateRegexp(token, *regexp.MustCompile("[A-Za-z0-9]{64}"))
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("key", key)
	query.Add("token", token)
	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var lists []map[string]any
	err = json.Unmarshal(data, &lists)
	if err != nil {
		return nil, err
	}

	return lists, nil
}

func getCards(ctx context.Context, url, id string) ([]map[string]any, error) {
	// Use globals
	var (
		key   string = os.Getenv("TRELLO_KEY")
		token string = os.Getenv("TRELLO_TOKEN")
	)

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	reqURL := fmt.Sprintf(url, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	err = validateRegexp(key, *regexp.MustCompile("[A-Za-z0-9]{32}"))
	if err != nil {
		return nil, err
	}

	err = validateRegexp(token, *regexp.MustCompile("[A-Za-z0-9]{64}"))
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("key", key)
	query.Add("token", token)
	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var cards []map[string]any
	err = json.Unmarshal(data, &cards)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func createCard(ctx context.Context, url, name, source, id string) error {
	// Use globals
	var (
		key   string = os.Getenv("TRELLO_KEY")
		token string = os.Getenv("TRELLO_TOKEN")
	)

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, http.NoBody)
	if err != nil {
		return err
	}

	err = validateRegexp(key, *regexp.MustCompile("[A-Za-z0-9]{32}"))
	if err != nil {
		return err
	}

	err = validateRegexp(token, *regexp.MustCompile("[A-Za-z0-9]{64}"))
	if err != nil {
		return err
	}

	query := req.URL.Query()
	query.Add("key", key)
	query.Add("token", token)
	query.Add("name", name)
	query.Add("urlSource", source)
	query.Add("idList", id)
	req.URL.RawQuery = query.Encode()

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func updateCard(ctx context.Context, url, id, listID string) error {
	// Use globals
	var (
		key   string = os.Getenv("TRELLO_KEY")
		token string = os.Getenv("TRELLO_TOKEN")
	)

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	reqURL := fmt.Sprintf(url, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, reqURL, http.NoBody)
	if err != nil {
		return err
	}

	err = validateRegexp(key, *regexp.MustCompile("[A-Za-z0-9]{32}"))
	if err != nil {
		return err
	}

	err = validateRegexp(token, *regexp.MustCompile("[A-Za-z0-9]{64}"))
	if err != nil {
		return err
	}

	query := req.URL.Query()
	query.Add("key", key)
	query.Add("token", token)
	query.Add("idList", listID)
	req.URL.RawQuery = query.Encode()

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func deleteCard(ctx context.Context, url, id string) error {
	// Use globals
	var (
		key   string = os.Getenv("TRELLO_KEY")
		token string = os.Getenv("TRELLO_TOKEN")
	)

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	reqURL := fmt.Sprintf(url, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, reqURL, http.NoBody)
	if err != nil {
		return err
	}

	err = validateRegexp(key, *regexp.MustCompile("[A-Za-z0-9]{32}"))
	if err != nil {
		return err
	}

	err = validateRegexp(token, *regexp.MustCompile("[A-Za-z0-9]{64}"))
	if err != nil {
		return err
	}

	query := req.URL.Query()
	query.Add("key", key)
	query.Add("token", token)
	req.URL.RawQuery = query.Encode()

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func validateStrings(str ...string) error {
	for _, s := range str {
		if len(strings.TrimSpace(s)) == 0 {
			errStr := "value of string cant be of zero length"
			return errors.New(errStr)
		}
	}

	return nil
}

func validateRegexp(str string, pattern regexp.Regexp) error {
	ok := pattern.MatchString(str)
	if !ok {
		errStr := "value of string does not match regexp"
		return errors.New(errStr)
	}

	return nil
}

func validateNotEmpty(values ...any) error {
	errStr := "value length cant be zero"

	for _, value := range values {
		switch value := value.(type) {
		case map[string]any:
			if len(value) == 0 {
				return errors.New(errStr)
			}
		case []map[string]any:
			if len(value) == 0 {
				return errors.New(errStr)
			}
		}
	}

	return nil
}
