package fedex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarrier_Detect(t *testing.T) {
	testCases := []struct {
		name    string
		trackId string
		ok      bool
	}{
		{name: "12 true", trackId: "123456789012", ok: true},
		{name: "15 true", trackId: "123456789012345", ok: true},
		{name: "20 true", trackId: "12345678901234567890", ok: true},
		{name: "22 true", trackId: "1234567890123456789012", ok: true},
		{name: "21 false", trackId: "123456789012345678901", ok: false},
		{name: "10 false", trackId: "1234567890", ok: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			fedex := NewCarrier(NewApi("", "", "", ""))
			ok := fedex.Detect(testCase.trackId)
			assert.Equal(t, testCase.ok, ok)
		})
	}
}
