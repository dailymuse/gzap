package gomuselogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is the global logger for the application.
var Logger = getLogger()

// logger is the package level pointer to an instantied Logger.
var logger *zap.Logger

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

	// By default if we can't determine the environment explicitly we'll
	// use the development logger.
	setDevelopmentLogger()

	return nil
}

// getLogger is an internal function that fetches a reference to an instantied Logger,
// or inits a new logger.
// TODO mention that it assumes an ENV is a test env if none is given.
func getLogger() *zap.Logger {
	if logger == nil {
		New(NewDefaultTestConfig())
	}

	return logger
}

func setProductionLogger(cfg *Config) error {
	graylog, err := NewGraylog(cfg)
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
				String: cfg.GetLogEnvName(),
				Type:   zapcore.StringType,
			},
		),
	)

	logger = zapProductionLogger

	return nil
}

func setStagingLogger(cfg *Config) error {
	graylog, err := NewGraylog(cfg)
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
				String: cfg.GetLogEnvName(),
				Type:   zapcore.StringType,
			},
		),
	)

	logger = zapProductionLogger

	return nil
}

func setDevelopmentLogger() {
	Config := zap.NewDevelopmentConfig()
	Config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapDevelopmentLogger, err := Config.Build()
	if err != nil {
		panic(err)
	}

	logger = zapDevelopmentLogger
}

func setTestLogger() {
	zapNopLogger := zap.NewNop()
	logger = zapNopLogger
}
