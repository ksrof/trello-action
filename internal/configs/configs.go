package configs

import (
	"fmt"
	"os"

	"github.com/ksrof/gha-trello/internal/models"
	"github.com/spf13/viper"
)

// Environment saves the values from the environment variables.
// It returns a struct containing the values of the environment variables.
func Environment() (*models.Env, error) {
	// When ENV equals "dev" load YAML file.
	if os.Getenv("ENV") == "dev" {
		viper.SetConfigName("env")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")

		err := viper.ReadInConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to load yaml file: %v", err)
		}

		return &models.Env{
			Trello: &models.Trello{
				Key:    viper.GetString("trello.key"),
				Token:  viper.GetString("trello.token"),
				IDList: viper.GetString("trello.idList"),
			},
			Github: &models.Github{
				Token: viper.GetString("github.token"),
				User:  viper.GetString("github.user"),
				Repo:  viper.GetString("github.repo"),
				Event: viper.GetString("github.event"),
				ID:    viper.GetString("github.id"),
			},
		}, nil
	}

	// Otherwise just get the values directly.
	return &models.Env{
		Trello: &models.Trello{
			Key:    os.Getenv("TRELLO_KEY"),
			Token:  os.Getenv("TRELLO_TOKEN"),
			IDList: os.Getenv("TRELLO_ID_LIST"),
		},
		Github: &models.Github{
			Token: os.Getenv("GH_TOKEN"),
			User:  os.Getenv("GH_USER"),
			Repo:  os.Getenv("GH_REPO"),
			Event: os.Getenv("GH_EVENT"),
			ID:    os.Getenv("GH_ID"),
		},
	}, nil
}
