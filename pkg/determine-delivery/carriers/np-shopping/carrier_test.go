package np_shopping

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
		{name: "NPI99999999999999 true", trackId: "NPI99999999999999", ok: true},
		{name: "NPI12999999999999 true", trackId: "NPI12999999999999", ok: true},
		{name: "NPI9999999999999Z false", trackId: "cv999999999zz", ok: false},
		{name: "NPI9999999NPG false", trackId: "NP9999999NPG", ok: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			n := Carrier{}
			ok := n.Detect(testCase.trackId)
			assert.Equal(t, testCase.ok, ok)
		})
	}
}
