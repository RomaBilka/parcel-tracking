package np_shopping

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

var npShopping *regexp.Regexp

func init() {
	//NP99999999999999NPG
	npShopping = regexp.MustCompile(`(?i)^NP[\d]{14}NPG$`)
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
	matched := npShopping.MatchString(trackId)
	if matched {
		return true
	}

	return false
}

func (d *Detector) GetCarrier() carriers.Carrier {
	return d.carrier
}
