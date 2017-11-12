# gzap - Graylog Integrated Zap Logger

[![GoDoc](https://godoc.org/github.com/dailymuse/gzap?status.svg)](https://godoc.org/github.com/dailymuse/gzap) [![CircleCI](https://circleci.com/gh/dailymuse/gzap.svg?style=svg)](https://circleci.com/gh/dailymuse/gzap) [![codecov](https://codecov.io/gh/dailymuse/gzap/branch/master/graph/badge.svg)](https://codecov.io/gh/dailymuse/gzap)


### Getting Stated

To use gzap, simply import:

```go
import "github.com/dailymuse/gzap"
```

### Example Usage

```go
package main

import (
    "time"

    "github.com/dailymuse/gzap"
)

func main() {
    // Instantiate a global logger.
    if err := gzap.New(&gzap.Config{
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

    // Example Info log.
    gzap.Logger.Info("this is an example Info log",
        gzap.String("process name", "some-fake-name"),
        gzap.Int64("expectedDocs", int64(255)),
        gzap.Int64("docsUploaded", int64(100)),
    )

    // Example Error log.
    gzap.Logger.Error("this is an example Error log",
        gzap.Error(errors.New("example error")),
        gzap.String("index name", "my-full-index-name"),
        gzap.Float64("time elapsed", float64(1002)),
    )

    // Example Debug log.
    gzap.Logger.Error("this is an example Debug log",
        gzap.String("variable", "some-variable-here"),
    )
}
```