package gml

import (
	"crypto/tls"

	graylog "github.com/Devatoria/go-graylog"
)

// Graylog is a unifying interface for using the internal
// Graylog package 'github.com/Devatoria/go-graylog', while
// also being able to use Mocks in it's place.
type Graylog interface {
	Close() error
	Send(graylog.Message) error
}

// NewGraylog returns a new Graylog instance.
func NewGraylog(cfg *Config) (Graylog, error) {
	if cfg._isMock {
		return cfg._mock, cfg._mockErr
	}

	if cfg.UseTLS {
		return getGraylogTLS(cfg)
	}

	return getGraylog(cfg)
}

func getGraylogTLS(cfg *Config) (Graylog, error) {
	g, err := graylog.NewGraylogTLS(
		graylog.Endpoint{
			Transport: graylog.TCP,
			Address:   cfg.GraylogAddress,
			Port:      cfg.GraylogPort,
		},
		cfg.GraylogConnectionTimeout,
		&tls.Config{
			InsecureSkipVerify: cfg.InsecureSkipVerify,
		},
	)

	if err != nil {
		return nil, err
	}

	return g, nil
}

func getGraylog(cfg *Config) (Graylog, error) {
	g, err := graylog.NewGraylog(
		graylog.Endpoint{
			Transport: graylog.TCP,
			Address:   cfg.GraylogAddress,
			Port:      cfg.GraylogPort,
		},
	)

	if err != nil {
		return nil, err
	}

	return g, nil
}
