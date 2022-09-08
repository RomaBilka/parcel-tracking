package np_shopping

import (
	"regexp"
	"strconv"
	"time"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

var patterns = map[string]*regexp.Regexp{
	"start NPI": regexp.MustCompile(`^NPI`),
	"end NPI":   regexp.MustCompile(`NPI$`),
	"start NPG": regexp.MustCompile(`^NPG`),
	"end NPG":   regexp.MustCompile(`NPG$`),
}

type api interface {
	TrackByTrackingNumber(string) (*TrackingDocumentResponse, error)
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

	places := make([]carriers.Place, len(response.TrackingHistory))
	for i, d := range response.TrackingHistory {
		timestamp, err := strconv.ParseInt(d.Date, 10, 64)
		if err != nil {
			panic(err)
		}

		places[i] = carriers.Place{
			County:  d.Country,
			Comment: d.Description,
			Date:    time.Unix(timestamp, 0),
		}
	}

	parcels := []carriers.Parcel{
		carriers.Parcel{
			TrackingNumber: response.WaybillNumber,
			Status:         response.State,
			Places:         places,
		},
	}

	return parcels, nil
}
