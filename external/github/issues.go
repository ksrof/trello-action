package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ksrof/trello-action/utils"
)

//go:generate mockgen -destination=mock/issues.go -package=mock . Issues
type Issues interface {
	Get(ctx context.Context, opts []utils.Field) (*IssuesResponse, error)
	GetLabels(ctx context.Context, opts []utils.Field) (*IssuesResponse, error)
}

type IssuesResponse struct {
	Response interface{}
	Status   string
	Code     int
	Error    string
}

func (r *IssuesResponse) Get(ctx context.Context, opts []utils.Field) (*IssuesResponse, error) {
	fields, err := utils.NewFieldsMapper(opts...)
	if err != nil {
		return &IssuesResponse{
			Status: http.StatusText(http.StatusBadRequest),
			Code:   http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	field := fields.(map[string]string)
	reqURL := fmt.Sprintf(field["request_url"], field["username"], field["repository"], field["issue_id"])

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return &IssuesResponse{
			Status: http.StatusText(http.StatusInternalServerError),
			Code:   http.StatusInternalServerError,
			Error:  err.Error(),
		}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", field["token"]))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &IssuesResponse{
			Status: http.StatusText(http.StatusInternalServerError),
			Code:   http.StatusInternalServerError,
			Error:  err.Error(),
		}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &IssuesResponse{
			Status: http.StatusText(http.StatusInternalServerError),
			Code:   http.StatusInternalServerError,
			Error:  err.Error(),
		}, err
	}

	var response map[string]any
	err = json.Unmarshal(data, &response)
	if err != nil {
		return &IssuesResponse{
			Status: http.StatusText(http.StatusInternalServerError),
			Code:   http.StatusInternalServerError,
			Error:  err.Error(),
		}, err
	}

	return &IssuesResponse{
		Response: response,
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
	}, nil
}

func (r *IssuesResponse) GetLabels(ctx context.Context, opts []utils.Field) (*IssuesResponse, error) {
	fields, err := utils.NewFieldsMapper(opts...)
	if err != nil {
		return &IssuesResponse{
			Status: http.StatusText(http.StatusBadRequest),
			Code:   http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	field := fields.(map[string]string)
	reqURL := fmt.Sprintf(field["request_url"], field["username"], field["repository"], field["issue_id"])

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return &IssuesResponse{
			Status: http.StatusText(http.StatusInternalServerError),
			Code:   http.StatusInternalServerError,
			Error:  err.Error(),
		}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", field["token"]))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &IssuesResponse{
			Status: http.StatusText(http.StatusInternalServerError),
			Code:   http.StatusInternalServerError,
			Error:  err.Error(),
		}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &IssuesResponse{
			Status: http.StatusText(http.StatusInternalServerError),
			Code:   http.StatusInternalServerError,
			Error:  err.Error(),
		}, err
	}

	var response []map[string]any
	err = json.Unmarshal(data, &response)
	if err != nil {
		return &IssuesResponse{
			Status: http.StatusText(http.StatusInternalServerError),
			Code:   http.StatusInternalServerError,
			Error:  err.Error(),
		}, err
	}

	return &IssuesResponse{
		Response: response,
		Status:   http.StatusText(http.StatusOK),
		Code:     http.StatusOK,
	}, nil
}
