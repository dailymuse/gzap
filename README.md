# Go Muse Logger

[![GoDoc](https://godoc.org/github.com/dailymuse/gml?status.svg)](https://godoc.org/github.com/dailymuse/gml) [![CircleCI](https://circleci.com/gh/dailymuse/gml.svg?style=svg)](https://circleci.com/gh/dailymuse/gml) [![codecov](https://codecov.io/gh/dailymuse/gml/branch/master/graph/badge.svg)](https://codecov.io/gh/dailymuse/gml)


### Getting Stated

To use gml, simply import:

```go
import "github.com/dailymuse/gml"
```

### Example Usage

```go
package main

import (
	"time"

	"github.com/dailymuse/gml"
)

func main() {
    // Instantiate a global logger.
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

    // Example Info log.
    gml.Logger.Info("this is an example Info log",
        gml.Error(errors.New("example error")),
        gml.String("process name", "some-fake-name"),
    )

    // Example Error log.
    gml.Logger.Error("this is an example Error log",
        gml.Int64("docsUploaded", int64(100)),
        gml.Int64("expectedDocs", int64(255)),
        gml.String("index name", "my-full-index-name"),
        gml.Float64("time elapsed", float64(1002)),
    )
}
```