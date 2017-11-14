package gzap

import (
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is the global logger for the application. Upon first initalization of
// the logger all calls to 'getLogger' are memoized with the instantiated 'logger'.
var Logger = getLogger()

// logger is the package level pointer to an instantied Logger.
var logger *zap.Logger

// logInitializer represents a function that initializes
type logInitializer func(cfg *Config) error

// envToLogInitializerMapping represents the different type of log initializers
// to their correlating env level.
var envToLogInitializerMapping = map[int]logInitializer{
	testEnv:    setTestLogger,
	devEnv:     setDevelopmentLogger,
	stagingEnv: setStagingLogger,
	prodEnv:    setProductionLogger,
}

// InitLogger initializes a global Logger based upon your env configurations.
func InitLogger() error {
	return initLogger(&Config{})
}

func initLogger(cfg *Config) error {
	env, err := cfg.getGraylogEnv()
	if err != nil {
		return err
	}

	logInitalizer, ok := envToLogInitializerMapping[env]
	if !ok {
		return errors.New(envNotSetErrorString)
	}

	return logInitalizer(cfg)
}

// getLogger is an internal function that returns an instantied Logger,
// or inits a new logger with default test configuration.
//
// This is because tests that run application code that use the logger will
// need an instaniated Logger to run. In this case we want to make sure we
// use a no-op logger, to reduce test noise.
func getLogger() *zap.Logger {
	if logger == nil {
		InitLogger()
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
				String: cfg.getGraylogLogEnvName(),
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
				String: cfg.getGraylogLogEnvName(),
				Type:   zapcore.StringType,
			},
		),
	)

	logger = zapProductionLogger

	return nil
}

func setDevelopmentLogger(cfg *Config) error {
	if cfg._mockDevErr != nil {
		return cfg._mockDevErr
	}

	Config := zap.NewDevelopmentConfig()
	Config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapDevelopmentLogger, err := Config.Build()
	if err != nil {
		return err
	}

	logger = zapDevelopmentLogger

	return nil
}

func setTestLogger(cfg *Config) error {
	zapNopLogger := zap.NewNop()
	logger = zapNopLogger
	return nil
}
