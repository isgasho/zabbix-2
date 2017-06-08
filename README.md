# Zabbix [![Build Status](https://travis-ci.org/Gman98ish/zabbix.svg?branch=master)](https://travis-ci.org/Gman98ish/zabbix)
This library acts as a wrapper around the zabbix API

# Notice
This library is still under development, and is likely to undergo a lot of change until it reaches a more
complete state

## Basic Usage
```go
package main

import (
    "github.com/gman98ish/zabbix"
)

func main() {
    api := zabbix.NewAPI()

    res, err := api.Request("history.get", map[string]interface{}{
        "itemids": "123",
    })

    // do stuff with response
}

```

## Contributing

All contributions are welcome, just open an issue and/or submit an MR, and I'll get back to you ASAP