package determine_delivery

import (
	"testing"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/me"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/np"
	np_shopping "github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/np-shopping"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/ups"
	"github.com/stretchr/testify/assert"
)

func TestDetermine(t *testing.T) {
	testCases := []struct {
		name    string
		trackId string
		isError bool
	}{
		{
			name:    "UPU",
			trackId: "1z0000000000000001",
		},
		{
			name:    "NovaPoshta",
			trackId: "59000000000001",
		},
		{
			name:    "MeestExpress",
			trackId: "CV999999999ZZ",
		},
		{
			name:    "NPShopping",
			trackId: "NP99999999999999NPG",
		},
		{
			name:    "unknown",
			trackId: "59000",
			isError: true,
		},
	}

	detector := NewDetector()
	detector.Registry(np.NewCarrier(np.NewApi("", "")))
	detector.Registry(me.NewCarrier(me.NewApi("", "", "", "")))
	detector.Registry(np_shopping.NewCarrier(np_shopping.NewApi()))
	detector.Registry(ups.NewCarrier(ups.NewApi()))

	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {

			_, err := detector.Detect(testCase.trackId)
			if !testCase.isError {
				assert.NoError(t, err)
			}
			if testCase.isError {
				assert.Error(t, err)
			}
		})
	}
}
