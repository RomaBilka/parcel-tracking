package determine_delivery

import (
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	response_errors "github.com/RomaBilka/parcel-tracking/pkg/response-errors"
)

type Detector struct {
	carries []carriers.Carrier
}

func NewDetector() *Detector {
	return &Detector{}
}

func (d *Detector) Registry(c carriers.Carrier) {
	d.carries = append(d.carries, c)
}

func (d *Detector) Detect(trackId string) (carriers.Carrier, error) {
	for _, carrier := range d.carries {
		if carrier.Detect(trackId) {
			return carrier, nil
		}
	}

	return nil, response_errors.CarrierNotFound
}
