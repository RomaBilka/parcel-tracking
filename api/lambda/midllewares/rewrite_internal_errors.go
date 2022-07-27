package midllewares

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/RomaBilka/parcel-tracking/api"
	"github.com/RomaBilka/parcel-tracking/api/lambda/handlers"
	"github.com/aws/aws-lambda-go/events"
)

// RewriteInternalErrors - to avoid logging at the lambda side, we need to not return an error, otherwise we will
// have logs duplications.
func RewriteInternalErrors(previous handlers.Handler) handlers.Handler {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		resp, err := previous(ctx, request)
		if err != nil {
			body, _ := json.Marshal(api.Error{Message: http.StatusText(http.StatusInternalServerError)})
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       string(body),
			}, nil
		}
		return resp, nil
	}
}
