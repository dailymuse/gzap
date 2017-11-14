package gzap

import (
	"errors"
	"flag"
	"os"
	"strconv"
	"time"
)

const testEnv = 0
const devEnv = 1
const stagingEnv = 2
const prodEnv = 3

var envNotSetErrorString = "no valid env was explicity set, and not currently running tests"
var hostname = "LOL"

// Config represents all the logger configurations available
// when instaniating a new Logger.
type Config struct {
	GraylogAddress           string
	GraylogPort              uint
	UseTLS                   bool
	InsecureSkipVerify       bool
	LogEnvName               string
	GraylogConnectionTimeout time.Duration
	_isMock                  bool
	_mockEnv                 int
	_mockEnvError            error
	_mockGraylog             Graylog
	_mockGraylogErr          error
	_mockDevErr              error
}

// NewDefaultTestConfig returns a noop logging Config used for run testing.
func NewDefaultTestConfig() *Config {
	return &Config{}
}

// NewConfig returns a new logging Config with the supplied arugments.
func NewConfig(
	GraylogAddress string,
	GraylogPort uint,
	UseTLS bool,
	InsecureSkipVerify bool,
	LogEnvName string,
	GraylogConnectionTimeout time.Duration,
) *Config {
	return &Config{
		GraylogAddress,
		GraylogPort,
		UseTLS,
		InsecureSkipVerify,
		LogEnvName,
		GraylogConnectionTimeout,
		false,
		0,
		nil,
		nil,
		nil,
		nil,
	}
}

func (c Config) getGraylogAppName() string {
	if c._isMock {
		return ""
	}

	appName := os.Getenv("GRAYLOG_APP_NAME")
	if appName == "" {
		panic("GRAYLOG_APP_NAME not set")
	}

	return appName
}

func (c Config) getGraylogEnv() (int, error) {
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
