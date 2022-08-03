package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func Test_DoSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != XmlContentType {
			_, _ = w.Write([]byte("Content-Type header is not correct"))
			return
		}
		_, _ = w.Write([]byte("OK"))
	}))

	defer server.Close()

	resp, err := Do(server.URL, "GET", func(req *fasthttp.Request) {
		req.Header.SetContentType(XmlContentType)
	})

	assert.Equal(t, "OK", string(resp.Body()))
	assert.Nil(t, err)
}

func Test_DoError(t *testing.T) {
	resp, err := Do("invalid url", "GET", nil)

	assert.Nil(t, resp)
	assert.Equal(t, errors.New("response failed with status code: 200 and body: "), err)
}
