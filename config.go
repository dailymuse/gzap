package gzap

import "time"

// Config represents all the logger configurations available
// when instaniating a new Logger.
type Config struct {
	AppName                  string
	IsProdEnv                bool
	IsStagingEnv             bool
	IsTestEnv                bool
	IsDevEnv                 bool
	GraylogAddress           string
	GraylogPort              uint
	GraylogVersion           string
	UseTLS                   bool
	InsecureSkipVerify       bool
	LogEnvName               string
	GraylogConnectionTimeout time.Duration
	_isMock                  bool
	_mock                    Graylog
	_mockErr                 error
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
	GraylogVersion string,
	UseTLS bool,
	InsecureSkipVerify bool,
	LogEnvName string,
	GraylogConnectionTimeout time.Duration,
) *Config {
	return &Config{
		AppName,
		IsProdEnv,
		IsStagingEnv,
		IsTestEnv,
		IsDevEnv,
		GraylogAddress,
		GraylogPort,
		GraylogVersion,
		UseTLS,
		InsecureSkipVerify,
		LogEnvName,
		GraylogConnectionTimeout,
		false,
		nil,
		nil,
	}
}

// NewDefaultTestConfig returns a noop logging Config used for run testing.
func NewDefaultTestConfig() *Config {
	return &Config{
		IsProdEnv:    false,
		IsStagingEnv: false,
		IsTestEnv:    true,
		IsDevEnv:     false,
	}
}
