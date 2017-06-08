[![Build Status](https://travis-ci.org/Gman98ish/zabbix.svg?branch=master)](https://travis-ci.org/Gman98ish/zabbix)

# Zabbix
This library acts as a wrapper around the zabbix API

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