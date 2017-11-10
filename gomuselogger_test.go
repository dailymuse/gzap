package gomuselogger

import (
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     string
	}{
		{
			"New should fail if Graylog fails to connect with Prod configuration",
			args{
				&Config{
					GetIsProdEnv: func() bool {
						return true
					},
					_isMock: true,
				},
			},
			true,
			"fake error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := New(tt.args.cfg); tt.wantErr && err.Error() != tt.err {
				t.Errorf("New() error = \"%v\", wantErrString \"%v\"", err, tt.err)
			}
		})
	}
}
