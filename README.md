# Go Muse Logger

[![GoDoc](https://godoc.org/github.com/dailymuse/gml?status.svg)](https://godoc.org/github.com/dailymuse/gml) [![CircleCI](https://circleci.com/gh/dailymuse/gml.svg?style=svg)](https://circleci.com/gh/dailymuse/gml)


```go
package main

import (
	"time"

	"github.com/dailymuse/gml"
)

func main() {
    err := gml.New(&gml.Config{
        GetAppName: func() string {

        },
        GetIsProdEnv: func() bool {

        },
        GetIsStagingEnv: func() bool {

        },
        GetIsTestEnv: func() bool {

        },
        GetGraylogAddress: func() string {

        },
        GetGraylogPort: func() uint {

        },
        GetGraylogVersion: func() string {

        },
        GetHostname: func() string {

        },
        GetUseTLS: func() bool {

        },
        GetInsecureSkipVerify: func() bool {

        },
        GetLogEnvName: func() string {

        },
        GetGraylogConnectionTimeout: func() time.Duration {
            
        },
    })

    gomuselog.Logger.Info()
}
```