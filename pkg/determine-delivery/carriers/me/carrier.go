package me

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

//Starts with CV, 9 numbers and 2 letters at the end
//CV999999999ZZ
var startCV = regexp.MustCompile(`(?i)^CV[\d]{9}[a-z]{2}$`)

//Starts with MYCV, 9 numbers and 2 letters at the end
//MYCV999999999ZZ
var startMYCV = regexp.MustCompile(`(?i)^MYCV[\d]{9}[a-z]{2}$`)

type Carrier struct {
	api *Api
}

func NewCarrier(api *Api) *Carrier {
	return &Carrier{
		api: api,
	}
}

func (c *Carrier) Detect(trackId string) bool {
	if startCV.MatchString(trackId) {
		return true
	}

	if startMYCV.MatchString(trackId) {
		return true
	}

	return false
}

func (c *Carrier) Track(trackingId string) ([]carriers.Parcel, error) {
	documents, err := c.api.ShipmentsTrack(trackingId)
	if err != nil {
		return nil, err
	}

	parcels := make([]carriers.Parcel, len(documents.ResultTable))
	for _, d := range documents.ResultTable {
		parcels = append(parcels, carriers.Parcel{
			Number:  d.ShipmentNumberSender,
			Address: d.CountryDel,
			Status:  d.ActionMessages_UA + " " + d.DetailMessages_UA,
		})
	}

	return parcels, nil
}
