package np_shopping

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
		{name: "starts with NPI", trackId: "NPI99999999999999", ok: true},
		{name: "ends with NPI", trackId: "NPMD000000001456NPI", ok: true},
		{name: "ends with NPG", trackId: "OG0212001001NPG", ok: true},
		{name: "starts with NPG true", trackId: "NPGMD000000001456", ok: true},
		{name: "invalid starts/ends with char", trackId: "cv999999999zz", ok: false},
		{name: "invalid starts/ends with number", trackId: "9999999", ok: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			n := Carrier{}
			ok := n.Detect(testCase.trackId)
			assert.Equal(t, testCase.ok, ok)
		})
	}
}

func TestCarrier_Track(t *testing.T) {
	testCases := []struct {
		name     string
		apiMock  func(m *apiMock, trackIds []string)
		expResp  []carriers.Parcel
		expErr   error
		trackIds []string
	}{
		{
			name: "failed to track by number",
			apiMock: func(m *apiMock, trackIds []string) {
				m.On("TrackByTrackingNumber", trackIds[0]).Once().
					Return(nil, assert.AnError)
			},
			expErr:   assert.AnError,
			trackIds: []string{"NPMD000000001455NPI"},
		},
		{
			name:     "successful track by number",
			trackIds: []string{"NPMD000000001455NPI", "NPMD000000001456NPI"},
			apiMock: func(m *apiMock, trackIds []string) {
				m.On("TrackByTrackingNumber", trackIds[0]).Once().
					Return(&TrackingDocumentResponse{
						WaybillNumber: trackIds[0],
						State:         "Delivered",
					}, nil).
					On("TrackByTrackingNumber", trackIds[1]).Once().
					Return(&TrackingDocumentResponse{
						WaybillNumber: trackIds[1],
						State:         "Delivered",
					}, nil)
			},
			expResp: []carriers.Parcel{
				{
					TrackingNumber: "NPMD000000001455NPI",
					Status:         "Delivered",
					Places:         []carriers.Place{},
				},
				{
					TrackingNumber: "NPMD000000001456NPI",
					Status:         "Delivered",
					Places:         []carriers.Place{},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			m := &apiMock{}
			testCase.apiMock(m, testCase.trackIds)

			n := NewCarrier(m)
			parcels, err := n.Track(testCase.trackIds)

			assert.Equal(t, testCase.expErr, err)
			assert.ElementsMatch(t, testCase.expResp, parcels)
			m.AssertExpectations(t)
		})
	}
}

type apiMock struct {
	mock.Mock
}

func (m *apiMock) TrackByTrackingNumber(trackNumber string) (*TrackingDocumentResponse, error) {
	arg := m.Called(trackNumber)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	}

	return arg.Get(0).(*TrackingDocumentResponse), arg.Error(1)
}
