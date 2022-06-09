package me

import (
	"testing"
)

func TestTrackingDocument(t *testing.T) {
	me := NewMeestExpress("0xA79E003048D2B47311E26B7D4A430FFC", "public", "PUBLIC", "https://apii.meest-group.com/T/1C_Query.php")

	me.ShipmentsTrack("TESTIK11")
}
