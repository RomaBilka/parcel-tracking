package me

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

var patterns = map[string]*regexp.Regexp{
	// CV999999999ZZ
	"startCV": regexp.MustCompile(`(?i)^CV[\d]{9}[a-z]{2}$`),

	"startMYCV": regexp.MustCompile(`(?i)^MYCV[\d]{9}[a-z]{2}$`),
}

type api interface {
	ShipmentsTrack(string) (*ShipmentsTrackResponse, error)
}

type Carrier struct {
	api api
}

func NewCarrier(api api) *Carrier {
	return &Carrier{
		api: api,
	}
}

func (c *Carrier) Detect(trackId string) bool {
	for _, pattern := range patterns {
		if pattern.MatchString(trackId) {
			return true
		}
	}

	return false
}

func (c *Carrier) Track(trackingId string) ([]carriers.Parcel, error) {
	documents, err := c.api.ShipmentsTrack(trackingId)
	if err != nil {
		return nil, err
	}

	parcels := make([]carriers.Parcel, len(documents.ResultTable))
	for i, d := range documents.ResultTable {
		parcels[i] = carriers.Parcel{
			Number:  d.ShipmentNumberSender,
			Address: d.CountryDel,
			Status:  d.ActionMessages_UA + " " + d.DetailMessages_UA,
		}
	}

	return parcels, nil
}
