package internal

import (
	"log"
	"os"
	"strings"

	"github.com/ksrof/gha-trello/models"
	"github.com/spf13/viper"
)

// Env sets the environment variables based on the environment.
// It returns the struct containing the environment variables.
func Env() (models.Env, error) {
	switch strings.ToLower(os.Getenv("ENV")) {
	case "dev":
		viper.SetConfigName("env")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("unable to read environment config: %v", err)
		}

		env := models.Env{
			TrelloKey:    viper.GetString("trello.key"),
			TrelloToken:  viper.GetString("trello.token"),
			TrelloIDList: viper.GetString("trello.list"),
			GithubToken:  viper.GetString("github.token"),
			GithubUser:   viper.GetString("github.user"),
			GithubRepo:   viper.GetString("github.repo"),
			Action:       viper.GetString("action"),
			ID:           viper.GetString("id"),
		}

		return env, nil
	default:
		env := models.Env{
			TrelloKey:    os.Getenv("TRELLO_KEY"),
			TrelloToken:  os.Getenv("TRELLO_TOKEN"),
			TrelloIDList: os.Getenv("TRELLO_ID_LIST"),
			GithubToken:  os.Getenv("GH_TOKEN"),
			GithubUser:   os.Getenv("GH_USER"),
			GithubRepo:   os.Getenv("GH_REPO"),
			Action:       os.Getenv("ACTION"),
			ID:           os.Getenv("ID"),
		}

		return env, nil
	}
}
