package gml

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is the global logger for the application. Upon first initalization of
// the logger all calls to 'getLogger' are memoized with the instantiated 'logger'.
var Logger = getLogger()

// logger is the package level pointer to an instantied Logger.
var logger *zap.Logger

// Init initializes a global Logger based upon the configurations you pass in.
func Init(cfg *Config) error {
	if cfg.IsProdEnv {
		if err := setProductionLogger(cfg); err != nil {
			return err
		}

		return nil
	}

	if cfg.IsStagingEnv {
		if err := setStagingLogger(cfg); err != nil {
			return err
		}

		return nil
	}

	if cfg.IsTestEnv {
		setTestLogger()
		return nil
	}

	// By default if we can't determine the environment explicitly we'll
	// use the development logger.
	setDevelopmentLogger()

	return nil
}

// getLogger is an internal function that returns an instantied Logger,
// or inits a new logger with default test configuration.
//
// This is because tests that run application code that use the logger will
// need an instaniated Logger to run. In this case we want to make sure we
// use a no-op logger, to reduce test noise.
func getLogger() *zap.Logger {
	if logger == nil {
		Init(NewDefaultTestConfig())
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
				String: cfg.LogEnvName,
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
				String: cfg.LogEnvName,
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
