package fedex

import (
	"errors"
	"testing"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCarrier_Detect(t *testing.T) {
	testCases := []struct {
		name    string
		trackId string
		ok      bool
	}{
		{name: "12 true", trackId: "123456789012", ok: true},
		{name: "15 true", trackId: "123456789012345", ok: true},
		{name: "20 true", trackId: "12345678901234567890", ok: true},
		{name: "22 true", trackId: "1234567890123456789012", ok: true},
		{name: "21 false", trackId: "123456789012345678901", ok: false},
		{name: "10 false", trackId: "1234567890", ok: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			fedex := NewCarrier(NewApi("", "", "", ""))
			ok := fedex.Detect(testCase.trackId)
			assert.Equal(t, testCase.ok, ok)
		})
	}
}

func TestCarrier_Track(t *testing.T) {
	testCases := []struct {
		name         string
		trackNumber  string
		setupApiMock func(api *apiMock, requestData TrackingRequest)
		parcels      []carriers.Parcel
		err          error
	}{
		{
			name:        "Ok response",
			trackNumber: "123456789012",
			setupApiMock: func(api *apiMock, requestData TrackingRequest) {
				response := &TrackingResponse{
					Output: Output{
						CompleteTrackResults: []CompleteTrackResult{
							CompleteTrackResult{
								TrackResults: []TrackResult{
									TrackResult{
										LatestStatusDetail: LatestStatusDetail{
											ScanLocation: ScanLocation{
												CountryName: "Country Name",
												City:        "City",
											},
											StatusByLocale: "Ok",
										},
									},
								},
							},
						},
					},
				}

				api.On("TrackByTrackingNumber", requestData).Once().Return(response, nil)
			},
			parcels: []carriers.Parcel{{Address: "Country Name City", Status: "Ok"}},
		},
		{
			name: "Bad response",
			setupApiMock: func(api *apiMock, requestData TrackingRequest) {
				api.On("TrackByTrackingNumber", requestData).Once().Return(nil, errors.New("bad request"))
			},
			err: errors.New("bad request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			api := &apiMock{}

			trackingInfo := TrackingInfo{
				TrackingNumberInfo: TrackingNumberInfo{
					TrackingNumber: testCase.trackNumber,
				},
			}

			requestData := TrackingRequest{IncludeDetailedScans: true}
			requestData.TrackingInfo = append(requestData.TrackingInfo, trackingInfo)

			testCase.setupApiMock(api, requestData)

			c := NewCarrier(api)
			parcels, err := c.Track(testCase.trackNumber)

			assert.Equal(t, testCase.err, err)
			assert.Equal(t, testCase.parcels, parcels)
			api.AssertExpectations(t)
		})
	}
}

type apiMock struct {
	mock.Mock
}

func (m *apiMock) TrackByTrackingNumber(trackingRequest TrackingRequest) (*TrackingResponse, error) {
	arg := m.Called(trackingRequest)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	}

	return arg.Get(0).(*TrackingResponse), arg.Error(1)
}
