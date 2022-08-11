package determine_delivery

import (
	"testing"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDetector_Detect_Fails(t *testing.T) {
	trackId := "test-track-id"

	m := &mockCarrier{}
	m.On("Detect", trackId).Twice().Return(false)

	detector := NewDetector()
	detector.Registry(m)
	detector.Registry(m)

	carrier, err := detector.Detect(trackId)
	assert.Nil(t, carrier, "carrier should be nil")
	assert.Equal(t, errCarrierNotDetected, err)
}

func TestDetector_Detect_Success(t *testing.T) {
	trackId := "test-track-id"

	invalidCarrier := &mockCarrier{}
	invalidCarrier.On("Detect", trackId).Twice().Return(false)

	validCarrier := &mockCarrier{}
	validCarrier.On("Detect", trackId).Once().Return(true)

	detector := NewDetector()
	detector.Registry(invalidCarrier)
	detector.Registry(invalidCarrier)
	detector.Registry(validCarrier)

	carrier, err := detector.Detect(trackId)
	assert.NoError(t, err)
	assert.Equal(t, carrier, validCarrier)
}

type mockCarrier struct {
	mock.Mock
}

func (m *mockCarrier) Track(s string) ([]carriers.Parcel, error) {
	args := m.Called(s)
	if args.Get(0) != nil {
		return args.Get(0).([]carriers.Parcel), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockCarrier) Detect(s string) bool {
	return m.Called(s).Bool(0)
}
