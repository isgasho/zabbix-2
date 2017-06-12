package zabbix

import (
	"encoding/json"
)

// Response is a JsonRPC response as defined in
// the JsonRPC spec.
type Response struct {
	JsonRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   Error
}

// IsError returns if the response failed
func (r *Response) IsError() bool {
	if r.Error.Error() != "0: " {
		return true
	}

	return false
}

// Success returns if the response was successful
// wraps aournd IsError
func (r *Response) Success() bool {
	return !r.IsError()
}
