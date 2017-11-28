package gzap

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	graylog "github.com/Devatoria/go-graylog"
)

const tlsTransport = "tls"

// Config is an interface representing all the logging configurations accessible
// via enironment
type Config interface {
	getGraylogAppName() string
	getGraylogHandlerType() graylog.Transport
	getGraylogHost() string
	useTLS() bool
	getGraylogPort() uint
	getGraylogTLSTimeout() time.Duration
	getGraylogLogEnvName() string
	getGraylogSkipInsecureSkipVerify() bool
	getIsTestEnv() bool
	useColoredConsolelogs() bool
}

// EnvConfig represents all the logger configurations available
// when instaniating a new Logger.
type EnvConfig struct{}

func (e EnvConfig) getGraylogAppName() string {
	appName := os.Getenv("GRAYLOG_APP_NAME")
	if appName == "" {
		panic("GRAYLOG_APP_NAME env not set")
	}

	return appName
}

func (e *EnvConfig) getGraylogHandlerType() graylog.Transport {
	defaultHandlerType := tlsTransport
	handlerType := os.Getenv("GRAYLOG_HANDLER_TYPE")

	var transportType graylog.Transport
	if handlerType == tlsTransport {
		transportType = graylog.TCP
	}

	if graylog.Transport(handlerType) == graylog.UDP {
		transportType = graylog.UDP
	}

	// If no transport type is set use tls by default.
	if transportType == "" {
		transportType = graylog.Transport(defaultHandlerType)
	}

	return transportType
}

func (e *EnvConfig) getGraylogHost() string {
	graylogHost := os.Getenv("GRAYLOG_HOST")
	return graylogHost
}

func (e *EnvConfig) useTLS() bool {
	handlerType := os.Getenv("GRAYLOG_HANDLER_TYPE")
	if handlerType == "" {
		panic("GRAYLOG_HANDLER_TYPE env not set")
	}

	if handlerType == tlsTransport {
		return true
	}

	return false
}

func (e *EnvConfig) getGraylogPort() uint {
	portString := "12201"

	if e.getGraylogHandlerType() == graylog.UDP {
		portString = os.Getenv("GRAYLOG_UDP_PORT")
	}

	if e.getGraylogHandlerType() == graylog.TCP {
		portString = os.Getenv("GRAYLOG_TLS_PORT")
	}

	port, err := strconv.ParseUint(portString, 10, 32)
	if err != nil {
		panic(fmt.Errorf("could not properly parse Graylog port: %s", portString))
	}

	return uint(port)
}

func (e *EnvConfig) getGraylogTLSTimeout() time.Duration {
	defaultTimeout := time.Second * 3

	timeoutString := os.Getenv("GRAYLOG_TLS_TIMEOUT_SECS")
	if timeoutString == "" {
		return defaultTimeout
	}

	timeoutSeconds, err := strconv.ParseInt(timeoutString, 10, 32)
	if err != nil {
		panic("invalid GRAYLOG_TLS_TIMEOUT_SECS could not parse int")
	}

	return time.Second * time.Duration(timeoutSeconds)
}

func (e *EnvConfig) getGraylogLogEnvName() string {
	envName := os.Getenv("GRAYLOG_ENV")
	if envName == "" {
		panic("GRAYLOG_ENV not set")
	}

	return envName
}

func (e *EnvConfig) getGraylogSkipInsecureSkipVerify() bool {
	skipInsecure := os.Getenv("GRAYLOG_SKIP_TLS_VERIFY")
	if skipInsecure == "true" {
		return true
	}

	return false
}

func (e *EnvConfig) getIsTestEnv() bool {
	// If we're running test return test logger env.
	if flag.Lookup("test.v") != nil {
		return true
	}

	return false
}

func (e *EnvConfig) useColoredConsolelogs() bool {
	envLevel := os.Getenv("THEMUSE_ENV_LEVEL")

	// If the env level is not set use colored logs.
	if envLevel == "" {
		return true
	}

	return false
}
