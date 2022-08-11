package usps

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
		{name: "9400100000000000000000", trackId: "9400100000000000000000", ok: true},
		{name: "9411100000000000000000", trackId: "9411100000000000000000", ok: false},
		{name: "94001000000000000000000", trackId: "94001000000000000000000", ok: false},
		{name: "940010000000000000000", trackId: "940010000000000000000", ok: false},
		{name: "9205500000000000000000", trackId: "9205500000000000000000", ok: true},
		{name: "9215500000000000000000", trackId: "9215500000000000000000", ok: false},
		{name: "92055000000000000000000", trackId: "92055000000000000000000", ok: false},
		{name: "920550000000000000000", trackId: "920550000000000000000", ok: false},
		{name: "9407300000000000000000", trackId: "9407300000000000000000", ok: true},
		{name: "9417300000000000000000", trackId: "9417300000000000000000", ok: false},
		{name: "94073000000000000000000", trackId: "94073000000000000000000", ok: false},
		{name: "940730000000000000000", trackId: "940730000000000000000", ok: false},
		{name: "9303300000000000000000", trackId: "9303300000000000000000", ok: true},
		{name: "9313300000000000000000", trackId: "9313300000000000000000", ok: false},
		{name: "930330000000000000000", trackId: "930330000000000000000", ok: false},
		{name: "93033000000000000000000", trackId: "93033000000000000000000", ok: false},
		{name: "9208800000000000000000", trackId: "9208800000000000000000", ok: true},
		{name: "9218800000000000000000", trackId: "9218800000000000000000", ok: false},
		{name: "920880000000000000000", trackId: "920880000000000000000", ok: false},
		{name: "92088000000000000000000", trackId: "92088000000000000000000", ok: false},
		{name: "9202100000000000000000", trackId: "9202100000000000000000", ok: true},
		{name: "9212100000000000000000", trackId: "9212100000000000000000", ok: false},
		{name: "920210000000000000000", trackId: "920210000000000000000", ok: false},
		{name: "92021000000000000000000", trackId: "92021000000000000000000", ok: false},
		{name: "9270100000000000000000", trackId: "9270100000000000000000", ok: true},
		{name: "9271100000000000000000", trackId: "9271100000000000000000", ok: false},
		{name: "927010000000000000000", trackId: "927010000000000000000", ok: false},
		{name: "92701000000000000000000", trackId: "92701000000000000000000", ok: false},
		{name: "EC000000000US", trackId: "EC000000000US", ok: true},
		{name: "EC00000000US", trackId: "EC00000000US", ok: false},
		{name: "EC0000000000US", trackId: "EC0000000000US", ok: false},
		{name: "EA000000000US", trackId: "EA000000000US", ok: true},
		{name: "EA00000000US", trackId: "EA00000000US", ok: false},
		{name: "EA0000000000US", trackId: "EA0000000000US", ok: false},
		{name: "CP000000000US", trackId: "CP000000000US", ok: true},
		{name: "CP00000000US", trackId: "CP00000000US", ok: false},
		{name: "CP0000000000US", trackId: "CP0000000000US", ok: false},
		{name: "AA0000000000US", trackId: "AA0000000000US", ok: false},
		{name: "8200000000", trackId: "8200000000", ok: true},
		{name: "1000000000", trackId: "1000000000", ok: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			u := NewCarrier(NewApi("", "", ""))
			ok := u.Detect(testCase.trackId)
			assert.Equal(t, testCase.ok, ok)
		})
	}
}
