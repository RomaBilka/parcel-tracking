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
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/ups"
	"github.com/stretchr/testify/assert"
)

func TestDetermine(t *testing.T) {
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
	npShoppingCarrier := np_shopping.NewCarrier(np_shopping.NewApi(""))
	detector.Registry(npShoppingCarrier)

	testCases := []struct {
		name    string
		trackId string
		carrier carriers.Carrier
		err     error
	}{
		{name: "UPS 1Z12345E6605272234", trackId: "1Z12345E6605272234", carrier: upsCarrier},
		{name: "UPS 1Z123456E6605272234", trackId: "1Z123456E6605272234", carrier: upsCarrier},
		{name: "UPS 1Z123456E660527223", trackId: "1Z123456E660527223", carrier: upsCarrier},
		{name: "UPS 1ZWX0692YP40636269", trackId: "1ZWX0692YP40636269", carrier: upsCarrier},
		{name: "UPS cgish000116630", trackId: "cgish000116630", carrier: upsCarrier},
		{name: "NovaPoshta 59000000000001", trackId: "59000000000001", carrier: npCarrier},
		{name: "NovaPoshta 10000000000001", trackId: "10000000000001", carrier: npCarrier},
		{name: "NovaPoshta 20000000000001", trackId: "20000000000001", carrier: npCarrier},
		{name: "MeestExpress CV999999999ZZ", trackId: "CV999999999ZZ", carrier: meCarrier},
		{name: "MeestExpress MYCV999999999ZZ", trackId: "MYCV999999999ZZ", carrier: meCarrier},

		{name: "NP Shopping NPI99999999999999", trackId: "NPI99999999999999", carrier: npShoppingCarrier},

		{name: "DHL 0001111111111", trackId: "0001111111111", carrier: dhlCarrier},
		{name: "DHL JVGL1111111111", trackId: "JVGL1111111111", carrier: dhlCarrier},
		{name: "DHL GM1111111111", trackId: "GM1111111111", carrier: dhlCarrier},
		{name: "DHL LX1111111111", trackId: "LX1111111111", carrier: dhlCarrier},
		{name: "DHL RX1111111111", trackId: "RX1111111111", carrier: dhlCarrier},
		{name: "DHL 3S1111111111", trackId: "3S1111111111", carrier: dhlCarrier},
		{name: "DHL JJD1111111111", trackId: "JJD1111111111", carrier: dhlCarrier},
		{name: "DHL 1234-12345", trackId: "1234-12345", carrier: dhlCarrier},
		{name: "DHL 123-12345678", trackId: "1234-12345", carrier: dhlCarrier},
		{name: "DHL ABC123456", trackId: "ABC123456", carrier: dhlCarrier},
		{name: "DHL 1AB123456", trackId: "1AB123456", carrier: dhlCarrier},
		{name: "DHL 1AB12345", trackId: "1AB12345", carrier: dhlCarrier},
		{name: "DHL 1AB1234", trackId: "1AB1234", carrier: dhlCarrier},
		{name: "DHL ABC-ABC-1234567", trackId: "ABC-ABC-1234567", carrier: dhlCarrier},
		{name: "DHL AB-AB-1234567", trackId: "AB-AB-1234567", carrier: dhlCarrier},
		{name: "DHL ABC-AB-1234567", trackId: "ABC-AB-1234567", carrier: dhlCarrier},
		{name: "DHL AB-ABC-1234567", trackId: "AB-ABC-1234567", carrier: dhlCarrier},
		{name: "Fedex 123456789012", trackId: "123456789012", carrier: fedexCarrier},
		{name: "Fedex 123456789012345", trackId: "123456789012345", carrier: fedexCarrier},
		{name: "Fedex 12345678901234567890", trackId: "12345678901234567890", carrier: fedexCarrier},
		{name: "unknown", trackId: "59000", err: errors.New("carrier not detected")},
	}

	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {
			carrier, err := detector.Detect(testCase.trackId)
			assert.Equal(t, testCase.carrier, carrier)
			assert.Equal(t, testCase.err, err)
		})
	}
}
