package dhl

import (
	"fmt"
	"testing"
)

func TestTrackingDocument(t *testing.T) {
	dhl := NewApi("https://api-eu.dhl.com/", "demo-key")

	fmt.Println(dhl.TrackingDocument("4211612230"))
	/*
		data, err := dhl.TrackingDocument("")
		fmt.Println(data, err)*/
}
