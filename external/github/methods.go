package github

import (
	"fmt"
	"net/http"
)

type Requester interface {
	GetIssueByID() *http.Response
	GetPullByID() *http.Response
	GetLabelsFromIssue() *http.Response
}

func (c *config) GetIssueByID() *http.Response {
	fmt.Println(c.token)
	fmt.Println(c.user)
	fmt.Println(c.repo)
	fmt.Println(c.event)
	fmt.Println(c.id)
	return nil
}

func (c *config) GetPullByID() *http.Response {
	return nil
}

func (c *config) GetLabelsFromIssue() *http.Response {
	return nil
}
