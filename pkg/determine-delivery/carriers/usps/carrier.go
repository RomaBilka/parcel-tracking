package usps

import "github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"

var patterns string

type Carrier struct {
	api *Api
}

func NewCarrier(api *Api) *Carrier {
	return &Carrier{
		api: api,
	}
}

func (c *Carrier) Detect(trackId string) bool {
	return false
}

func (c *Carrier) Track(trackNumber string) ([]carriers.Parcel, error) {
	resp, err := c.api.TrackingDocument(trackNumber)
	if err != nil {
		return nil, err
	}

	parcels := make([]carriers.Parcel, len(resp.details))
	for i, d := range resp.details {
		parcels[i] = carriers.Parcel{
			Number: resp.number,
			Status: d,
		}
	}

	return parcels, nil
}
