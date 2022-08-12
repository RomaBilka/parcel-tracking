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
		{name: "start 94001 & 17 digits", trackId: "9400100000000000000000", ok: true},
		{name: "start 94111 & 17 digits", trackId: "9411100000000000000000", ok: false},
		{name: "start 94001 & 18 digits", trackId: "94001000000000000000000", ok: false},
		{name: "start 94001 & 16 digits", trackId: "940010000000000000000", ok: false},
		{name: "start 92055 & 17 digits", trackId: "9205500000000000000000", ok: true},
		{name: "start 92155 & 17 digits", trackId: "9215500000000000000000", ok: false},
		{name: "start 92055 & 18 digits", trackId: "92055000000000000000000", ok: false},
		{name: "start 92055 & 16 digits", trackId: "920550000000000000000", ok: false},
		{name: "start 94073 & 17 digits", trackId: "9407300000000000000000", ok: true},
		{name: "start 94173 & 17 digits", trackId: "9417300000000000000000", ok: false},
		{name: "start 94073 & 18 digits", trackId: "94073000000000000000000", ok: false},
		{name: "start 94073 & 16 digits", trackId: "940730000000000000000", ok: false},
		{name: "start 93033 & 17 digits", trackId: "9303300000000000000000", ok: true},
		{name: "start 93133 & 17 digits", trackId: "9313300000000000000000", ok: false},
		{name: "start 93033 & 16 digits", trackId: "930330000000000000000", ok: false},
		{name: "start 93033 & 18 digits", trackId: "93033000000000000000000", ok: false},
		{name: "start 92088 & 17 digits", trackId: "9208800000000000000000", ok: true},
		{name: "start 92188 & 17 digits", trackId: "9218800000000000000000", ok: false},
		{name: "start 92088 & 16 digits", trackId: "920880000000000000000", ok: false},
		{name: "start 92088 & 18 digits", trackId: "92088000000000000000000", ok: false},
		{name: "start 92021 & 17 digits", trackId: "9202100000000000000000", ok: true},
		{name: "start 92121 & 17 digits", trackId: "9212100000000000000000", ok: false},
		{name: "start 92021 & 16 digits", trackId: "920210000000000000000", ok: false},
		{name: "start 92021 & 18 digits", trackId: "92021000000000000000000", ok: false},
		{name: "start 92701 & 17 digits", trackId: "9270100000000000000000", ok: true},
		{name: "start 92711 & 17 digits", trackId: "9271100000000000000000", ok: false},
		{name: "start 92701 & 16 digits", trackId: "927010000000000000000", ok: false},
		{name: "start 92701 & 18 digits", trackId: "92701000000000000000000", ok: false},
		{name: "start EC & 9 digits US", trackId: "EC000000000US", ok: true},
		{name: "start EC & 8 digits US", trackId: "EC00000000US", ok: false},
		{name: "start EC & 10 digits US", trackId: "EC0000000000US", ok: false},
		{name: "start EA & 9 digits US", trackId: "EA000000000US", ok: true},
		{name: "start EA & 8 digits US", trackId: "EA00000000US", ok: false},
		{name: "start EA & 10 digits US", trackId: "EA0000000000US", ok: false},
		{name: "start CP & 9 digits US", trackId: "CP000000000US", ok: true},
		{name: "start CP & 8 digits US", trackId: "CP00000000US", ok: false},
		{name: "start CP & 10 digits US", trackId: "CP0000000000US", ok: false},
		{name: "start AA & 9 digits US", trackId: "AA0000000000US", ok: false},
		{name: "82  & 8 digits", trackId: "8200000000", ok: true},
		{name: "10 digits", trackId: "1000000000", ok: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			u := NewCarrier(NewApi("", "", ""))
			ok := u.Detect(testCase.trackId)
			assert.Equal(t, testCase.ok, ok)
		})
	}
}
