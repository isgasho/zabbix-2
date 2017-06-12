package zabbix

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

const (
	// RPCVersion is the RPC version to use
	RPCVersion = "2.0"
)

// API makes JSON RPC requests to a zabbix API pointed at
// by URL.
type API struct {
	AuthToken string
	URL       string

	Requester Requester
	Logger    Logger

	idCount uint
}

// NewAPI creates a default API object and sets the requester
// to a http implementation
func NewAPI(url string) *API {
	api := &API{
		URL: url,
	}

	api.Requester = new(HTTPRequester)
	api.Logger = new(StdOutLogger)

	return api
}

// Request makes a zabbix API call and returns its response.
// The error isn't indicative of an API error, instead representing
// and error during processing the parameters, for example, `error
// parsing JSON`. For API specific errors, check the error object
// as a part of the response
func (a *API) Request(method string, params map[string]interface{}) (*Response, error) {
	request := Request{
		Auth:    a.AuthToken,
		ID:      1,
		JsonRPC: RPCVersion,
		Method:  method,
		Params:  params,
	}

	requestData, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	requestReader := bytes.NewReader(requestData)
	res, err := a.Requester.Post(a.URL, "application/json", requestReader)
	if err != nil {
		return nil, err
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(resData, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetID returns a new, sequentially generated zabbix
// ID
func (a *API) GetID() uint {
	a.idCount++
	return a.idCount
}

// Auth makes an authentication request to the zabbix
// API with the supplied credentials. The token
// sent back will be used for subsequent requests.
func (a *API) Auth(username, password string) error {
	res, err := a.Request("user.login", map[string]interface{}{
		"user":     username,
		"password": password,
	})

	if err != nil {
		return err
	}

	if res.IsError() {
		return &res.Error
	}

	var tok string
	err = json.Unmarshal(res.Result, &tok)
	if err != nil {
		return err
	}

	a.AuthToken = tok

	return nil
}
