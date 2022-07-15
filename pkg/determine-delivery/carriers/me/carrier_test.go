package me

import (
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
		{name: "CV true", trackId: "CV999999999ZZ", ok: true},
		{name: "cv true", trackId: "cv999999999zz", ok: true},
		{name: "CV_1 false", trackId: "CV999999999ZZz", ok: false},
		{name: "CV_2 false", trackId: "CV9999999999ZZ", ok: false},
		{name: "MYCV true", trackId: "MYCV999999999ZZ", ok: true},
		{name: "unknown", trackId: "ZZZZ999999999ZZ", ok: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			m := NewCarrier(NewApi("", "", "", ""))
			ok := m.Detect(testCase.trackId)
			assert.Equal(t, testCase.ok, ok)
		})
	}
}

func TestCarrier_Track(t *testing.T) {
	testCases := []struct {
		name         string
		trackNumber  string
		setupApiMock func(api *apiMock, trackNumber string)
		parcel       carriers.Parcel
	}{
		{
			name: "Ok response",
			setupApiMock: func(api *apiMock, trackNumber string) {
				shipment := ShipmentTrackResponse{
					ShipmentNumberSender: trackNumber,
					CountryDel:           "UA",
					ActionMessages_UA:    "Action Messages",
					DetailMessages_UA:    "Detail Messages",
				}

				res := &ShipmentsTrackResponse{}
				res.ResultTable = append(res.ResultTable, shipment)

				api.On("ShipmentsTrack", trackNumber).Once().Return(res, nil)
			},
			parcel: carriers.Parcel{Address: "UA", Status: "Action Messages Detail Messages"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			api := &apiMock{}
			testCase.setupApiMock(api, testCase.trackNumber)

			c := NewCarrier(api)
			parcels, err := c.Track(testCase.trackNumber)

			assert.NoError(t, err)
			assert.Equal(t, testCase.parcel, parcels[0])
			api.AssertExpectations(t)
		})
	}
}

type apiMock struct {
	mock.Mock
}

func (m *apiMock) ShipmentsTrack(trackNumber string) (*ShipmentsTrackResponse, error) {
	arg := m.Called(trackNumber)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	}

	return arg.Get(0).(*ShipmentsTrackResponse), arg.Error(1)
}
