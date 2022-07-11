package np_shopping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetect(t *testing.T) {
	testCases := []struct {
		name    string
		trackId string
		ok      bool
	}{
		{name: "NP99999999999999NPG true", trackId: "NP99999999999999NPG", ok: true},
		{name: "NP12999999999999NPG true", trackId: "NP12999999999999NPG", ok: true},
		{name: "NP9999999999999ZNPG false", trackId: "cv999999999zz", ok: false},
		{name: "NP9999999NPG false", trackId: "NP9999999NPG", ok: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			n := NewCarrier(NewApi())
			ok := n.Detect(testCase.trackId)
			assert.Equal(t, testCase.ok, ok)
		})
	}
}
