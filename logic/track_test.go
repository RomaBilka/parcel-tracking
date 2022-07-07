package logic

import (
	"context"
	"errors"
	"testing"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestParcelsTracker_TrackParcel(t *testing.T) {
	testID := "testID-string"

	tests := []struct {
		name       string
		setupMocks func(dm *detectorMock, cm *carrierMock)
		expParcel  carriers.Parcel
		expErr     error
	}{
		{
			name: "failed to detect parcel",
			setupMocks: func(dm *detectorMock, cm *carrierMock) {
				dm.On("Detect", testID).Once().Return(nil, assert.AnError)
			},
			expErr: assert.AnError,
		},
		{
			name: "failed to track parcel",
			setupMocks: func(dm *detectorMock, cm *carrierMock) {
				dm.On("Detect", testID).Once().Return(cm, nil)
				cm.On("Track", testID).Once().Return(nil, assert.AnError)
			},
			expErr: assert.AnError,
		},
		{
			name: "invalid number of parcels, more than 1",
			setupMocks: func(dm *detectorMock, cm *carrierMock) {
				dm.On("Detect", testID).Once().Return(cm, nil)
				cm.On("Track", testID).Once().Return([]carriers.Parcel{{}, {}}, nil)
			},
			expErr: errors.New("invalid number of parcels, expected 1 - got 2"),
		},
		{
			name: "invalid number of parcels, less than 1",
			setupMocks: func(dm *detectorMock, cm *carrierMock) {
				dm.On("Detect", testID).Once().Return(cm, nil)
				cm.On("Track", testID).Once().Return([]carriers.Parcel{}, nil)
			},
			expErr: errors.New("invalid number of parcels, expected 1 - got 0"),
		},
		{
			name: "success",
			setupMocks: func(dm *detectorMock, cm *carrierMock) {
				dm.On("Detect", testID).Once().Return(cm, nil)
				cm.On("Track", testID).Once().
					Return([]carriers.Parcel{{Number: "123", Address: "223", Status: "323"}}, nil)
			},
			expParcel: carriers.Parcel{Number: "123", Address: "223", Status: "323"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dm := &detectorMock{}
			cm := &carrierMock{}

			tc.setupMocks(dm, cm)

			tr := NewParcelsTracker(dm)
			gotParcel, gotErr := tr.TrackParcel(context.Background(), testID)

			assert.Equal(t, tc.expParcel, gotParcel)
			assert.Equal(t, tc.expErr, gotErr)
		})
	}
}

type detectorMock struct {
	mock.Mock
}

func (m *detectorMock) Detect(trackID string) (carriers.Carrier, error) {
	ret := m.Called(trackID)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).(carriers.Carrier), ret.Error(1)
}

type carrierMock struct {
	mock.Mock
}

func (m *carrierMock) Track(trackID string) ([]carriers.Parcel, error) {
	ret := m.Called(trackID)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).([]carriers.Parcel), ret.Error(1)
}

func (m *carrierMock) Detect(trackID string) bool {
	return m.Called(trackID).Bool(0)
}
