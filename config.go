package gml

import "time"

// Config TODO
type Config struct {
	AppName                  string
	IsProdEnv                bool
	IsStagingEnv             bool
	IsTestEnv                bool
	GraylogAddress           string
	GraylogPort              uint
	GraylogVersion           string
	Hostname                 string
	UseTLS                   bool
	InsecureSkipVerify       bool
	LogEnvName               string
	GraylogConnectionTimeout time.Duration
	_isMock                  bool
	_mock                    Graylog
	_mockErr                 error
}

// NewConfig TODO
func NewConfig(
	AppName string,
	IsProdEnv bool,
	IsStagingEnv bool,
	IsTestEnv bool,
	GraylogAddress string,
	GraylogPort uint,
	GraylogVersion string,
	Hostname string,
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
		GraylogAddress,
		GraylogPort,
		GraylogVersion,
		Hostname,
		UseTLS,
		InsecureSkipVerify,
		LogEnvName,
		GraylogConnectionTimeout,
		false,
		nil,
		nil,
	}
}

// NewDefaultTestConfig TODO
func NewDefaultTestConfig() *Config {
	return &Config{
		IsProdEnv:    false,
		IsStagingEnv: false,
		IsTestEnv:    false,
	}
}
