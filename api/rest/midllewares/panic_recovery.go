package midllewares

import (
	"net/http"
)

type unknownArgumentsLogger interface {
	Error(args ...interface{})
}

func PanicRecovery(log unknownArgumentsLogger) Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					// trigger rewrite internal errors middleware, to write correct error message
					_, _ = rw.Write(nil)

					log.Error(r)
				}
			}()
			handler.ServeHTTP(rw, request)
		})
	}
}
