package midllewares

import (
	"context"
	"net/http"
	"testing"

	"github.com/RomaBilka/parcel-tracking/api/lambda/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestRewriteInternalErrors(t *testing.T) {
	testCases := []struct {
		name    string
		prev    handlers.Handler
		expResp events.APIGatewayProxyResponse
	}{
		{
			name: "ignore, if no errors",
			prev: func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
				return events.APIGatewayProxyResponse{Body: "ok", StatusCode: http.StatusOK}, nil
			},
			expResp: events.APIGatewayProxyResponse{Body: "ok", StatusCode: http.StatusOK},
		},
		{
			name: "rewrite if error",
			prev: func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
				return events.APIGatewayProxyResponse{Body: "ok", StatusCode: http.StatusOK}, assert.AnError
			},
			expResp: events.APIGatewayProxyResponse{Body: "internal server error", StatusCode: http.StatusInternalServerError},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := RewriteInternalErrors(tc.prev)(context.Background(), events.APIGatewayProxyRequest{})
			assert.Equal(t, tc.expResp, resp)
			assert.Nil(t, err)
		})
	}
}
