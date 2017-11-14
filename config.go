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

// Config represents all the logger configurations available
// when instaniating a new Logger.
type Config struct {
	AppName                  string
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

// NewConfig returns a new logging Config with the supplied arugments.
func NewConfig(
	AppName string,
	IsProdEnv bool,
	IsStagingEnv bool,
	IsTestEnv bool,
	IsDevEnv bool,
	GraylogAddress string,
	GraylogPort uint,
	UseTLS bool,
	InsecureSkipVerify bool,
	LogEnvName string,
	GraylogConnectionTimeout time.Duration,
) *Config {
	return &Config{
		AppName,
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

// NewDefaultTestConfig returns a noop logging Config used for run testing.
func NewDefaultTestConfig() *Config {
	return &Config{}
}

func getGraylogEnv(cfg *Config) (int, error) {
	if cfg._isMock {
		return cfg._mockEnv, cfg._mockEnvError
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
