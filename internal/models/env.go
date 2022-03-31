package models

// Trello struct handles all data needed to work with the Trello API.
type Trello struct {
	Key     string
	Token   string
	IDBoard string
	IDList  string
}

// Github struct handles all data needed to work with the Github API.
type Github struct {
	Token string
	User  string
	Repo  string
	Event string
	ID    string
}

// Env struct handles all data needed for the application to work.
type Env struct {
	Trello *Trello
	Github *Github
}
