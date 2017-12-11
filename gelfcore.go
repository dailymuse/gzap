package gzap

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/Devatoria/go-graylog"
	"go.uber.org/zap/zapcore"
)

// GelfCore implements the https://godoc.org/go.uber.org/zap/zapcore#Core interface
// Messages are written to a graylog endpoint using the GELF format + protocol
type GelfCore struct {
	Graylog            Graylog
	Context            []zapcore.Field
	cfg                Config
	encoder            zapcore.Encoder
	graylogConstructor GraylogConstructor
}

// NewGelfCore creates a new GelfCore with empty context.
func NewGelfCore(cfg Config, gl Graylog) GelfCore {
	encoderConfigs := zapcore.EncoderConfig{
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewJSONEncoder(encoderConfigs)

	return GelfCore{
		Graylog:            gl,
		cfg:                cfg,
		encoder:            encoder,
		graylogConstructor: NewGraylog,
	}
}

// map zapcore's log levels to standard syslog levels used by gelf, approximately.
var zapToSyslog = map[zapcore.Level]uint{
	zapcore.DebugLevel:  7,
	zapcore.InfoLevel:   6,
	zapcore.WarnLevel:   4,
	zapcore.ErrorLevel:  3,
	zapcore.DPanicLevel: 2,
	zapcore.PanicLevel:  2,
	zapcore.FatalLevel:  1,
}

// Write writes messages to the configured Graylog endpoint.
func (gc GelfCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	extraFields := map[string]string{
		"file":        entry.Caller.File,
		"line":        strconv.Itoa(entry.Caller.Line),
		"logger_name": entry.LoggerName,
		"app_name":    gc.cfg.getGraylogAppName(),
	}

	// the order here is important,
	// as fields supplied at the log site should overwrite fields supplied in the context
	for _, field := range gc.Context {
		extraFields[field.Key] = field.String
	}

	for _, field := range fields {
		extraFields[field.Key] = field.String
	}

	// Encode the zap fields from fields to JSON with proper types.
	buf, err := gc.encoder.EncodeEntry(entry, fields)
	if err != nil {
		panic(err)
	}

	// Unmarshal the JSON into a map.
	m := make(map[string]interface{})
	if err = json.Unmarshal(buf.Bytes(), &m); err != nil {
		return err
	}

	// Parse the map and return only strings.
	for k, v := range m {
		extraFields[k] = fmt.Sprintf("%v", v)
	}

	msg := graylog.Message{
		Version:      "1.1",
		Host:         hostname,
		ShortMessage: entry.Message,
		FullMessage:  entry.Stack,
		Timestamp:    entry.Time.Unix(),
		Level:        zapToSyslog[entry.Level],
		Extra:        extraFields,
	}

	if err := gc.Graylog.Send(msg); err != nil {
		if err := attemptRetry(gc.cfg, gc, msg, gc.graylogConstructor); err != nil {
			panic(err)
		}
	}

	return nil
}

// With adds structured context to the logger.
func (gc GelfCore) With(fields []zapcore.Field) zapcore.Core {
	gc.Context = append(gc.Context, fields...)
	return gc
}

// Sync is a no-op.
func (gc GelfCore) Sync() error {
	return nil
}

// Check determines whether the supplied entry should be logged.
func (gc GelfCore) Check(entry zapcore.Entry, checkedEntry *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if gc.Enabled(entry.Level) {
		return checkedEntry.AddCore(entry, gc)
	}

	return checkedEntry
}

// Enabled only enables info messages and above.
func (gc GelfCore) Enabled(level zapcore.Level) bool {
	return zapcore.InfoLevel.Enabled(level)
}

func attemptRetry(cfg Config, gc GelfCore, msg graylog.Message, newGraylog GraylogConstructor) error {
	attempts := 3
	var retryErr error

	for i := 0; i < attempts; i++ {
		// Attempt to create new client.
		graylog, err := newGraylog(cfg)
		if err != nil {
			retryErr = err
			continue
		}

		// Attempt to send message.
		err = graylog.Send(msg)
		if err != nil {
			retryErr = err
			continue
		}

		// Assign new client to Gelfcore.
		gc.Graylog = graylog
		retryErr = nil
		break
	}

	return retryErr
}
