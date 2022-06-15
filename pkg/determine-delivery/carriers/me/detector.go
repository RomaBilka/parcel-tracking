package me

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

var startCV *regexp.Regexp
var startMYCV *regexp.Regexp

func init() {
	//CV999999999ZZ
	startCV = regexp.MustCompile(`(?i)^CV[\d]{9}[a-z][a-z]$`)

	//MYCV999999999ZZ
	startMYCV = regexp.MustCompile(`(?i)^MYCV[\d]{9}[a-z][a-z]$`)
}

type Detector struct {
	carrier carriers.Carrier
}

func NewDetector(carrier carriers.Carrier) *Detector {
	return &Detector{
		carrier: carrier,
	}
}

func (d *Detector) Detect(trackId string) bool {
	matched := startCV.MatchString(trackId)
	if matched {
		return true
	}

	matched = startMYCV.MatchString(trackId)
	if matched {
		return true
	}

	return false
}

func (d *Detector) GetCarrier() carriers.Carrier {
	return d.carrier
}
