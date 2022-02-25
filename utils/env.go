package utils

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var env string = "dev"

// Config represents the data structure
// that holds the environment variables
type Config struct {
	TrelloKey    string
	TrelloToken  string
	TrelloIDList string
	GithubToken  string
	GithubUser   string
	GithubRepo   string
	Action       string
}

// SetEnv sets the environment variables
// using the Config struct and returns it
func SetEnv() (Config, error) {
	switch strings.ToLower(env) {
	case "dev":
		// Load environment config from YAML file
		viper.SetConfigName("env")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("unable to read environment config: %v", err)
		}

		config := Config{
			TrelloKey:    viper.GetString("trello.key"),
			TrelloToken:  viper.GetString("trello.token"),
			TrelloIDList: viper.GetString("trello.list"),
			GithubToken:  viper.GetString("github.token"),
			GithubUser:   viper.GetString("github.user"),
			GithubRepo:   viper.GetString("github.repo"),
			Action:       viper.GetString("action"),
		}

		return config, nil
	case "prod":
		config := Config{
			TrelloKey:    os.Getenv("TRELLO_KEY"),
			TrelloToken:  os.Getenv("TRELLO_TOKEN"),
			TrelloIDList: os.Getenv("TRELLO_ID_LIST"),
			GithubToken:  os.Getenv("GH_TOKEN"),
			GithubUser:   os.Getenv("GH_USER"),
			GithubRepo:   os.Getenv("GH_REPO"),
			Action:       os.Getenv("ACTION"),
		}

		return config, nil
	}

	return Config{}, nil
}
