package fedex

import (
	"fmt"
	"testing"
)

func TestTrackByTrackingNumber(t *testing.T) {
	api := NewApi("", "", "", "")
	fmt.Println(api.TrackByTrackingNumber())
}
