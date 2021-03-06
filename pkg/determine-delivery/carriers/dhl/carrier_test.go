package dhl

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
		{name: "000", trackId: "0001111111111", ok: true},
		{name: "JVGL", trackId: "JVGL1111111111", ok: true},
		{name: "GM", trackId: "GM1111111111", ok: true},
		{name: "LX", trackId: "LX1111111111", ok: true},
		{name: "RX", trackId: "RX1111111111", ok: true},
		{name: "3S", trackId: "3S1111111111", ok: true},
		{name: "JJD", trackId: "JJD1111111111", ok: true},
		{name: "6 digits", trackId: "123456", ok: false},
		{name: "7 digits", trackId: "1234567", ok: true},
		{name: "9 digits", trackId: "123456789", ok: true},
		{name: "10 digits", trackId: "1234567890", ok: true},
		{name: "11 digits", trackId: "12345678901", ok: false},
		{name: "12 digits", trackId: "123456789012", ok: false},
		{name: "13 digits", trackId: "123456789013", ok: false},
		{name: "14 digits", trackId: "12345678901234", ok: true},
		{name: "15 digits", trackId: "123456789012345", ok: false},
		{name: "1234-12345", trackId: "1234-12345", ok: true},
		{name: "12345-12345", trackId: "12345-12345", ok: false},
		{name: "1234-123456", trackId: "1234-123456", ok: false},
		{name: "ABC-ABC-1234567", trackId: "ABC-ABC-1234567", ok: true},
		{name: "AB-AB-1234567", trackId: "AB-AB-1234567", ok: true},
		{name: "ABC-AB-1234567", trackId: "ABC-AB-1234567", ok: true},
		{name: "AB-ABC-1234567", trackId: "AB-ABC-1234567", ok: true},
		{name: "AB-ABC-123456", trackId: "AB-ABC-123456", ok: false},
		{name: "AB-A-1234567", trackId: "AB-A-1234567", ok: false},
		{name: "AB-ABCD-1234567", trackId: "AB-ABCD-1234567", ok: false},
		{name: "A-AB-1234567", trackId: "A-AB-1234567", ok: false},
		{name: "123-12345678", trackId: "123-12345678", ok: true},
		{name: "12-12345678", trackId: "12-12345678", ok: false},
		{name: "123-123456789", trackId: "123-123456789", ok: false},
		{name: "123-1234567", trackId: "123-1234567", ok: false},
		{name: "ABC123456", trackId: "ABC123456", ok: true},
		{name: "ABC12345", trackId: "ABC12345", ok: false},
		{name: "ABC1234567", trackId: "ABC1234567", ok: false},
		{name: "AB1234567", trackId: "AB1234567", ok: false},
		{name: "AB123456", trackId: "AB123456", ok: false},
		{name: "1AB123456", trackId: "1AB123456", ok: true},
		{name: "1AB12345", trackId: "1AB12345", ok: true},
		{name: "1AB1234", trackId: "1AB1234", ok: true},
		{name: "1AB1234567", trackId: "1AB1234567", ok: false},
		{name: "12AB1234", trackId: "12AB1234", ok: false},
		{name: "12A1234", trackId: "12A1234", ok: false},
		{name: "12A12345", trackId: "12A12345", ok: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			u := NewCarrier(NewApi("", ""))
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
			name: "Ok response",
			setupApiMock: func(api *apiMock, trackNumber string) {
				s := shipment{
					Id: trackNumber,
				}
				s.Status.Location = address{StreetAddress: "UA"}
				s.Status.Status = "Ok"

				res := &response{}
				res.Shipments = append(res.Shipments, s)

				api.On("TrackingDocument", trackNumber).Once().Return(res, nil)
			},
			parcels: []carriers.Parcel{{Address: "UA", Status: "Ok"}},
		},
		{
			name: "Bad response",
			setupApiMock: func(api *apiMock, trackNumber string) {
				api.On("TrackingDocument", trackNumber).Once().Return(nil, errors.New("bad request"))
			},
			err: errors.New("bad request"),
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

func (m *apiMock) TrackingDocument(trackNumber string) (*response, error) {
	arg := m.Called(trackNumber)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	}

	return arg.Get(0).(*response), arg.Error(1)
}
