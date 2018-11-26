package gzap

import (
	"time"

	graylog "github.com/Devatoria/go-graylog"
	"github.com/stretchr/testify/mock"
)

// MockEnvConfig represents all the logger configurations available
// when instaniating a new Logger via a mock.
type MockEnvConfig struct {
	mock.Mock
}

func (m *MockEnvConfig) enableJSONFormatter() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockEnvConfig) getGraylogAppName() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockEnvConfig) getGraylogHandlerType() graylog.Transport {
	args := m.Called()
	return args.Get(0).(graylog.Transport)
}

func (m *MockEnvConfig) getGraylogHost() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockEnvConfig) useTLS() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockEnvConfig) getGraylogPort() uint {
	args := m.Called()
	return args.Get(0).(uint)
}

func (m *MockEnvConfig) getGraylogTLSTimeout() time.Duration {
	args := m.Called()
	return args.Get(0).(time.Duration)
}

func (m *MockEnvConfig) getGraylogLogEnvName() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockEnvConfig) getGraylogSkipInsecureSkipVerify() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockEnvConfig) getIsTestEnv() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockEnvConfig) useColoredConsolelogs() bool {
	args := m.Called()
	return args.Bool(0)
}
