package midllewares

import (
	"context"
	"encoding/json"
	"time"

	"github.com/RomaBilka/parcel-tracking/api"
	"github.com/RomaBilka/parcel-tracking/api/lambda/handlers"
	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"
)

type Middleware func(handler handlers.Handler) handlers.Handler

type logger interface {
	Error(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
}

func Logging(logger logger) Middleware {
	return func(previous handlers.Handler) handlers.Handler {
		return func(ctx context.Context, request events.APIGatewayProxyRequest) (resp events.APIGatewayProxyResponse, respErr error) {
			tn := time.Now()
			resp, respErr = previous(ctx, request)
			if respErr != nil {
				logger.Error(respErr.Error(), zap.Duration("duration", time.Since(tn)))
				return resp, respErr
			}

			fields := []zap.Field{
				zap.Int("statusCode", resp.StatusCode),
				zap.Duration("duration", time.Since(tn)),
			}

			if api.IsWarn(resp.StatusCode) || api.IsErr(resp.StatusCode) {
				apiErr := api.Error{}
				err := json.Unmarshal([]byte(resp.Body), &apiErr)
				if err != nil {
					fields = append(fields, zap.Error(err), zap.String("response", resp.Body))
					logger.Error("failed to unmarshall error", fields...)
					return
				}

				if api.IsWarn(resp.StatusCode) {
					logger.Warn(apiErr.Message, fields...)
				}
				if api.IsErr(resp.StatusCode) {
					logger.Error(apiErr.Message, fields...)
				}
			}
			return
		}
	}
}
