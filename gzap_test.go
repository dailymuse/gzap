package gzap

import (
	"errors"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
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
					IsProdEnv: true,
					_isMock:   true,
					_mockErr:  errors.New("could not connect to Graylog"),
				},
			},
			true,
			"could not connect to Graylog",
		},
		{
			"New should pass if Graylog connects with Prod configuration",
			args{
				&Config{
					IsProdEnv: true,
					_isMock:   true,
					_mock:     &MockGraylog{},
					_mockErr:  nil,
				},
			},
			false,
			"",
		},
		{
			"New should fail if Graylog fails to connect with Staging configuration",
			args{
				&Config{
					IsStagingEnv: true,
					_isMock:      true,
					_mockErr:     errors.New("could not connect to Graylog"),
				},
			},
			true,
			"could not connect to Graylog",
		},
		{
			"New should pass if Graylog connects with Staging configuration",
			args{
				&Config{
					IsStagingEnv: true,
					_isMock:      true,
					_mock:        &MockGraylog{},
					_mockErr:     nil,
				},
			},
			false,
			"",
		},
		{
			"New should pass if using test configuration",
			args{
				NewDefaultTestConfig(),
			},
			false,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Init(tt.args.cfg)

			if tt.wantErr && err == nil {
				t.Errorf("New() expected error = \"%v\"; got \"nil\"", tt.err)
			}

			if err != nil && err.Error() != tt.err {
				t.Errorf("New() expected error = \"%v\";  got \"%v\"", tt.err, err.Error())
			}
		})
	}
}

func ExampleInit() {
	if err := Init(&Config{
		AppName:                  "",
		IsProdEnv:                false,
		IsStagingEnv:             false,
		IsTestEnv:                false,
		GraylogAddress:           "",
		GraylogPort:              0,
		GraylogVersion:           "",
		Hostname:                 "",
		UseTLS:                   false,
		InsecureSkipVerify:       false,
		LogEnvName:               "",
		GraylogConnectionTimeout: time.Second * 0,
	}); err != nil {
		panic(err)
	}

	defer Logger.Sync()

	Logger.Info("this is a test info log")
}
