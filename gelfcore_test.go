package gml

import (
	"testing"

	"github.com/Devatoria/go-graylog"
	"go.uber.org/zap/zapcore"
)

func TestGelfCore_Write(t *testing.T) {
	type fields struct {
		Graylog *graylog.Graylog
		Context []zapcore.Field
		cfg     *Config
	}
	type args struct {
		entry  zapcore.Entry
		fields []zapcore.Field
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gc := GelfCore{
				Graylog: tt.fields.Graylog,
				Context: tt.fields.Context,
				cfg:     tt.fields.cfg,
			}
			if err := gc.Write(tt.args.entry, tt.args.fields); (err != nil) != tt.wantErr {
				t.Errorf("GelfCore.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
