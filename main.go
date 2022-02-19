package main

// TODO's
// Perform a GET requests to the github api endpoint and
// return the data from a specific user.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

const baseURL string = "https://api.github.com"

// User represents the data structure of the
// api.github.com/user response body.
type User struct {
	Login                   string      `json:"login"`
	ID                      int         `json:"id"`
	NodeID                  string      `json:"node_id"`
	AvatarURL               string      `json:"avatar_url"`
	GravatarID              string      `json:"gravatar_id"`
	URL                     string      `json:"url"`
	HTMLURL                 string      `json:"html_url"`
	FollowersURL            string      `json:"followers_url"`
	FollowingURL            string      `json:"following_url"`
	GistsURL                string      `json:"gists_url"`
	StarredURL              string      `json:"starred_url"`
	SubscriptionsURL        string      `json:"subscriptions_url"`
	OrganizationsURL        string      `json:"organizations_url"`
	ReposURL                string      `json:"repos_url"`
	EventsURL               string      `json:"events_url"`
	ReceivedEventsURL       string      `json:"received_events_url"`
	Type                    string      `json:"type"`
	SiteAdmin               bool        `json:"site_admin"`
	Name                    string      `json:"name"`
	Company                 string      `json:"company"`
	Blog                    string      `json:"blog"`
	Location                string      `json:"location"`
	Email                   interface{} `json:"email"`
	Hireable                bool        `json:"hireable"`
	Bio                     string      `json:"bio"`
	TwitterUsername         string      `json:"twitter_username"`
	PublicRepos             int         `json:"public_repos"`
	PublicGists             int         `json:"public_gists"`
	Followers               int         `json:"followers"`
	Following               int         `json:"following"`
	CreatedAt               time.Time   `json:"created_at"`
	UpdatedAt               time.Time   `json:"updated_at"`
	PrivateGists            int         `json:"private_gists"`
	TotalPrivateRepos       int         `json:"total_private_repos"`
	OwnedPrivateRepos       int         `json:"owned_private_repos"`
	DiskUsage               int         `json:"disk_usage"`
	Collaborators           int         `json:"collaborators"`
	TwoFactorAuthentication bool        `json:"two_factor_authentication"`
	Plan                    struct {
		Name          string `json:"name"`
		Space         int    `json:"space"`
		Collaborators int    `json:"collaborators"`
		PrivateRepos  int    `json:"private_repos"`
	} `json:"plan"`
}

// jsonify takes a response body converts it to json
// using the User struct and returns it as a string.
func jsonify(body []byte) (string, error) {
	var user User

	err := json.Unmarshal(body, &user)
	if err != nil {
		return "", fmt.Errorf("unable to unmarshal body response: %v", err)
	}

	json, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		return "", fmt.Errorf("unable to marshal user struct: %v", err)
	}

	return string(json), nil
}

// getUser performs a GET request to api.github.com/user
// with an Authorization token and prints the response formatted.
func getUser(token string) error {
	reqURL := fmt.Sprintf("%s/user", baseURL)
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return fmt.Errorf("unable to perform new request: %v", err)
	}

	// Set Authorization Header with Github Personal Access Token
	authToken := fmt.Sprintf("token %s", token)
	req.Header.Set("Authorization", authToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("unable to perform GET request: %v", err)
	}

	// Check the Error return value
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response body: %v", err)
	}

	// Format the response body
	respJSON, err := jsonify(body)
	if err != nil {
		return fmt.Errorf("unable to parse respons: %v", err)
	}

	// Output the formatted response
	fmt.Println(respJSON)

	return nil
}

func main() {
	// Load environment configuration from YAML file
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("unable to read environment configuration: %v", err)
	}

	// Save environment configuration values as variables
	token := viper.GetString("gh_config.token")

	err = getUser(token)
	if err != nil {
		log.Fatalf("unable to get user: %v", err)
	}
}
