package ups

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

var start1z *regexp.Regexp
var start8 *regexp.Regexp
var start9 *regexp.Regexp

func init() {
	//1Z**************** length 18
	start1z = regexp.MustCompile(`(?i)^1Z[\d]{16}$`)

	//8***************** length 18
	start8 = regexp.MustCompile(`^8[\d]{17}$`)

	//9***************** length 18
	start9 = regexp.MustCompile(`^9[\d]{17}$`)
}

type Detector struct {
	carrier carriers.Carrier
}

func NewDetector(carrier carriers.Carrier) *Detector {
	return &Detector{carrier: carrier}
}

func (d *Detector) Detect(trackId string) bool {

	matched := start1z.MatchString(trackId)
	if matched {
		return true
	}

	matched = start8.MatchString(trackId)
	if matched {
		return true
	}

	matched = start9.MatchString(trackId)

	return matched
}

func (d *Detector) GetCarrier() carriers.Carrier {
	return d.carrier
}
