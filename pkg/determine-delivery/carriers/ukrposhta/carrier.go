package ukrposhta

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/RomaBilka/parcel-tracking/pkg/helpers"
)

var patterns = map[string]*regexp.Regexp{
	//Starts with 2 letters 9 numbers and ends with 2 letters
	"pattern1": regexp.MustCompile(`(?i)^[a-z]{2}[\d]{9}[a-z]{2}$`),

	//Numeric only with the length 13
	"numbers13": regexp.MustCompile(`^[\d]{13}$`),
}

const layout = "2019-02-07T16:36:00"

type api interface {
	TrackByTrackingNumber(barcodes []string) (*Response, error)
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
	response, err := c.api.TrackByTrackingNumber([]string{trackingId})
	if err != nil {
		return nil, err
	}

	parcels := make([]carriers.Parcel, len(response.Found))
	i := 0
	for trackingNumber, b := range response.Found {
		place := make([]carriers.Place, len(b))
		for j, item := range b {
			date, err := helpers.ParseTime(item.Date, layout)
			if err != nil {
				return nil, err
			}

			place[j] = carriers.Place{
				Country: item.Country,
				Address: item.Name,
				Comment: item.EventName,
				Date:    date,
			}
		}

		parcels[i] = carriers.Parcel{
			TrackingNumber: trackingNumber,
			Places:         place,
		}
		i++
	}

	return parcels, nil
}
