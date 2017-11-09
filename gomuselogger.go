package gomuselogger

import (
	"crypto/tls"
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
	GetIsProdEnv       func() bool
	GetIsStagingEnv    func() bool
	GetIsTestEnv       func() bool
	GraylogAddress     string
	GraylogPort        uint
	GraylogVersion     string
	GetHostname        func() string
	UseTLS             bool
	InsecureSkipVerify bool
}

//func NewDefault()

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

// getLogger fetchess a reference to an instantied Logger, or inits a new logger.
// ASSUMES IT'S RUNNING IN TEST
func getLogger() *zap.Logger {
	if logger == nil {
		//New()
	}

	return logger
}

// func getGraylog ()

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
	graylog, err := getGraylogTLS(cfg)
	if err != nil {
		return err
	}

	zapProductionLogger := zap.New(
		NewGelfCore(graylog),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.Fields(zapcore.Field{Key: "Env", String: "production", Type: zapcore.StringType}))

	logger = zapProductionLogger

	return nil
}

func setStagingLogger(cfg *Config) error {
	graylog, err := getGraylogTLS(cfg)
	if err != nil {
		return err
	}

	zapProductionLogger := zap.New(
		NewGelfCore(graylog),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.Fields(zapcore.Field{Key: "Env", String: "staging", Type: zapcore.StringType}),
	)

	logger = zapProductionLogger

	return nil
}

func setDevelopmentLogger() {
	config := zap.NewDevelopmentConfig()
	// This allows us to have colored logs for development.
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

// DefaultGetIsProdEnv TODO
func DefaultGetIsProdEnv() bool {
	return false
}

// DefaultGetIsTest TODO
func DefaultGetIsTest() bool {
	return false
}

// DefaultGetIsStagingEnv TODO
func DefaultGetIsStagingEnv() bool {
	return false
}
