# gzap - Graylog Integrated Zap Logger

[![GoDoc](https://godoc.org/github.com/dailymuse/gzap?status.svg)](https://godoc.org/github.com/dailymuse/gzap) [![CircleCI](https://circleci.com/gh/dailymuse/gzap.svg?style=svg)](https://circleci.com/gh/dailymuse/gzap) [![codecov](https://codecov.io/gh/dailymuse/gzap/branch/master/graph/badge.svg)](https://codecov.io/gh/dailymuse/gzap) [![Go Report Card](https://goreportcard.com/badge/github.com/dailymuse/gzap)](https://goreportcard.com/report/github.com/dailymuse/gzap)

Gzap provide fast structured leveled logging using [zap](https://github.com/uber-go/zap), and a TCP/UDP Graylog logsink (TLS supported). Both [zap](https://github.com/uber-go/zap) and [Graylog](https://github.com/Devatoria/go-graylog) librarys are versioned locked within the applications so no other external dependencies required. 

### Getting Stated

To use gzap, first import it:

```go
import "gopkg.in/dailymuse/gzap.v1"
```

> The Graylog logsink is only enabled for Production and Staging environments. So you'll need to set a `GRAYLOG_ENV` environment variable with either of the following correlating states.

GRAYLOG_ENV | Environment  | Graylog enabled?
--- | --- | ---
0 | Test (no-op logger) | ❌
1 | Dev | ❌
2 | Staging | ✅
3 | Production | ✅


### Environment Variables

To properly use `gzap` you'll need to set your configurations via Environment variables. The following are configurable Envs:

Env Name | Description
--- | --- |
GRAYLOG_ENV | A number 0 - 3 describing the Graylog loggin environment you wish to use (Refrence table above)
GRAYLOG_HOST | Hostname that your graylog is currently listening on `example.graylog.com`


### Internal API
The logger that is publicly exposed is the zap [Logger](https://godoc.org/go.uber.org/zap#Logger). You can reference what log levels are available for use [here](https://godoc.org/go.uber.org/zap#Logger)). Below are a few examples:

```
func (log *Logger) DPanic(msg string, fields ...zapcore.Field)
func (log *Logger) Debug(msg string, fields ...zapcore.Field)
func (log *Logger) Error(msg string, fields ...zapcore.Field)
func (log *Logger) Fatal(msg string, fields ...zapcore.Field)
func (log *Logger) Info(msg string, fields ...zapcore.Field)
func (log *Logger) Panic(msg string, fields ...zapcore.Field)
func (log *Logger) Warn(msg string, fields ...zapcore.Field)
```

All zap [fields](https://godoc.org/go.uber.org/zap/zapcore#Field) needed for logging are also exposed by gzap.

```go
gzap.Logger.Error("this is an example Error log",
        gzap.String("variable", "some-variable-here"),
)
```

For any other information please take a look at the gzap [Godoc](https://godoc.org/github.com/dailymuse/gzap).

### Example Usage

```go
package main

import (
    "time"

    "gopkg.in/dailymuse/gzap.v1"
)

func main() {
    // Instantiate a global logger.
    if err := gzap.InitLogger(); err != nil {
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
    gzap.Logger.Debug("this is an example Debug log",
        gzap.String("variable", "some-variable-here"),
    )
}
```

### Important info

#### Contributing

In order to contribute you'll need to have a valid go environment setup.

> If you need to install go, see installation instructions [here](https://golang.org/doc/install#install).

##### Tests Logs are a no-op

Tests that run application code containing logs will not print those logs by default. The Test logger is a no-op to reduce noise during testing.
