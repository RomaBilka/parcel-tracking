package midllewares

import (
	"context"
	"net/http"
	"testing"

	"github.com/RomaBilka/parcel-tracking/api/lambda/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestEnableLogging(t *testing.T) {
	testCases := []struct {
		name   string
		prev   handlers.Handler
		logger func(m *loggerMock)
	}{
		{
			name: "log message from body",
			prev: func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
				return events.APIGatewayProxyResponse{Body: "body error", StatusCode: http.StatusInternalServerError}, nil
			},
			logger: func(m *loggerMock) {
				m.On("Error", "internal server error, body: body error",
					zap.Int("StatusCode", http.StatusInternalServerError))
			},
		},
		{
			name: "log message from error",
			prev: func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
				return events.APIGatewayProxyResponse{}, assert.AnError
			},
			logger: func(m *loggerMock) {
				m.On("Error", "internal server error, error: assert.AnError general error for testing", zap.Int("StatusCode", http.StatusInternalServerError))
			},
		},
		{
			name: "log message from body and error",
			prev: func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
				return events.APIGatewayProxyResponse{Body: "body error", StatusCode: http.StatusInternalServerError}, assert.AnError
			},
			logger: func(m *loggerMock) {
				m.On("Error", "internal server error, body: body error, error: assert.AnError general error for testing",
					zap.Int("StatusCode", http.StatusInternalServerError))
			},
		},
		{
			name: "do not log if status not found",
			prev: func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
				return events.APIGatewayProxyResponse{Body: "body", StatusCode: http.StatusNotFound}, nil
			},
			logger: func(m *loggerMock) {},
		},
		{
			name: "do not log if status success",
			prev: func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
				return events.APIGatewayProxyResponse{Body: "body", StatusCode: http.StatusOK}, nil
			},
			logger: func(m *loggerMock) {},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := &loggerMock{}
			tc.logger(l)

			emptyCtx := context.Background()
			emptyReq := events.APIGatewayProxyRequest{}

			expResp, expErr := tc.prev(emptyCtx, emptyReq)

			resp, err := Logging(l)(tc.prev)(emptyCtx, emptyReq)
			assert.Equal(t, expResp, resp)
			assert.Equal(t, expErr, err)

			l.AssertExpectations(t)
		})
	}
}

type loggerMock struct {
	mock.Mock
}

func (l *loggerMock) Error(msg string, fields ...zap.Field) {
	args := []interface{}{msg}
	for _, field := range fields {
		args = append(args, field)
	}
	l.Called(args...)
}
