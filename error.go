package zabbix

import (
	"fmt"
)

// Error is a JsonRPC error object as defined in the
// standard. It implements the error interface.
type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
