package gzap

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	graylog "github.com/Devatoria/go-graylog"
)

const testEnv = 0
const devEnv = 1
const stagingEnv = 2
const prodEnv = 3

const tlsTransport = "tls"

var envNotSetErrorString = "no valid env was explicity set, and not currently running tests"

// Config represents all the logger configurations available
// when instaniating a new Logger.
type Config struct {
	logEnvName      *string
	_isMock         bool
	_mockEnv        int
	_mockEnvError   error
	_mockGraylog    Graylog
	_mockGraylogErr error
	_mockDevErr     error
}

func (c *Config) getGraylogAppName() string {
	if c._isMock {
		return ""
	}

	appName := os.Getenv("GRAYLOG_APP_NAME")
	if appName == "" {
		panic("GRAYLOG_APP_NAME env not set")
	}

	return appName
}

func (c *Config) getGraylogEnv() (int, error) {
	if c._isMock {
		return c._mockEnv, c._mockEnvError
	}

	// If we're running test return test logger env.
	if flag.Lookup("test.v") != nil {
		return testEnv, nil
	}

	envLevelString := os.Getenv("GRAYLOG_ENV")
	if envLevelString == "" {
		return 0, errors.New(envNotSetErrorString)
	}

	envLevel, err := strconv.Atoi(envLevelString)
	if err != nil {
		return 0, errors.New("could not properly parse GRAYLOG_ENV")
	}

	return envLevel, nil
}

func (c *Config) getGraylogHandlerType() graylog.Transport {
	handlerType := os.Getenv("GRAYLOG_HANDLER_TYPE")
	if handlerType == "" {
		panic("GRAYLOG_HANDLER_TYPE env not set")
	}

	var transportType graylog.Transport
	if handlerType == tlsTransport {
		transportType = graylog.TCP
	}

	if graylog.Transport(handlerType) == graylog.UDP {
		transportType = graylog.UDP
	}

	return transportType
}

func (c *Config) getGraylogHost() string {
	graylogHost := os.Getenv("GRAYLOG_HOST")
	if graylogHost == "" {
		panic("GRAYLOG_HOST env not set")
	}

	return graylogHost
}

func (c *Config) useTLS() bool {
	handlerType := os.Getenv("GRAYLOG_HANDLER_TYPE")
	if handlerType == "" {
		panic("GRAYLOG_HANDLER_TYPE env not set")
	}

	if handlerType == tlsTransport {
		return true
	}

	return false
}

func (c *Config) getGraylogPort() uint {
	var portString string
	if c.getGraylogHandlerType() == graylog.UDP {
		portString = os.Getenv("GRAYLOG_UDP_PORT")
	}

	if c.getGraylogHandlerType() == graylog.TCP {
		portString = os.Getenv("GRAYLOG_TLS_PORT")
	}

	if portString == "" {
		panic("GRAYLOG_UDP_PORT or GRAYLOG_TLS_PORT env not set")
	}

	port, err := strconv.ParseUint(portString, 10, 32)
	if err != nil {
		panic(fmt.Errorf("could not properly parse Graylog port: %s", portString))
	}

	return uint(port)
}

func (c *Config) getGraylogTLSTimeout() time.Duration {
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

func (c *Config) getGraylogLogEnvName() string {
	// Check if we already memoized the log env name.
	if c.logEnvName != nil {
		return *c.logEnvName
	}

	env, err := c.getGraylogEnv()
	if err != nil {
		panic(err)
	}

	var logEnvName string
	if env == prodEnv {
		logEnvName = "prd"
	}

	if env == stagingEnv {
		logEnvName = "stg"
	}

	// Memoize the log env name.
	c.logEnvName = &logEnvName

	return logEnvName
}

func (c *Config) getGraylogSkipInsecureSkipVerify() bool {
	skipInsecure := os.Getenv("GRAYLOG_SKIP_TLS_VERIFY")
	if skipInsecure == "true" {
		return true
	}

	return false
}
