package gzap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LevedLogger is a unifying type for all of the exported zap leveled loggers,
// Info, Warn, Error, Debug.
type LevedLogger func(msg string, fields ...zapcore.Field)

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
	return initLogger(&EnvConfig{}, false)
}

func initLogger(cfg Config, disableGraylog bool) error {
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
	if graylogHost != "" && !disableGraylog {
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
			// If the logger fails to instaniate with it's current configuration
			// attempt to initLogger and skip setting up Graylog. This is to prevent
			// panicing and shutting down a service when Graylog experiences difficulties.
			// If this also fails we have a problem unrelated to Graylog and we need
			// to panic as logging has issues that are unrecoverable.
			if err != initLogger(&EnvConfig{}, true) {
				panic(err)
			}
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
	logger = zap.New(
		core,
		zap.AddCaller(),
	)
	return nil
}
