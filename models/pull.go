package models

import "time"

// Pull represents the data structure of the api.github.com/repos/user/repo/pulls response body.
type Pull []struct {
	URL      string `json:"url"`
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	HTMLURL  string `json:"html_url"`
	DiffURL  string `json:"diff_url"`
	PatchURL string `json:"patch_url"`
	IssueURL string `json:"issue_url"`
	Number   int    `json:"number"`
	State    string `json:"state"`
	Locked   bool   `json:"locked"`
	Title    string `json:"title"`
	User     struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"user"`
	Body               interface{}   `json:"body"`
	CreatedAt          time.Time     `json:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at"`
	ClosedAt           interface{}   `json:"closed_at"`
	MergedAt           interface{}   `json:"merged_at"`
	MergeCommitSha     string        `json:"merge_commit_sha"`
	Assignee           interface{}   `json:"assignee"`
	Assignees          []interface{} `json:"assignees"`
	RequestedReviewers []interface{} `json:"requested_reviewers"`
	RequestedTeams     []interface{} `json:"requested_teams"`
	Labels             []interface{} `json:"labels"`
	Milestone          interface{}   `json:"milestone"`
	Draft              bool          `json:"draft"`
	CommitsURL         string        `json:"commits_url"`
	ReviewCommentsURL  string        `json:"review_comments_url"`
	ReviewCommentURL   string        `json:"review_comment_url"`
	CommentsURL        string        `json:"comments_url"`
	StatusesURL        string        `json:"statuses_url"`
}
