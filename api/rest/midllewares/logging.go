package midllewares

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/RomaBilka/parcel-tracking/api"
	"go.uber.org/zap"
)

type responseErrorsWriter struct {
	http.ResponseWriter

	status int
	error  []byte
}

func (r *responseErrorsWriter) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func (r *responseErrorsWriter) Write(b []byte) (int, error) {
	if api.IsWarn(r.status) || api.IsErr(r.status) {
		r.error = b
	}
	return r.ResponseWriter.Write(b)
}

type (
	Middleware func(handler http.Handler) http.Handler

	logger interface {
		Error(msg string, fields ...zap.Field)
		Warn(msg string, fields ...zap.Field)
	}
)

func Logging(logger logger) Middleware {
	return func(prev http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			tn := time.Now()

			rw := &responseErrorsWriter{ResponseWriter: writer}
			prev.ServeHTTP(rw, request)

			if len(rw.error) == 0 {
				return
			}

			fields := []zap.Field{
				zap.Int("statusCode", rw.status),
				zap.Duration("duration", time.Since(tn)),
			}

			apiErr := api.Error{}
			if err := json.Unmarshal(rw.error, &apiErr); err != nil {
				fields = append(fields, zap.Error(err), zap.String("response", string(rw.error)))

				logger.Error("failed to unmarshall error", fields...)
				return
			}

			if api.IsWarn(rw.status) {
				logger.Warn(apiErr.Message, fields...)
			}
			if api.IsErr(rw.status) {
				logger.Error(apiErr.Message, fields...)
			}
		})
	}
}
