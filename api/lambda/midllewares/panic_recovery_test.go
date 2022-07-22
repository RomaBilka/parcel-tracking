package midllewares

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/RomaBilka/parcel-tracking/api/lambda/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPanicRecovery(t *testing.T) {
	testCases := []struct {
		name   string
		prev   handlers.Handler
		logger func(m *unknownArgsLoggerMock)

		expResp events.APIGatewayProxyResponse
		expErr  error
	}{
		{
			name: "recover",
			prev: func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
				panic("test panic")
			},
			logger: func(m *unknownArgsLoggerMock) {
				m.On("Error", "test panic")
			},
			expErr: errors.New(""),
		},
		{
			name:   "do not recover",
			logger: func(m *unknownArgsLoggerMock) {},
			prev: func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
				return events.APIGatewayProxyResponse{Body: `{"message":"ok"}`, StatusCode: http.StatusOK}, nil
			},
			expResp: events.APIGatewayProxyResponse{Body: `{"message":"ok"}`, StatusCode: http.StatusOK},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := &unknownArgsLoggerMock{}
			tc.logger(l)

			resp, err := PanicRecovery(l)(tc.prev)(context.Background(), events.APIGatewayProxyRequest{})
			assert.Equal(t, tc.expResp, resp)
			assert.Equal(t, tc.expErr, err)

			l.AssertExpectations(t)
		})
	}
}

type unknownArgsLoggerMock struct {
	mock.Mock
}

func (u *unknownArgsLoggerMock) Error(args ...interface{}) {
	u.Called(args...)
}
