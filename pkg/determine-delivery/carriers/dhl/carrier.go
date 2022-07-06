package dhl

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

//10 digits numerical only
var digits10 = regexp.MustCompile(`[\d]{10}$`)

type Carrier struct {
	api *Api
}

func NewCarrier(api *Api) *Carrier {
	return &Carrier{
		api: api,
	}
}

func (c *Carrier) Detect(trackId string) bool {
	return true
}

func (c *Carrier) Track(trackNumber string) ([]carriers.Parcel, error) {
	parcels := make([]carriers.Parcel, 1)

	response, err := c.api.TrackingDocument(trackNumber)
	if err != nil {
		return nil, err
	}

	for _, d := range response.Shipments {
		parcels = append(parcels, carriers.Parcel{
			Number:  d.Id,
			Address: d.Status.Location.StreetAddress,
			Status:  d.Status.Status,
		})
	}

	return parcels, nil
}
