package midllewares

import (
	"encoding/json"
	"net/http"

	"github.com/RomaBilka/parcel-tracking/api"
)

type internalErrorsReWriter struct {
	http.ResponseWriter
	status int
}

func (r *internalErrorsReWriter) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func (r *internalErrorsReWriter) Write(b []byte) (int, error) {
	if r.status == http.StatusInternalServerError {
		b, _ = json.Marshal(api.Error{Message: http.StatusText(http.StatusInternalServerError)})
		return r.ResponseWriter.Write(b)
	}
	return r.ResponseWriter.Write(b)
}

func RewriteInternalErrors(previous http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		rw := &internalErrorsReWriter{ResponseWriter: writer}
		previous.ServeHTTP(rw, request)
	})
}
