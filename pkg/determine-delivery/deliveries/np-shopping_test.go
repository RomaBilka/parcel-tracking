package deliveries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNpShopping(t *testing.T) {
	testCases := []struct {
		name    string
		trackId string
		ok      bool
	}{
		{
			name:    "NP99999999999999NPG true",
			trackId: "NP99999999999999NPG",
			ok:      true,
		},
		{
			name:    "NP12999999999999NPG true",
			trackId: "NP12999999999999NPG",
			ok:      true,
		},
		{
			name:    "NP9999999999999ZNPG false",
			trackId: "cv999999999zz",
			ok:      false,
		},
		{
			name:    "NP9999999NPG false",
			trackId: "NP9999999NPG",
			ok:      false,
		},
	}

	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {
			ok, err := IsNpShopping(testCase.trackId)
			assert.NoError(t, err)
			assert.Equal(t, testCase.ok, ok)
		})
	}
}
