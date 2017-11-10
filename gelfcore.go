package gml

import (
	"strconv"

	"github.com/Devatoria/go-graylog"
	"go.uber.org/zap/zapcore"
)

// GelfCore implements the https://godoc.org/go.uber.org/zap/zapcore#Core interface
// Messages are written to a graylog endpoint using the GELF format + protocol
type GelfCore struct {
	Graylog Graylog
	Context []zapcore.Field
	cfg     *Config
}

// NewGelfCore creates a new GelfCore with empty context
func NewGelfCore(cfg *Config, gl Graylog) GelfCore {
	return GelfCore{
		Graylog: gl,
		cfg:     cfg,
	}
}

// map zapcore's log levels to standard syslog levels used by gelf, approximately
var zapToSyslog = map[zapcore.Level]uint{
	zapcore.DebugLevel:  7,
	zapcore.InfoLevel:   6,
	zapcore.WarnLevel:   4,
	zapcore.ErrorLevel:  3,
	zapcore.DPanicLevel: 2,
	zapcore.PanicLevel:  2,
	zapcore.FatalLevel:  1,
}

// Write the message to the endpoint
// TODO test this method
func (gc GelfCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	extraFields := map[string]string{
		"File":       entry.Caller.File,
		"Line":       strconv.Itoa(entry.Caller.Line),
		"LoggerName": entry.LoggerName,
		"AppName":    gc.cfg.AppName,
	}

	// the order here is important,
	// as fields supplied at the log site should overwrite fields supplied in the context
	for _, field := range gc.Context {
		extraFields[field.Key] = field.String
	}

	for _, field := range fields {
		extraFields[field.Key] = field.String
	}

	if err := gc.Graylog.Send(graylog.Message{
		Version:      gc.cfg.GraylogVersion,
		Host:         gc.cfg.Hostname,
		ShortMessage: entry.Message,
		FullMessage:  entry.Stack,
		Timestamp:    entry.Time.Unix(),
		Level:        zapToSyslog[entry.Level],
		Extra:        extraFields,
	}); err != nil {
		return err
	}

	return nil
}

// With adds structured context to the logger
func (gc GelfCore) With(fields []zapcore.Field) zapcore.Core {
	return GelfCore{
		Graylog: gc.Graylog,
		Context: append(gc.Context, fields...),
	}
}

// Sync is a no-op for us
func (gc GelfCore) Sync() error {
	return nil
}

// Check determines whether the supplied entry should be logged
func (gc GelfCore) Check(entry zapcore.Entry, checkedEntry *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if gc.Enabled(entry.Level) {
		return checkedEntry.AddCore(entry, gc)
	}

	return checkedEntry
}

// Enabled only enables info messages and above
func (gc GelfCore) Enabled(level zapcore.Level) bool {
	return zapcore.InfoLevel.Enabled(level)
}
