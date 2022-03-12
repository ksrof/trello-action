package models

// Env represents the data structure of the environment variables.
type Env struct {
	TrelloKey    string
	TrelloToken  string
	TrelloIDList string
	GithubToken  string
	GithubUser   string
	GithubRepo   string
	Action       string
}
