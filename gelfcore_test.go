package gzap

import (
	"errors"
	"testing"

	"go.uber.org/zap"

	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zapcore"
)

func TestGelfCore_Write(t *testing.T) {
	type fields struct {
		Graylog Graylog
		Context []zapcore.Field
		cfg     Config
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
		err     string
		sendErr error
	}{
		{
			"GelfCore.Write should return an error when failing to send message",
			fields{
				&MockGraylog{},
				[]zapcore.Field{},
				&MockEnvConfig{},
			},
			args{
				zapcore.Entry{},
				[]zapcore.Field{},
			},
			true,
			"failed to send Graylog message",
			errors.New("failed to send Graylog message"),
		},
		{
			"GelfCore.Write should return nil when successfully sending a log",
			fields{
				&MockGraylog{},
				[]zapcore.Field{},
				&MockEnvConfig{},
			},
			args{
				zapcore.Entry{},
				[]zapcore.Field{},
			},
			false,
			"",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGraylog := NewMockGraylog()
			mockGraylog.On("Send", mock.AnythingOfType("graylog.Message")).Return(tt.sendErr)
			tt.fields.Graylog = &mockGraylog

			mockEnvConfig := &MockEnvConfig{}
			mockEnvConfig.On("getGraylogAppName").Return("LOL")
			mockEnvConfig.On("getIsTestEnv").Return(false)
			tt.fields.cfg = mockEnvConfig

			gc := GelfCore{
				Graylog: tt.fields.Graylog,
				Context: tt.fields.Context,
				cfg:     tt.fields.cfg,
				encoder: zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
			}

			err := gc.Write(tt.args.entry, tt.args.fields)

			if tt.wantErr && err == nil {
				t.Errorf("GelfCore.Write() expected error = \"%v\"; got \"nil\"", tt.err)
			}

			if err != nil && err.Error() != tt.err {
				t.Errorf("GelfCore.Write() expected error = \"%v\";  got \"%v\"", tt.err, err.Error())
			}
		})
	}
}
