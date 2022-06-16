package me

import (
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

type Carrier struct {
	api *Api
}

func NewCarrier(api *Api) *Carrier {
	return &Carrier{
		api: api,
	}
}

func (c *Carrier) Tracking(trackingId string) ([]carriers.Parcel, error) {
	return []carriers.Parcel{}, nil
}