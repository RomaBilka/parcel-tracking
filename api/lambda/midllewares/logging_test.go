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
			name: "log error on internal error",
			prev: func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
				return events.APIGatewayProxyResponse{Body: `{"message": "body error"}`, StatusCode: http.StatusInternalServerError}, nil
			},
			logger: func(m *loggerMock) {
				m.On("Error", "body error",
					zap.Int("statusCode", http.StatusInternalServerError), mock.Anything)
			},
		},
		{
			name: "log error from error",
			prev: func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
				return events.APIGatewayProxyResponse{}, assert.AnError
			},
			logger: func(m *loggerMock) {
				m.On("Error", assert.AnError.Error(), mock.Anything)
			},
		},
		{
			name: "log warn on bad request",
			prev: func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
				return events.APIGatewayProxyResponse{Body: `{"message": "body error"}`, StatusCode: http.StatusBadRequest}, nil
			},
			logger: func(m *loggerMock) {
				m.On("Warn", "body error",
					zap.Int("statusCode", http.StatusBadRequest), mock.Anything)
			},
		},
		{
			name: "failed to unmarshall",
			prev: func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
				return events.APIGatewayProxyResponse{Body: "body error", StatusCode: http.StatusNotFound}, nil
			},
			logger: func(m *loggerMock) {
				m.On("Error", "failed to unmarshall error",
					zap.Int("statusCode", http.StatusNotFound), mock.Anything, mock.Anything,
					zap.String("response", "body error"),
				)
			},
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

func (l *loggerMock) Warn(msg string, fields ...zap.Field) {
	args := []interface{}{msg}
	for _, field := range fields {
		args = append(args, field)
	}
	l.Called(args...)
}
