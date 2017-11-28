package gzap

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
func NewGraylog(cfg Config) (Graylog, error) {
	var gl Graylog
	var err error

	if cfg.getGraylogHandlerType() == graylog.UDP {
		gl, err = getGraylogUDP(cfg)
	}

	if cfg.getGraylogHandlerType() == graylog.TCP {
		gl, err = getGraylogTLS(cfg)
	}

	return gl, err
}

func getGraylogTLS(cfg Config) (Graylog, error) {
	g, err := graylog.NewGraylogTLS(
		graylog.Endpoint{
			Transport: graylog.TCP,
			Address:   cfg.getGraylogHost(),
			Port:      cfg.getGraylogPort(),
		},
		cfg.getGraylogTLSTimeout(),
		&tls.Config{
			InsecureSkipVerify: cfg.getGraylogSkipInsecureSkipVerify(),
		},
	)

	if err != nil {
		return nil, err
	}

	return g, nil
}

func getGraylogUDP(cfg Config) (Graylog, error) {
	g, err := graylog.NewGraylog(
		graylog.Endpoint{
			Transport: graylog.UDP,
			Address:   cfg.getGraylogHost(),
			Port:      cfg.getGraylogPort(),
		},
	)

	if err != nil {
		return nil, err
	}

	return g, nil
}
