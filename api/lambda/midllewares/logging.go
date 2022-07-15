package midllewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RomaBilka/parcel-tracking/api/lambda/handlers"
	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"
)

type Middleware func(handler handlers.Handler) handlers.Handler

type logger interface {
	Error(msg string, fields ...zap.Field)
}

func Logging(logger logger) Middleware {
	return func(previous handlers.Handler) handlers.Handler {
		return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
			resp, err := previous(ctx, request)

			var msg string
			if resp.StatusCode == http.StatusInternalServerError {
				msg += fmt.Sprintf(", body: %s", resp.Body)
			}
			if err != nil {
				msg += fmt.Sprintf(", error: %s", err)
			}
			if msg != "" {
				logger.Error("internal server error"+msg, zap.Int("StatusCode", http.StatusInternalServerError))
			}

			return resp, err
		}
	}
}
