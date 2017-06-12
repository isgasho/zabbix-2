package zabbix

type Request struct {
	JsonRPC string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
	Auth    string                 `json:"auth,omitempty"`
	ID      uint                   `json:"id"`
}
