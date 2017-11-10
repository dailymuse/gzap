package gml

import (
	"crypto/tls"
	"errors"

	graylog "github.com/Devatoria/go-graylog"
)

// NewGraylog TODO
func NewGraylog(cfg *Config) (*graylog.Graylog, error) {
	if cfg._isMock {
		return nil, errors.New("fake error")
	}

	if cfg.UseTLS {
		return getGraylogTLS(cfg)
	}

	return getGraylog(cfg)
}

// getGraylogTLS MORE TODO
func getGraylogTLS(cfg *Config) (*graylog.Graylog, error) {
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

// getGraylog TODO
func getGraylog(cfg *Config) (*graylog.Graylog, error) {
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
