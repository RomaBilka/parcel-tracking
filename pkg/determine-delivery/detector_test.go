package determine_delivery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetermine(t *testing.T) {
	testCases := []struct {
		name     string
		trackId  string
		delivery string
	}{
		{
			name:     "UPU",
			trackId:  "1z0000000000000001",
			delivery: UPS,
		},
		{
			name:     "NovaPoshta",
			trackId:  "59000000000001",
			delivery: NOVA_POSHTA,
		},
		{
			name:     "MeestExpress",
			trackId:  "CV999999999ZZ",
			delivery: MEEST_EXPRESS,
		},
		{
			name:     "NPShopping",
			trackId:  "NP99999999999999NPG",
			delivery: NP_SHOPPING,
		},
		{
			name:     "unknown",
			trackId:  "59000",
			delivery: "",
		},
	}

	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {
			delivery := Detect(testCase.trackId)
			assert.Equal(t, testCase.delivery, delivery)
		})
	}
}
