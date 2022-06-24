package ups

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUPS(t *testing.T) {
	testCases := []struct {
		name    string
		trackId string
		ok      bool
	}{
		{
			name:    "1Z true",
			trackId: "1z0000000000000001",
			ok:      true,
		},
		{
			name:    "8 true",
			trackId: "800000000000000001",
			ok:      true,
		},
		{
			name:    "9 true",
			trackId: "900000000000000001",
			ok:      true,
		},
		{
			name:    "1Z false",
			trackId: "1z000000000000000000000000",
			ok:      false,
		},
		{
			name:    "8 false",
			trackId: "80000000000000000000",
			ok:      false,
		},
		{
			name:    "9 false",
			trackId: "90000000000000000",
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
			u := NewCarrier(NewApi())
			ok := u.Detect(testCase.trackId)
			assert.Equal(t, testCase.ok, ok)
		})
	}
}
