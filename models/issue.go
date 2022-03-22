package models

import "time"

// Issue represents the data structure of the api.github.com/repos/user/repo/issues/id response body.
type Issue struct {
	URL                   string        `json:"url"`
	RepositoryURL         string        `json:"repository_url"`
	LabelsURL             string        `json:"labels_url"`
	CommentsURL           string        `json:"comments_url"`
	EventsURL             string        `json:"events_url"`
	HTMLURL               string        `json:"html_url"`
	ID                    int           `json:"id"`
	NodeID                string        `json:"node_id"`
	Number                int           `json:"number"`
	Title                 string        `json:"title"`
	Labels                []interface{} `json:"labels"`
	State                 string        `json:"state"`
	Locked                bool          `json:"locked"`
	Assignee              interface{}   `json:"assignee"`
	Assignees             []interface{} `json:"assignees"`
	Milestone             interface{}   `json:"milestone"`
	Comments              int           `json:"comments"`
	CreatedAt             time.Time     `json:"created_at"`
	UpdatedAt             time.Time     `json:"updated_at"`
	ClosedAt              interface{}   `json:"closed_at"`
	AuthorAssociation     string        `json:"author_association"`
	ActiveLockReason      interface{}   `json:"active_lock_reason"`
	Body                  interface{}   `json:"body"`
	ClosedBy              interface{}   `json:"closed_by"`
	TimelineURL           string        `json:"timeline_url"`
	PerformedViaGithubApp interface{}   `json:"performed_via_github_app"`
}
