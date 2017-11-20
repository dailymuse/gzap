package gzap

import (
	"errors"
	"testing"
)

func TestInitLogger(t *testing.T) {
	type args struct {
		cfg Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     string
	}{
		{
			"Init should fail if Graylog fails to connect with Prod configuration",
			args{
				Config{
					_isMock:         true,
					_mockEnv:        prodEnv,
					_mockGraylogErr: errors.New("could not connect to Graylog"),
				},
			},
			true,
			"could not connect to Graylog",
		},
		{
			"initLogger should pass if Graylog connects with Prod configuration",
			args{
				Config{
					_isMock:         true,
					_mockEnv:        prodEnv,
					_mockGraylog:    &MockGraylog{},
					_mockGraylogErr: nil,
				},
			},
			false,
			"",
		},
		{
			"initLogger should fail if Graylog fails to connect with Staging configuration",
			args{
				Config{
					_isMock:         true,
					_mockEnv:        stagingEnv,
					_mockGraylogErr: errors.New("could not connect to Graylog"),
				},
			},
			true,
			"could not connect to Graylog",
		},
		{
			"initLogger should pass if Graylog connects with Staging configuration",
			args{
				Config{
					_isMock:         true,
					_mockEnv:        stagingEnv,
					_mockGraylog:    &MockGraylog{},
					_mockGraylogErr: nil,
				},
			},
			false,
			"",
		},
		{
			"initLogger should pass if using test configuration",
			args{
				Config{
					_isMock:  true,
					_mockEnv: testEnv,
				},
			},
			false,
			"",
		},
		{
			"initLogger should pass if using dev configuration",
			args{
				Config{
					_isMock:  true,
					_mockEnv: devEnv,
				},
			},
			false,
			"",
		},
		{
			"initLogger should fail if invalid configuration is passed",
			args{
				Config{
					_isMock:  true,
					_mockEnv: 99,
				},
			},
			true,
			"no valid env was explicity set, and not currently running tests",
		},
		{
			"initLogger should fail if it cannot parse GRAYLOG_ENV",
			args{
				Config{
					_isMock:       true,
					_mockEnvError: errors.New("error occured when parsing env"),
				},
			},
			true,
			"error occured when parsing env",
		},
		{
			"initLogger should fail if dev logger throws error",
			args{
				Config{
					_isMock:     true,
					_mockEnv:    devEnv,
					_mockDevErr: errors.New("could not build development logger"),
				},
			},
			true,
			"could not build development logger",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := initLogger(tt.args.cfg)

			if tt.wantErr && err == nil {
				t.Errorf("initLogger() expected error = \"%v\"; got \"nil\"", tt.err)
			}

			if err != nil && err.Error() != tt.err {
				t.Errorf("initLogger() expected error = \"%v\";  got \"%v\"", tt.err, err.Error())
			}
		})
	}
}

func ExampleInitLogger() {
	if err := InitLogger(); err != nil {
		panic(err)
	}

	defer Logger.Sync()

	Logger.Info("this is a test info log")
}
