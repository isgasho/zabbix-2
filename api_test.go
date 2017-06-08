package zabbix

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

type TestRequester struct {
	Body io.Reader
}

func (tr *TestRequester) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	tr.Body = body

	responseRecorder := httptest.NewRecorder()

	return responseRecorder.Result(), nil
}

func (tr *TestRequester) GetBody() []byte {
	data, err := ioutil.ReadAll(tr.Body)
	if err != nil {
		return nil
	}

	return data
}

func newTestRequester() *TestRequester {
	return new(TestRequester)
}

func TestRequestFormat(t *testing.T) {
	assert := assert.New(t)
	api := NewAPI("test-url")
	testRequester := newTestRequester()
	api.Requester = testRequester

	api.Request("history.get", map[string]interface{}{
		"test": "param",
	})

	body := testRequester.Body
	bodyData, err := ioutil.ReadAll(body)
	if err != nil {
		t.Fail()
	}

	var request Request
	err = json.Unmarshal(bodyData, &request)
	if err != nil {
		t.Fail()
	}

	assert.Equal(RPCVersion, request.JsonRPC, "JSON RPC version incorrect")
	assert.Equal("history.get", request.Action, "Actions do not match")
	assert.Equal("", request.Auth, "Auth should be empty")

	decodedParams := request.Params
	value, exists := decodedParams["test"]
	if passed := assert.Truef(exists, "test not in params"); !passed {
		return
	}

	trueValue, correctType := value.(string)
	if passed := assert.Truef(correctType, "test not a string"); !passed {
		return
	}

	assert.Equal(trueValue, "param")
}
