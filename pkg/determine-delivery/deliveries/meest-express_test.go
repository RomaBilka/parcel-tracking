package deliveries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMeestExpress(t *testing.T) {
	testCases := []struct {
		name    string
		trackId string
		ok      bool
	}{
		{
			name:    "CV true",
			trackId: "CV999999999ZZ",
			ok:      true,
		},
		{
			name:    "cv true",
			trackId: "cv999999999zz",
			ok:      true,
		},
		{
			name:    "CV_1 false",
			trackId: "CV999999999ZZz",
			ok:      false,
		},
		{
			name:    "CV_2 false",
			trackId: "CV9999999999ZZ",
			ok:      false,
		},
		{
			name:    "MYCV true",
			trackId: "MYCV999999999ZZ",
			ok:      true,
		},
		{
			name:    "unknown",
			trackId: "ZZZZ999999999ZZ",
			ok:      false,
		},
	}

	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {
			ok, err := IsMeestExpress(testCase.trackId)
			assert.NoError(t, err)
			assert.Equal(t, testCase.ok, ok)
		})
	}
}
