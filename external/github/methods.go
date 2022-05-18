package github

type Requester interface {
	GetIssueByID() (interface{}, error)
	GetPullByID() (interface{}, error)
	GetLabelsFromIssue() (interface{}, error)
}

// TODO: Decode http.Response in a separated function.

func (c *config) GetIssueByID() (interface{}, error) {
	return nil, nil
}

func (c *config) GetPullByID() (interface{}, error) {
	return nil, nil
}

func (c *config) GetLabelsFromIssue() (interface{}, error) {
	return nil, nil
}
