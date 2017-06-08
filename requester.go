package zabbix

import (
	"io"
	"net/http"
)

// Requester is an object that can make requests to a URL
type Requester interface {
	Post(url, contentType string, body io.Reader) (*http.Response, error)
}

// HTTPRequester is the net/http implementation of requester
type HTTPRequester struct {
}

// Post sends a post request with the given data/.
func (hr *HTTPRequester) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	return http.Post(url, contentType, body)
}
