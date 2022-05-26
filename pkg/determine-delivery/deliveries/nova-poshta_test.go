package deliveries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNovaPoshta(t *testing.T) {
	testCases := []struct {
		name    string
		trackId string
		ok      bool
	}{
		{
			name:    "59 true",
			trackId: "59000000000001",
			ok:      true,
		},
		{
			name:    "20 true",
			trackId: "20000000000001",
			ok:      true,
		},
		{
			name:    "1 true",
			trackId: "100000000000001",
			ok:      true,
		},
		{
			name:    "59 false",
			trackId: "5900000000000",
			ok:      false,
		},
		{
			name:    "20 false",
			trackId: "2000000000000",
			ok:      false,
		},
		{
			name:    "1 false",
			trackId: "1000000000000",
			ok:      false,
		},
		{
			name:    "unknown",
			trackId: "01234567891011",
			ok:      false,
		},
	}

	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {
			ok, err := IsNovaPoshta(testCase.trackId)
			assert.NoError(t, err)
			assert.Equal(t, testCase.ok, ok)
		})
	}
}
