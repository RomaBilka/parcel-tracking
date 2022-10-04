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
	trackId := "NPMD000000001456NPI"

	testCases := []struct {
		name    string
		apiMock func(m *apiMock)
		expResp []carriers.Parcel
		expErr  error
	}{
		{
			name: "failed to track by number",
			apiMock: func(m *apiMock) {
				m.On("TrackByTrackingNumber", trackId).Once().
					Return(nil, assert.AnError)
			},
			expErr: assert.AnError,
		},
		{
			name: "successful track by number",
			apiMock: func(m *apiMock) {
				m.On("TrackByTrackingNumber", trackId).Once().
					Return(&TrackingDocumentResponse{
						WaybillNumber: trackId,
						State:         "Delivered",
					}, nil)
			},
			expResp: []carriers.Parcel{{
				TrackingNumber: trackId,
				Status:         "Delivered",
				Places:         []carriers.Place{},
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := &apiMock{}
			tc.apiMock(m)

			n := NewCarrier(m)
			resp, err := n.Track([]string{trackId})

			assert.Equal(t, tc.expResp, resp)
			assert.Equal(t, tc.expErr, err)

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
