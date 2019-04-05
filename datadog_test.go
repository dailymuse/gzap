package gzap

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"

)

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v - Got %v ", b, a)
	}
}

func Test_Datadog(t *testing.T) {
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("UGOOD"))
	})

	wrappedHandler := DatadogRequestLoggerHandler(handler)

	cfg := MockEnvConfig{}
	cfg.On("enableJSONFormatter").Return(true)
	cfg.On("getGraylogHost").Return("")
	cfg.On("getIsTestEnv").Return(false)
	cfg.On("useColoredConsolelogs").Return(true)

	err := initLogger(&cfg, true)
	if err != nil {
		t.Error(err)
	}

	// Add an observer so that we can check that the expected fields are present
	core, logs := observer.New(zap.InfoLevel)

	// We can set the logger core to an observer, as we do not actually care about the json formatting
	// aspect. What we are testing is that the log context (i.e.,the zapcore fields) gets populated.
	setLoggerFromCore(core)

	req, err := http.NewRequest("GET", "http://localhost:3000/foobar", strings.NewReader("Read ME"))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("X-Forwarded-For", "127.0.0.1,53.53.22.9,8.8.8.8")
	req.Header.Set("X-Amzn-Trace-Id", "Root=GETSOME")

	wrappedHandler.ServeHTTP(recorder, req)

	all := logs.AllUntimed()
	expect(t, len(all), 1)

	ctx := all[0].ContextMap()
	expect(t, ctx["network.client.ip"], "127.0.0.1")
	expect(t, ctx["http.status_code"], int64(200))
	expect(t, ctx["http.method"], "GET")
	expect(t, ctx["http.url"], "http://localhost:3000/foobar")
	expect(t, ctx["http.request_id"], "Root=GETSOME")
	expect(t, ctx["network.bytes_written"], int64(5))
	expect(t, ctx["network.bytes_read"], int64(7))

}
