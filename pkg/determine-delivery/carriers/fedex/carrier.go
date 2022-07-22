package fedex

import (
	"regexp"
)

var patterns = map[string]*regexp.Regexp{
	//Numeric only with the length 12
	"numbers12": regexp.MustCompile(`^[\d]{12}$`),
	//Numeric only with the length 15
	"numbers15": regexp.MustCompile(`^[\d]{15}$`),
	//Numeric only with the length 20
	"numbers20": regexp.MustCompile(`^[\d]{20}$`),
	//Numeric only with the length 22
	"numbers22": regexp.MustCompile(`^[\d]{22}$`),
}

type api interface {
	TrackByTrackingNumber(TrackingRequest) (*Response, error)
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

/*
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
*/
