package zabbix

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

type authTestRequester struct {
	Body io.Reader
}

func (atr *authTestRequester) Post(url, action string, data io.Reader) (*http.Response, error) {
	atr.Body = data

	responseJson := []byte(`{"jsonrpc": "2.0", "result": "abc123"}`)

	recorder := httptest.NewRecorder()
	recorder.Body = bytes.NewBuffer(responseJson)

	return recorder.Result(), nil
}

func TestAuth(t *testing.T) {
	assert := assert.New(t)
	tr := new(authTestRequester)

	api := NewAPI("test")
	api.Requester = tr

	err := api.Auth("user", "password")

	assert.Nil(err, "Error should be nil, instead is: ", err)
	assert.Equal("abc123", api.AuthToken)
}
