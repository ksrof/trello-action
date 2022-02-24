package utils

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var env string = "dev"

type Config struct {
	TrelloKey    string
	TrelloToken  string
	TrelloIDList string
	GithubToken  string
	GithubUser   string
	GithubRepo   string
	Action       string
}

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
			GithubToken:  os.Getenv("GITHUB_TOKEN"),
			GithubUser:   os.Getenv("GITHUB_USER"),
			GithubRepo:   os.Getenv("GITHUB_REPO"),
			Action:       os.Getenv("ACTION"),
		}

		return config, nil
	}

	return Config{}, nil
}
