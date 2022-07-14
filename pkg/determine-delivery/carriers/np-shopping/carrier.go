package np_shopping

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

////Starts with NP, 14 numbers and NPG at the end
//NP99999999999999NPG
var npShopping = regexp.MustCompile(`(?i)^NP[\d]{14}NPG$`)

type api interface {
}

type Carrier struct {
}

func NewCarrier() *Carrier {
	return &Carrier{}
}

func (c *Carrier) Detect(trackId string) bool {
	return npShopping.MatchString(trackId)
}

func (c *Carrier) Track(trackingId string) ([]carriers.Parcel, error) {
	return []carriers.Parcel{}, nil
}
