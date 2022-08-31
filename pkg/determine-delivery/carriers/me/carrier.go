package me

import (
	"regexp"
	"time"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

var patterns = map[string]*regexp.Regexp{
	//Starts with CV, 9 numbers and 2 letters at the end
	//CV999999999ZZ
	"startCV": regexp.MustCompile(`(?i)^CV[\d]{9}[a-z]{2}$`),

	//Starts with MYCV, 9 numbers and 2 letters at the end
	//MYCV999999999ZZ
	"startMYCV": regexp.MustCompile(`(?i)^MYCV[\d]{9}[a-z]{2}$`),
}

const layout = "2016-06-30 13: 42: 11"

type api interface {
	TrackByTrackingNumber(string) (*ShipmentsTrackResponse, error)
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
	response, err := c.api.TrackByTrackingNumber(trackingId)
	if err != nil {
		return nil, err
	}

	parcels := make([]carriers.Parcel, len(response.ResultTable))
	for i, d := range response.ResultTable {
		parcels[i] = carriers.Parcel{
			Number:  d.ShipmentNumberSender,
			Address: d.CountryDel,
			Status:  d.ActionMessages_UA + " " + d.DetailMessages_UA,
		}
	}

	return parcels, nil
}

func (c *Carrier) Track_draft(trackingId string) ([]carriers.Parcel_draft, error) {
	response, err := c.api.TrackByTrackingNumber(trackingId)
	if err != nil {
		return nil, err
	}

	if len(response.ResultTable) == 0 {
		return nil, nil
	}

	places, err := getPlaces(response.ResultTable)
	if err != nil {
		return nil, err
	}

	parcels := []carriers.Parcel_draft{
		carriers.Parcel_draft{
			TrackingNumber: response.ResultTable[0].ShipmentNumberSender,
			Places:         places,
		},
	}

	return parcels, nil
}

func getPlaces(result []ShipmentTrackResponse) ([]carriers.Place, error) {
	places := make([]carriers.Place, len(result))

	for i, s := range result {
		date, err := time.Parse(layout, s.DateTimeAction)
		if err != nil {
			return nil, err
		}
		places[i] = carriers.Place{
			County: s.Country,
			City:   s.City,
			Date:   date,
		}
	}

	return places, nil
}
