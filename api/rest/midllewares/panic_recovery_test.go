package midllewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPanicRecovery(t *testing.T) {
	testCases := []struct {
		name    string
		prev    http.Handler
		logger  func(m *unknownArgsLoggerMock)
		expCode int
		expResp string
	}{
		{
			name: "recover",
			prev: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				panic("panic test")
			}),
			logger: func(m *unknownArgsLoggerMock) {
				m.On("Error", "panic test")
			},
			expResp: ``,
			expCode: http.StatusInternalServerError,
		},
		{
			name:   "do not recover",
			logger: func(m *unknownArgsLoggerMock) {},
			prev: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusBadRequest)
				_, _ = writer.Write([]byte(`{"message":"bad request"}`))
			}),
			expResp: `{"message":"bad request"}`,
			expCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			logger := &unknownArgsLoggerMock{}
			tc.logger(logger)

			PanicRecovery(logger)(tc.prev).ServeHTTP(rw, &http.Request{})

			assert.Equal(t, tc.expCode, rw.Code)
			assert.Equal(t, tc.expResp, rw.Body.String())

			logger.AssertExpectations(t)
		})
	}
}

type unknownArgsLoggerMock struct {
	mock.Mock
}

func (u *unknownArgsLoggerMock) Error(args ...interface{}) {
	u.Called(args...)
}
