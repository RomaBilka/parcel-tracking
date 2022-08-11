package ups

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
		{name: "8 true", trackId: "800000000000000001", ok: true},
		{name: "9 true", trackId: "900000000000000001", ok: true},
		{name: "1Z false", trackId: "1z000000000000000000000000", ok: false},
		{name: "8 false", trackId: "80000000000000000000", ok: false},
		{name: "9 false", trackId: "90000000000000000", ok: false},
		{name: "unknown", trackId: "01234567891011", ok: false},
		{name: "9 false", trackId: "90000000000000000", ok: false},
		{name: "1Z12345E6605272234", trackId: "1Z12345E6605272234", ok: true},
		{name: "1Z123456E66052722", trackId: "1Z123456E66052722", ok: false},
		{name: "1Z123456E6605272234", trackId: "1Z123456E6605272234", ok: true},
		{name: "1Z123456E660527223", trackId: "1Z123456E660527223", ok: true},
		{name: "1ZWX0692YP40636269", trackId: "1ZWX0692YP40636269", ok: true},
		{name: "1ZWX0692YP406362690", trackId: "1ZWX0692YP406362690", ok: false},
		{name: "1ZWX0692Y40636269", trackId: "1ZWX0692YP406362690", ok: false},
		{name: "123456789", trackId: "123456789", ok: true},
		{name: "1234567890", trackId: "1234567890", ok: true},
		{name: "12345678901", trackId: "12345678901", ok: false},
		{name: "21 true", trackId: "123456789012345678901", ok: false},
		{name: "22 true", trackId: "1234567890123456789012", ok: true},
		{name: "23 true", trackId: "12345678901234567890123", ok: false},
		{name: "cgish000116630", trackId: "cgish000116630", ok: true},
		{name: "cgish0001166301", trackId: "cgish0001166301", ok: false},
		{name: "cgish00011663", trackId: "cgish00011663", ok: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			u := NewCarrier(NewApi("", "", "", ""))
			ok := u.Detect(testCase.trackId)
			assert.Equal(t, testCase.ok, ok)
		})
	}
}

func TestCarrier_Track(t *testing.T) {
	testCases := []struct {
		name         string
		trackNumber  string
		setupApiMock func(api *apiMock, trackNumber string)
		parcels      []carriers.Parcel
		err          error
	}{
		{
			name:        "Ok response",
			trackNumber: "1Z12345E0291980793",
			setupApiMock: func(api *apiMock, trackNumber string) {
				res := &TrackResponse{
					Shipment: Shipment{
						ShipmentIdentificationNumber: trackNumber,
						Package: Package{
							Activity: []Activity{{}},
						},
					},
				}
				api.On("TrackByTrackingNumber", trackNumber).Once().Return(res, nil)
			},
			parcels: []carriers.Parcel{{Number: "1Z12345E0291980793"}},
		},
		{
			name: "Bad response",
			setupApiMock: func(api *apiMock, trackNumber string) {
				api.On("TrackByTrackingNumber", trackNumber).Once().Return(nil, errors.New("Invalid tracking number"))
			},
			err: errors.New("Invalid tracking number"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			api := &apiMock{}
			testCase.setupApiMock(api, testCase.trackNumber)

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

func (m *apiMock) TrackByTrackingNumber(trackNumber string) (*TrackResponse, error) {
	arg := m.Called(trackNumber)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	}

	return arg.Get(0).(*TrackResponse), arg.Error(1)
}
