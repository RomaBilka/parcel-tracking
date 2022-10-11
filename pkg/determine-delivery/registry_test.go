package determine_delivery

import (
	"errors"
	"testing"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/dhl"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/fedex"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/me"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/np"
	np_shopping "github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/np-shopping"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/ukrposhta"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/ups"
	response_errors "github.com/RomaBilka/parcel-tracking/pkg/response-errors"
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
	assert.Equal(t, response_errors.CarrierNotFound, err)

	mock.AssertExpectationsForObjects(t, m)
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

	mock.AssertExpectationsForObjects(t, invalidCarrier, validCarrier)
}

type mockCarrier struct {
	mock.Mock
}

func (m *mockCarrier) Track(s []string) ([]carriers.Parcel, error) {
	args := m.Called(s)
	if args.Get(0) != nil {
		return args.Get(0).([]carriers.Parcel), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockCarrier) Detect(s string) bool {
	return m.Called(s).Bool(0)
}

func FuzzDetector_Detect(t *testing.F) {
	detector := NewDetector()
	npCarrier := np.NewCarrier(np.NewApi("", ""))
	detector.Registry(npCarrier)
	meCarrier := me.NewCarrier(me.NewApi("", "", "", ""))
	detector.Registry(meCarrier)
	fedexCarrier := fedex.NewCarrier(fedex.NewApi("", "", "", ""))
	detector.Registry(fedexCarrier)
	dhlCarrier := dhl.NewCarrier(dhl.NewApi("", ""))
	detector.Registry(dhlCarrier)
	upsCarrier := ups.NewCarrier(ups.NewApi("", "", "", ""))
	detector.Registry(upsCarrier)
	np_shoppingCarrier := np_shopping.NewCarrier(np_shopping.NewApi(""))
	detector.Registry(np_shoppingCarrier)
	ukrPoshtaCarrier := ukrposhta.NewCarrier(ukrposhta.NewApi("", ""))
	detector.Registry(ukrPoshtaCarrier)

	t.Fuzz(func(t *testing.T, trackId string) {
		_, err := detector.Detect(trackId)
		if err != nil && !errors.Is(err, response_errors.CarrierNotFound) {
			t.Errorf("Error: %v", err)
		}
	})
}
