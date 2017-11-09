package gomuselogger

import (
	"crypto/tls"
	"errors"
	"time"

	"github.com/Devatoria/go-graylog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is the global logger for the application.
var (
	Logger = getLogger()
)

// logger is the package level pointer to an instantied Logger.
var (
	logger *zap.Logger
)

// graylogConnectionTimeout TODO
var graylogConnectionTimeout = 3 * time.Second

// Config write the logs here TODO
type Config struct {
	GetAppName         func() string
	GetIsProdEnv       func() bool
	GetIsStagingEnv    func() bool
	GetIsTestEnv       func() bool
	GraylogAddress     string
	GraylogPort        uint
	GraylogVersion     string
	GetHostname        func() string
	UseTLS             bool
	InsecureSkipVerify bool
	GetEnvName         func() string
	isMock             bool
}

// New sets up the basic logger for either a Production or development
// environment.
func New(cfg *Config) error {
	if cfg.GetIsProdEnv() {
		if err := setProductionLogger(cfg); err != nil {
			return err
		}

		return nil
	}

	if cfg.GetIsStagingEnv() {
		if err := setStagingLogger(cfg); err != nil {
			return nil
		}

		return nil
	}

	if cfg.GetIsTestEnv() {
		setTestLogger()
		return nil
	}

	setDevelopmentLogger()

	return nil
}

// getLogger is an internal function that fetches a reference to an instantied Logger,
// or inits a new logger.
// TODO mention that it assumes an ENV is a test env if none is given.
func getLogger() *zap.Logger {
	if logger == nil {
		New(&Config{
			GetIsProdEnv: func() bool {
				return false
			},
			GetIsStagingEnv: func() bool {
				return false
			},
			GetIsTestEnv: func() bool {
				return true
			},
		})
	}

	return logger
}

func getGraylog(cfg *Config) (*graylog.Graylog, error) {
	if cfg.isMock {
		return nil, errors.New("GOT EYM")
	}

	if cfg.UseTLS {
		return getGraylogTLS(cfg)
	}

	// TODO fix this
	return getGraylogTLS(cfg)
}

// getGraylog TODO
func getGraylogTLS(cfg *Config) (*graylog.Graylog, error) {
	g, err := graylog.NewGraylogTLS(
		graylog.Endpoint{
			Transport: graylog.TCP,
			Address:   cfg.GraylogAddress,
			Port:      cfg.GraylogPort,
		},
		graylogConnectionTimeout,
		&tls.Config{
			InsecureSkipVerify: cfg.InsecureSkipVerify,
		},
	)

	if err != nil {
		return nil, err
	}

	return g, nil
}

func setProductionLogger(cfg *Config) error {
	graylog, err := getGraylog(cfg)
	if err != nil {
		return err
	}

	zapProductionLogger := zap.New(
		NewGelfCore(cfg, graylog),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.Fields(
			zapcore.Field{
				Key:    "Env",
				String: cfg.GetEnvName(),
				Type:   zapcore.StringType,
			},
		),
	)

	logger = zapProductionLogger

	return nil
}

func setStagingLogger(cfg *Config) error {
	graylog, err := getGraylog(cfg)
	if err != nil {
		return err
	}

	zapProductionLogger := zap.New(
		NewGelfCore(cfg, graylog),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.Fields(
			zapcore.Field{
				Key:    "Env",
				String: cfg.GetEnvName(),
				Type:   zapcore.StringType,
			},
		),
	)

	logger = zapProductionLogger

	return nil
}

func setDevelopmentLogger() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapDevelopmentLogger, err := config.Build()
	if err != nil {
		panic(err)
	}

	logger = zapDevelopmentLogger
}

func setTestLogger() {
	zapNopLogger := zap.NewNop()
	logger = zapNopLogger
}
