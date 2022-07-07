package handlers

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleLambdaEvent(t *testing.T) {
	testID := "testID-string"

	tests := []struct {
		name             string
		trackID          string
		method           string
		setupTrackerMock func(tracker *parcelTrackerMock)
		expResp          string
		expCode          int
	}{
		{
			name:             "invalid http method",
			setupTrackerMock: func(tracker *parcelTrackerMock) {},
			expCode:          http.StatusMethodNotAllowed,
		},
		{
			name:             "failed: empty track id",
			method:           http.MethodGet,
			setupTrackerMock: func(tracker *parcelTrackerMock) {},
			expCode:          http.StatusBadRequest,
			expResp:          `{"message":"track_id cannot be empty"}`,
		},
		{
			name:    "failed to track parcel",
			trackID: testID,
			method:  http.MethodGet,
			setupTrackerMock: func(tracker *parcelTrackerMock) {
				tracker.On("TrackParcel", mock.Anything, testID).Once().
					Return(carriers.Parcel{}, assert.AnError)
			},
			expResp: `{"message":"` + assert.AnError.Error() + `"}`,
			expCode: http.StatusBadRequest,
		},
		{
			name:    "success",
			method:  http.MethodGet,
			trackID: testID,
			setupTrackerMock: func(tracker *parcelTrackerMock) {
				tracker.On("TrackParcel", mock.Anything, testID).Once().
					Return(carriers.Parcel{Number: "number", Address: "address", Status: "status"}, nil)
			},
			expResp: `{"Number":"number","Address":"address","Status":"status"}` + "\n",
			expCode: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tracker := &parcelTrackerMock{}
			tc.setupTrackerMock(tracker)

			req := &http.Request{
				Method: tc.method,
				URL:    &url.URL{RawQuery: "track_id=" + tc.trackID},
			}
			rec := httptest.NewRecorder()

			Tracking(tracker)(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			gotResp, err := ioutil.ReadAll(res.Body)

			assert.NoError(t, err)
			assert.Equal(t, tc.expResp, string(gotResp))
			assert.Equal(t, tc.expCode, rec.Code)
		})
	}
}

type parcelTrackerMock struct {
	mock.Mock
}

func (m *parcelTrackerMock) TrackParcel(ctx context.Context, parcelID string) (carriers.Parcel, error) {
	ret := m.Called(ctx, parcelID)
	return ret.Get(0).(carriers.Parcel), ret.Error(1)
}
