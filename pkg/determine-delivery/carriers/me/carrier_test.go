package me

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
		trackNumbers []string
		setupApiMock func(api *apiMock, trackNumbers []string)
		parcels      []carriers.Parcel
		err          error
	}{
		{
			name:         "Ok response",
			trackNumbers: []string{"12A12345", "12A12346"},
			setupApiMock: func(api *apiMock, trackNumbers []string) {
				res1 := &ShipmentsTrackResponse{
					ResultTable: []ShipmentTrackResponse{
						ShipmentTrackResponse{
							ShipmentNumberSender: trackNumbers[0],
							Country:              "UA",
							City:                 "City",
						},
					},
				}

				res2 := &ShipmentsTrackResponse{
					ResultTable: []ShipmentTrackResponse{
						ShipmentTrackResponse{
							ShipmentNumberSender: trackNumbers[1],
							Country:              "UA",
							City:                 "City",
						},
					},
				}

				api.
					On("TrackByTrackingNumber", trackNumbers[0]).
					Once().Return(res1, nil).
					On("TrackByTrackingNumber", trackNumbers[1]).
					Once().Return(res2, nil)
			},
			parcels: []carriers.Parcel{
				carriers.Parcel{TrackingNumber: "12A12345", Places: []carriers.Place{carriers.Place{Country: "UA", City: "City"}}},
				carriers.Parcel{TrackingNumber: "12A12346", Places: []carriers.Place{carriers.Place{Country: "UA", City: "City"}}},
			},
		},
		{
			name:         "Bad response",
			trackNumbers: []string{"12A12345"},
			setupApiMock: func(api *apiMock, trackNumbers []string) {
				api.On("TrackByTrackingNumber", trackNumbers[0]).Once().Return(nil, errors.New("bad request"))
			},
			err: errors.New("bad request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			api := &apiMock{}
			testCase.setupApiMock(api, testCase.trackNumbers)

			c := NewCarrier(api)
			parcels, err := c.Track(testCase.trackNumbers)

			assert.Equal(t, testCase.err, err)
			assert.ElementsMatch(t, testCase.parcels, parcels)
			api.AssertExpectations(t)
		})
	}
}

type apiMock struct {
	mock.Mock
}

func (m *apiMock) TrackByTrackingNumber(trackNumber string) (*ShipmentsTrackResponse, error) {
	arg := m.Called(trackNumber)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	}

	return arg.Get(0).(*ShipmentsTrackResponse), arg.Error(1)
}
