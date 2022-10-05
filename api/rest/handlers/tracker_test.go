package handlers

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleRest(t *testing.T) {
	testId := "testId-string"

	tests := []struct {
		name             string
		trackId          string
		method           string
		setupTrackerMock func(tracker *parcelTrackerMock)
		expResp          string
		expCode          int
	}{
		{
			name:             "invalid http method",
			setupTrackerMock: func(tracker *parcelTrackerMock) {},
			expCode:          http.StatusMethodNotAllowed,
			method:           http.MethodGet,
		},
		{
			name:             "failed: empty track id",
			method:           http.MethodPost,
			setupTrackerMock: func(tracker *parcelTrackerMock) {},
			expCode:          http.StatusBadRequest,
			expResp:          `{"message":"track_id cannot be empty"}`,
		},
		{
			name:    "failed to track parcel",
			trackId: testId,
			method:  http.MethodPost,
			setupTrackerMock: func(tracker *parcelTrackerMock) {
				tracker.On("TrackParcels", mock.Anything, []string{testId}).Once().
					Return(map[string]carriers.Parcel{}, assert.AnError)
			},
			expResp: `{"message":"` + assert.AnError.Error() + `"}`,
			expCode: http.StatusInternalServerError,
		},
		{
			name:    "success",
			method:  http.MethodPost,
			trackId: testId,
			setupTrackerMock: func(tracker *parcelTrackerMock) {
				data := map[string]carriers.Parcel{}
				data["testId-string"] = carriers.Parcel{
					TrackingNumber: "number",
					Places:         []carriers.Place{carriers.Place{Address: "address"}},
					Status:         "status",
				}

				tracker.On("TrackParcels", mock.Anything, []string{testId}).Once().
					Return(data, nil)
			},
			expResp: `{"testId-string":{"TrackingNumber":"number","Places":[{"Address":"address","Date":"0001-01-01T00:00:00Z"}],"Status":"status","DeliveryDate":"0001-01-01T00:00:00Z"}}` + "\n",
			expCode: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tracker := &parcelTrackerMock{}
			tc.setupTrackerMock(tracker)

			data := strings.NewReader("track_id=" + testId)

			req, err := http.NewRequest(tc.method, "", data)
			assert.NoError(t, err)

			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			rec := httptest.NewRecorder()

			Tracking(tracker, 10)(rec, req)

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

func (m *parcelTrackerMock) TrackParcels(ctx context.Context, parcelIds []string) (map[string]carriers.Parcel, error) {
	ret := m.Called(ctx, parcelIds)
	return ret.Get(0).(map[string]carriers.Parcel), ret.Error(1)
}
