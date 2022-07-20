package midllewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRewriteInternalErrors(t *testing.T) {
	testCases := []struct {
		name    string
		prev    http.Handler
		expCode int
		expResp string
	}{
		{
			name: "ignore, if success",
			prev: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusOK)
				_, _ = writer.Write([]byte(`{"message":"ok"}`))
			}),
			expResp: `{"message":"ok"}`,
			expCode: http.StatusOK,
		},
		{
			name: "ignore, if bad request",
			prev: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusBadRequest)
				_, _ = writer.Write([]byte(`{"message":"bad request"}`))
			}),
			expResp: `{"message":"bad request"}`,
			expCode: http.StatusBadRequest,
		},
		{
			name: "rewrite if internal error",
			prev: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusInternalServerError)
				_, _ = writer.Write([]byte(`{"message":"must be overwritten"}`))
			}),
			expResp: `{"message":"Internal Server Error"}`,
			expCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			RewriteInternalErrors(tc.prev).ServeHTTP(rw, &http.Request{})

			assert.Equal(t, tc.expCode, rw.Code)
			assert.Equal(t, tc.expResp, rw.Body.String())
		})
	}
}
