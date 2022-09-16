package ukrposhta

import (
	"fmt"
	"testing"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

func TestCarrier_Detect(t *testing.T) {
	testCases := []struct {
		name    string
		trackId string
		ok      bool
	}{
		{name: "Numeric only 13 true", trackId: "0500128254610", ok: true},
		{name: "Numeric only 14 false", trackId: "10500128254610", ok: false},
		{name: "Numeric only 12 false", trackId: "105001282546", ok: false},
		{name: "RA067022878UA", trackId: "RA067022878UA", ok: true},
		{name: "UU123456789CN", trackId: "UU123456789CN", ok: true},
		{name: "LO123456789FR", trackId: "LO123456789FR", ok: true},
		{name: "LOQ123456789FR", trackId: "LOQ123456789FR", ok: false},
		{name: "UU1234567890CN", trackId: "UU1234567890CN", ok: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			n := NewCarrier(NewApi("", ""))
			ok := n.Detect(testCase.trackId)
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
			trackNumber: "RA067022878UA",
			setupApiMock: func(api *apiMock, trackNumber string) {
				res := &Response{
					Found: map[string][]Item{
						trackNumber: {
							Item{
								Barcode: trackNumber,
								Country: "UA",
								Name:    "Lviv",
							},
						},
					},
				}

				api.On("TrackByTrackingNumber", []string{trackNumber}).Once().Return(res, nil)
			},
			parcels: []carriers.Parcel{{TrackingNumber: "RA067022878UA", Places: []carriers.Place{carriers.Place{County: "UA", Address: "Lviv"}}}},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			api := &apiMock{}
			testCase.setupApiMock(api, testCase.trackNumber)

			c := NewCarrier(api)
			parcels, err := c.Track(testCase.trackNumber)
			fmt.Println(parcels)
			assert.Equal(t, testCase.err, err)
			assert.Equal(t, testCase.parcels, parcels)
			api.AssertExpectations(t)
		})
	}
}

type apiMock struct {
	mock.Mock
}

func (m *apiMock) TrackByTrackingNumber(barcodes []string) (*Response, error) {
	arg := m.Called(barcodes)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	}

	return arg.Get(0).(*Response), arg.Error(1)
}
