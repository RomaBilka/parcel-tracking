package determine_delivery

import (
	"errors"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

type C interface {
	Detect(string) bool
	GetCarrier() carriers.Carrier
}

type Detector struct {
	carries []C
}

func NewDetector() *Detector {
	return &Detector{}
}

func (d *Detector) Registry(c C) {
	d.carries = append(d.carries, c)
}

func (d *Detector) Detect(trackId string) (carriers.Carrier, error) {
	for _, carrier := range d.carries {
		if carrier.Detect(trackId) {
			return carrier.GetCarrier(), nil
		}
	}

	return nil, errors.New("Carrier not detected")
}
