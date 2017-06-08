package zabbix

type Request struct {
	JsonRPC string                 `json:"jsonrpc"`
	Action  string                 `json:"action"`
	Params  map[string]interface{} `json:"params"`
	Auth    string                 `json:"auth"`
	ID      uint                   `json:"id"`
}
