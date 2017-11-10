package gml

import "time"

// Config TODO
type Config struct {
	GetAppName                  func() string
	GetIsProdEnv                func() bool
	GetIsStagingEnv             func() bool
	GetIsTestEnv                func() bool
	GetGraylogAddress           func() string
	GetGraylogPort              func() uint
	GetGraylogVersion           func() string
	GetHostname                 func() string
	GetUseTLS                   func() bool
	GetInsecureSkipVerify       func() bool
	GetLogEnvName               func() string
	GetGraylogConnectionTimeout func() time.Duration
	_isMock                     bool
	_mockErr                    string
}

// NewConfig TODO
func NewConfig(
	GetAppName func() string,
	GetIsProdEnv func() bool,
	GetIsStagingEnv func() bool,
	GetIsTestEnv func() bool,
	GetGraylogAddress func() string,
	GetGraylogPort func() uint,
	GetGraylogVersion func() string,
	GetHostname func() string,
	GetUseTLS func() bool,
	GetInsecureSkipVerify func() bool,
	GetLogEnvName func() string,
	GetGraylogConnectionTimeout func() time.Duration,
) *Config {
	return &Config{
		GetAppName,
		GetIsProdEnv,
		GetIsStagingEnv,
		GetIsTestEnv,
		GetGraylogAddress,
		GetGraylogPort,
		GetGraylogVersion,
		GetHostname,
		GetUseTLS,
		GetInsecureSkipVerify,
		GetLogEnvName,
		GetGraylogConnectionTimeout,
		false,
		"",
	}
}

// NewDefaultTestConfig TODO
func NewDefaultTestConfig() *Config {
	return &Config{
		GetIsProdEnv: func() bool {
			return false
		},
		GetIsStagingEnv: func() bool {
			return false
		},
		GetIsTestEnv: func() bool {
			return true
		},
	}
}
