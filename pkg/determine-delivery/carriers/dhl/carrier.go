package dhl

import "github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"

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

	/*
		document, err := c.api.TrackingDocument(trackNumber)

	*/
	return parcels, nil
}
