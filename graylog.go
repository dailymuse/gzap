package gomuselogger

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

	if cfg.GetUseTLS() {
		return getGraylogTLS(cfg)
	}

	return getGraylog(cfg)
}

// getGraylogTLS TODO
func getGraylogTLS(cfg *Config) (*graylog.Graylog, error) {
	g, err := graylog.NewGraylogTLS(
		graylog.Endpoint{
			Transport: graylog.TCP,
			Address:   cfg.GetGraylogAddress(),
			Port:      cfg.GetGraylogPort(),
		},
		cfg.GetGraylogConnectionTimeout(),
		&tls.Config{
			InsecureSkipVerify: cfg.GetInsecureSkipVerify(),
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
			Address:   cfg.GetGraylogAddress(),
			Port:      cfg.GetGraylogPort(),
		},
	)

	if err != nil {
		return nil, err
	}

	return g, nil
}
