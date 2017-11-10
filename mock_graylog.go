package gml

import (
	graylog "github.com/Devatoria/go-graylog"
	"github.com/stretchr/testify/mock"
)

// MockGraylog TODO
type MockGraylog struct {
	mock.Mock
}

// NewMockGraylog TODO
func NewMockGraylog() MockGraylog {
	return MockGraylog{}
}

// Close closes the opened connections of the given client
func (m *MockGraylog) Close() error {
	args := m.Called()
	return args.Error(0)
}

// Send writes the given message to the given graylog client
func (m *MockGraylog) Send(msg graylog.Message) error {
	args := m.Called(msg)
	return args.Error(0)
}
