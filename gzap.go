package gzap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is the global logger for the application. Upon first initalization of
// the logger all calls to 'getLogger' are memoized with the instantiated 'logger'.
var Logger = getLogger()

// logger is the package level pointer to an instantied Logger.
var logger *zap.Logger

var highPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	return lvl >= zapcore.ErrorLevel
})

var lowPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	return lvl < zapcore.ErrorLevel
})

// InitLogger initializes a global Logger based upon your env configurations.
func InitLogger() error {
	return initLogger(&EnvConfig{})
}

func initLogger(cfg Config) error {
	// If we're running in a Go test return the test logger.
	if cfg.getIsTestEnv() {
		return setTestLogger(cfg)
	}

	// Create a console output enabled zapcore.
	zapcore := enableConsoleLogging(cfg)

	// Check if Graylog host is defined
	// if so return a Graylog Logger with
	// console logging enabled.
	graylogHost := cfg.getGraylogHost()
	if graylogHost != "" {
		return setGraylogLogger(cfg, zapcore)
	}

	// Return a console logger by default.
	return setLoggerFromCore(zapcore)
}

// getLogger is an internal function that returns an instantied Logger,
// or inits a new logger with default test configuration.
//
// This is because tests that run application code that use the logger will
// need an instaniated Logger to run. In this case we want to make sure we
// use a no-op logger, to reduce test noise.
func getLogger() *zap.Logger {
	if logger == nil {
		if err := InitLogger(); err != nil {
			panic(err)
		}
	}

	return logger
}

func setGraylogLogger(cfg Config, consoleLoggingCore zapcore.Core) error {
	graylog, err := NewGraylog(cfg)
	if err != nil {
		return err
	}

	zapcore := zap.New(
		zapcore.NewTee(
			NewGelfCore(cfg, graylog),
			consoleLoggingCore,
		),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.Fields(
			zapcore.Field{
				Key:    "env",
				String: cfg.getGraylogLogEnvName(),
				Type:   zapcore.StringType,
			},
		),
	)

	logger = zapcore

	return nil
}

func enableConsoleLogging(cfg Config) zapcore.Core {
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	encoderConfig := zap.NewDevelopmentEncoderConfig()

	if cfg.useColoredConsolelogs() {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	zapcore := zapcore.NewTee(
		zapcore.NewCore(
			consoleEncoder,
			consoleDebugging,
			lowPriority,
		),
		zapcore.NewCore(
			consoleEncoder,
			consoleErrors,
			highPriority,
		),
	)

	return zapcore
}

func setTestLogger(cfg Config) error {
	zapNopLogger := zap.NewNop()
	logger = zapNopLogger
	return nil
}

func setLoggerFromCore(core zapcore.Core) error {
	logger = zap.New(core)
	return nil
}
