package internal

import (
	"fmt"
	"net/http"

	"github.com/ksrof/gha-trello/models"
)

// Params adds different values to different keys.
// It returns the given request with the query parameters set.
func Params(params models.Params, req *http.Request) (*http.Request, error) {
	query := req.URL.Query()
	query.Add("idList", params.IDList)
	query.Add("key", params.Key)
	query.Add("token", params.Token)
	query.Add("name", fmt.Sprintf("%s %s", params.Title, params.Number))
	query.Add("urlSource", params.URL)
	req.URL.RawQuery = query.Encode()

	return req, nil
}
