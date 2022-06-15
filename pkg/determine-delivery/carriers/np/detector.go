package np

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

var start59 *regexp.Regexp
var start20 *regexp.Regexp
var start1 *regexp.Regexp

func init() {
	//59************ length 14
	start59 = regexp.MustCompile(`^59[\d]{12}$`)

	//20************ length 14
	start20 = regexp.MustCompile(`^20[\d]{12}$`)

	//1************* length 14
	start1 = regexp.MustCompile(`^1[\d]{13}$`)
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
	matched := start59.MatchString(trackId)
	if matched {
		return true
	}

	matched = start20.MatchString(trackId)
	if matched {
		return true
	}

	matched = start1.MatchString(trackId)

	return matched
}

func (d *Detector) GetCarrier() carriers.Carrier {
	return d.carrier
}
