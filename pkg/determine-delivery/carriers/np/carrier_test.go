package np

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
		{name: "59 true", trackId: "59000000000001", ok: true},
		{name: "20 true", trackId: "20000000000001", ok: true},
		{name: "1 true", trackId: "10000000000000", ok: true},
		{name: "59 false", trackId: "5900000000000", ok: false},
		{name: "20 false", trackId: "2000000000000", ok: false},
		{name: "1 false", trackId: "1000000000000", ok: false},
		{name: "unknown", trackId: "01234567891011", ok: false},
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
			name: "Ok response",
			setupApiMock: func(api *apiMock, trackNumber string) {
				trackingDocument := TrackingDocument{
					DocumentNumber: trackNumber,
				}
				methodProperties := TrackingDocuments{CheckWeightMethod: "3"}
				methodProperties.Documents = append(methodProperties.Documents, trackingDocument)

				document := TrackingDocumentResponse{
					Number:             trackNumber,
					CityRecipient:      "City Recipient",
					CitySender:         "City Sender",
					WarehouseRecipient: "Warehouse Recipient",
					WarehouseSender:    "Warehouse Sender",
					Status:             "Ok",
				}

				res := &TrackingDocumentsResponse{}
				res.Data = append(res.Data, document)

				api.On("TrackByTrackingNumber", methodProperties).Once().Return(res, nil)
			},
			parcels: []carriers.Parcel{{Places: []carriers.Place{carriers.Place{City: "City Sender", Address: "Warehouse Sender"}, carriers.Place{City: "City Recipient", Address: "Warehouse Recipient"}}, Status: "Ok"}},
		},
		{
			name: "Bad response",
			setupApiMock: func(api *apiMock, trackNumber string) {
				trackingDocument := TrackingDocument{
					DocumentNumber: trackNumber,
				}
				methodProperties := TrackingDocuments{CheckWeightMethod: "3"}
				methodProperties.Documents = append(methodProperties.Documents, trackingDocument)

				api.On("TrackByTrackingNumber", methodProperties).Once().Return(nil, errors.New("bad request"))
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

func (m *apiMock) TrackByTrackingNumber(methodProperties TrackingDocuments) (*TrackingDocumentsResponse, error) {
	arg := m.Called(methodProperties)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	}

	return arg.Get(0).(*TrackingDocumentsResponse), arg.Error(1)
}
