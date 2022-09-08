package handlers

import (
	"context"
	"net/http"
	"testing"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleLambdaEvent(t *testing.T) {
	testId := "testId-string"

	tests := []struct {
		name             string
		trackId          string
		setupTrackerMock func(tracker *parcelTrackerMock)
		expResp          events.APIGatewayProxyResponse
		expErr           error
	}{
		{
			name:             "failed: empty track id",
			setupTrackerMock: func(tracker *parcelTrackerMock) {},
			expResp: events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       `{"message":"track_id cannot be empty"}`,
			},
		},
		{
			name:    "failed to track parcel",
			trackId: testId,
			setupTrackerMock: func(tracker *parcelTrackerMock) {
				tracker.On("TrackParcel", mock.Anything, testId).Once().
					Return(carriers.Parcel{}, assert.AnError)
			},
			expResp: events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       `{"message":"` + assert.AnError.Error() + `"}`,
			},
		},
		{
			name:    "success",
			trackId: testId,
			setupTrackerMock: func(tracker *parcelTrackerMock) {
				tracker.On("TrackParcel", mock.Anything, testId).Once().
					Return(carriers.Parcel{TrackingNumber: "number", Places: []carriers.Place{carriers.Place{Address: "address"}}, Status: "status"}, nil)
			},
			expResp: events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Body:       `{"TrackingNumber":"number","Places":[{"Address":"address","Date":"0001-01-01T00:00:00Z"}],"Status":"status","DeliveryDate":"0001-01-01T00:00:00Z"}`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tracker := &parcelTrackerMock{}
			tc.setupTrackerMock(tracker)

			req := events.APIGatewayProxyRequest{QueryStringParameters: map[string]string{"track_id": tc.trackId}}
			gotResp, gotErr := Tracking(tracker)(context.Background(), req)

			assert.Equal(t, tc.expResp, gotResp)
			assert.Equal(t, tc.expErr, gotErr)
		})
	}
}

type parcelTrackerMock struct {
	mock.Mock
}

func (m *parcelTrackerMock) TrackParcel(ctx context.Context, parcelId string) (carriers.Parcel, error) {
	ret := m.Called(ctx, parcelId)
	return ret.Get(0).(carriers.Parcel), ret.Error(1)
}
