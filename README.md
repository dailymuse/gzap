# Go Muse Logger

[![GoDoc](https://godoc.org/github.com/dailymuse/gml?status.svg)](https://godoc.org/github.com/dailymuse/gml) [![CircleCI](https://circleci.com/gh/dailymuse/gml.svg?style=svg)](https://circleci.com/gh/dailymuse/gml) [![codecov](https://codecov.io/gh/dailymuse/gml/branch/develop/graph/badge.svg)](https://codecov.io/gh/dailymuse/gml)




```go
package main

import (
	"time"

	"github.com/dailymuse/gml"
)

func main() {
    if err := gml.New(&gml.Config{
        AppName: "app-name",
        IsProdEnv: true,
        IsStagingEnv: false,
        IsTestEnv: false,
        GraylogAddress: "127.0.0.1",
        GraylogPort: 1234,
        GraylogVersion: "1.1",
        Hostname: "myhostname",
        UseTLS: true,
        InsecureSkipVerify: true,
        LogEnvName: "prod",
        GraylogConnectionTimeout: time.Second * 3,
    }); err != nil {
        panic(err)
    }

    gml.Logger.Info(..)
}
```