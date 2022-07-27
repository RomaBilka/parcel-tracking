package midllewares

import (
	"context"
	"errors"

	"github.com/RomaBilka/parcel-tracking/api/lambda/handlers"
	"github.com/aws/aws-lambda-go/events"
)

type unknownArgumentsLogger interface {
	Error(args ...interface{})
}

func PanicRecovery(log unknownArgumentsLogger) Middleware {
	return func(previous handlers.Handler) handlers.Handler {
		return func(ctx context.Context, request events.APIGatewayProxyRequest) (resp events.APIGatewayProxyResponse, respErr error) {
			defer func() {
				if r := recover(); r != nil {
					// trigger rewrite internal errors middleware, to write correct error message
					respErr = errors.New("")
					log.Error(r)
				}
			}()

			return previous(ctx, request)
		}
	}
}
