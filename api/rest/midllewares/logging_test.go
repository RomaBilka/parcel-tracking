package midllewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestEnableLogging(t *testing.T) {
	testCases := []struct {
		name   string
		prev   http.Handler
		logger func(m *loggerMock)

		expCode int
		expResp string
	}{
		{
			name: "log on bad request",
			prev: http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
				writer.WriteHeader(http.StatusBadRequest)
				_, _ = writer.Write([]byte(`{"message":"test error"}`))
			}),
			logger: func(m *loggerMock) {
				m.On("Warn", "test error",
					zap.Int("statusCode", http.StatusBadRequest), mock.Anything)
			},
			expResp: `{"message":"test error"}`,
			expCode: http.StatusBadRequest,
		},
		{
			name: "log on internal error",
			prev: http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
				writer.WriteHeader(http.StatusInternalServerError)
				_, _ = writer.Write([]byte(`{"message":"test error"}`))
			}),
			logger: func(m *loggerMock) {
				m.On("Error", "test error",
					zap.Int("statusCode", http.StatusInternalServerError), mock.Anything)
			},
			expResp: `{"message":"test error"}`,
			expCode: http.StatusInternalServerError,
		},
		{
			name: "success, no logs",
			prev: http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
				writer.WriteHeader(http.StatusOK)
				_, _ = writer.Write([]byte(`{"message":"success"}`))
			}),
			logger:  func(m *loggerMock) {},
			expResp: `{"message":"success"}`,
			expCode: http.StatusOK,
		},
		{
			name: "failed to unmarshal error",
			prev: http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
				writer.WriteHeader(http.StatusNotFound)
				_, _ = writer.Write([]byte(`not a json`))
			}),
			logger: func(m *loggerMock) {
				m.On("Error", "failed to unmarshall error",
					zap.Int("statusCode", http.StatusNotFound), mock.Anything, mock.Anything,
					zap.String("response", "not a json"))
			},
			expResp: `not a json`,
			expCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := &loggerMock{}
			tc.logger(l)

			rw := httptest.NewRecorder()
			Logging(l)(tc.prev).ServeHTTP(rw, &http.Request{})

			assert.Equal(t, tc.expCode, rw.Code)
			assert.Equal(t, tc.expResp, rw.Body.String())

			l.AssertExpectations(t)
		})
	}
}

type loggerMock struct {
	mock.Mock
}

func (l *loggerMock) Error(msg string, fields ...zap.Field) {
	args := []interface{}{msg}
	for _, field := range fields {
		args = append(args, field)
	}
	l.Called(args...)
}

func (l *loggerMock) Warn(msg string, fields ...zap.Field) {
	args := []interface{}{msg}
	for _, field := range fields {
		args = append(args, field)
	}
	l.Called(args...)
}
