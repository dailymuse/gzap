package gomuselogger

import (
	"fmt"
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
			"my first test",
			args{
				&Config{
					GetIsProdEnv: func() bool {
						return true
					},
					isMock: true,
				},
			},
			true,
			"something",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := New(tt.args.cfg); tt.wantErr && err.Error() != tt.err {
				t.Errorf("New() error = \"%v\", wantErrString \"%v\"", err, tt.err)
			}

			err := New(tt.args.cfg)
			fmt.Println(err.Error())
		})
	}
}
