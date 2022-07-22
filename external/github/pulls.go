package github

import (
	"context"
	"fmt"

	"github.com/ksrof/trello-action/utils"
)

type Pulls interface {
	Get(ctx context.Context, opts []utils.Field) (*PullsResponse, error)
}

type PullsResponse struct {
	Response interface{}
	Status   string
	Code     int
	Error    string
}

func (r *PullsResponse) Get(ctx context.Context, opts []utils.Field) (*PullsResponse, error) {
	fmt.Println("--- to be implemented ---")
	return nil, nil
}
