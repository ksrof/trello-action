package internal

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// NewReq builds a new http request with a given method, url and payload.
// The url gets parsed before being used in the new request.
// It returns the newly built request ready to be sent from the client to the server.
func NewReq(method, URL string, payload io.Reader) (req *http.Request, err error) {
	reqURL, err := url.Parse(URL)
	if err != nil {
		return req, fmt.Errorf("failed to parse url: %v", err)
	}

	req, err = http.NewRequest(method, reqURL.String(), payload)
	if err != nil {
		return req, fmt.Errorf("failed to build a new http request: %v", err)
	}

	return req, nil
}
