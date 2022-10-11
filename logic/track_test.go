package logic

import (
	"context"
	"testing"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestParcelsTracker_TrackParcel(t *testing.T) {
	testId := "testId-string"

	tests := []struct {
		name       string
		setupMocks func(dm *detectorMock, cm *carrierMock)
		expParcel  map[string]carriers.Parcel
		expErr     error
	}{
		{
			name: "failed to detect parcel",
			setupMocks: func(dm *detectorMock, cm *carrierMock) {
				dm.On("Detect", testId).Once().Return(nil, assert.AnError)
			},
			expErr: assert.AnError,
		},
		{
			name: "failed to track parcel",
			setupMocks: func(dm *detectorMock, cm *carrierMock) {
				dm.On("Detect", testId).Once().Return(cm, nil)
				cm.On("Track", []string{testId}).Once().Return([]carriers.Parcel{}, nil)
			},
			expParcel: make(map[string]carriers.Parcel),
		},
		{
			name: "success",
			setupMocks: func(dm *detectorMock, cm *carrierMock) {
				dm.On("Detect", testId).Once().Return(cm, nil)
				cm.On("Track", []string{testId}).Once().
					Return([]carriers.Parcel{{TrackingNumber: "123", Places: []carriers.Place{carriers.Place{Address: "223"}}, Status: "323"}}, nil)
			},
			expParcel: map[string]carriers.Parcel{"123": carriers.Parcel{TrackingNumber: "123", Places: []carriers.Place{carriers.Place{Address: "223"}}, Status: "323"}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dm := &detectorMock{}
			cm := &carrierMock{}

			tc.setupMocks(dm, cm)

			tr := NewParcelsTracker(dm)
			gotParcel, gotErr := tr.TrackParcels(context.Background(), []string{testId})

			assert.Equal(t, tc.expParcel, gotParcel)
			assert.Equal(t, tc.expErr, gotErr)
		})
	}
}

type detectorMock struct {
	mock.Mock
}

func (m *detectorMock) Detect(trackId string) (carriers.Carrier, error) {
	ret := m.Called(trackId)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).(carriers.Carrier), ret.Error(1)
}

type carrierMock struct {
	mock.Mock
}

func (m *carrierMock) Track(trackId []string) ([]carriers.Parcel, error) {
	ret := m.Called(trackId)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).([]carriers.Parcel), ret.Error(1)
}

func (m *carrierMock) Detect(trackId string) bool {
	return m.Called(trackId).Bool(0)
}
