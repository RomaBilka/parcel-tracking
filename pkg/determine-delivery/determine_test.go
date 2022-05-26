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
			name:     "unknown",
			trackId:  "59000",
			delivery: "",
		},
	}

	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {
			delivery, err := Determine(testCase.trackId)
			assert.NoError(t, err)
			assert.Equal(t, testCase.delivery, delivery)
		})
	}
}
