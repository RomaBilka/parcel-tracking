package me

import (
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

func (c *Carrier) Track(trackingId string) ([]carriers.Parcel, error) {
	parcels := []carriers.Parcel{}

	documents, err := c.api.ShipmentsTrack(trackingId)
	if err != nil {
		return parcels, err
	}

	for _, d := range documents.ResultTable {
		parcels = append(parcels, carriers.Parcel{
			Number:  d.ShipmentNumberSender,
			Address: d.CountryDel,
			Status:  d.ActionMessages_UA + " " + d.DetailMessages_UA,
		})
	}

	return parcels, nil
}
