package models

import "time"

// Pull represents the data structure of the api.github.com/repos/user/repo/pulls/id response body.
type Pull struct {
	URL                 string        `json:"url"`
	ID                  int           `json:"id"`
	NodeID              string        `json:"node_id"`
	HTMLURL             string        `json:"html_url"`
	DiffURL             string        `json:"diff_url"`
	PatchURL            string        `json:"patch_url"`
	IssueURL            string        `json:"issue_url"`
	Number              int           `json:"number"`
	State               string        `json:"state"`
	Locked              bool          `json:"locked"`
	Title               string        `json:"title"`
	Body                interface{}   `json:"body"`
	CreatedAt           time.Time     `json:"created_at"`
	UpdatedAt           time.Time     `json:"updated_at"`
	ClosedAt            interface{}   `json:"closed_at"`
	MergedAt            interface{}   `json:"merged_at"`
	MergeCommitSha      string        `json:"merge_commit_sha"`
	Assignee            interface{}   `json:"assignee"`
	Assignees           []interface{} `json:"assignees"`
	RequestedReviewers  []interface{} `json:"requested_reviewers"`
	RequestedTeams      []interface{} `json:"requested_teams"`
	Labels              []interface{} `json:"labels"`
	Milestone           interface{}   `json:"milestone"`
	Draft               bool          `json:"draft"`
	CommitsURL          string        `json:"commits_url"`
	ReviewCommentsURL   string        `json:"review_comments_url"`
	ReviewCommentURL    string        `json:"review_comment_url"`
	CommentsURL         string        `json:"comments_url"`
	StatusesURL         string        `json:"statuses_url"`
	AuthorAssociation   string        `json:"author_association"`
	AutoMerge           interface{}   `json:"auto_merge"`
	ActiveLockReason    interface{}   `json:"active_lock_reason"`
	Merged              bool          `json:"merged"`
	Mergeable           bool          `json:"mergeable"`
	Rebaseable          bool          `json:"rebaseable"`
	MergeableState      string        `json:"mergeable_state"`
	MergedBy            interface{}   `json:"merged_by"`
	Comments            int           `json:"comments"`
	ReviewComments      int           `json:"review_comments"`
	MaintainerCanModify bool          `json:"maintainer_can_modify"`
	Commits             int           `json:"commits"`
	Additions           int           `json:"additions"`
	Deletions           int           `json:"deletions"`
	ChangedFiles        int           `json:"changed_files"`
}
