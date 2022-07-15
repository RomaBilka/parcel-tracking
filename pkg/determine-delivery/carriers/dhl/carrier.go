package dhl

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

var patterns = map[string]*regexp.Regexp{
	"start000": regexp.MustCompile(`^000`),

	"startJVGL": regexp.MustCompile(`(?i)^JVGL`),

	"startGM": regexp.MustCompile(`(?i)^GM`),

	"startLX": regexp.MustCompile(`(?i)^LX`),

	"startRX": regexp.MustCompile(`(?i)^RX`),

	"start3S": regexp.MustCompile(`(?i)^3S`),

	"startJJD": regexp.MustCompile(`(?i)^JJD`),

	// Example: 1AB12345
	"pattern1": regexp.MustCompile(`(?i)^[\d]{1}[a-z]{2}[\d]{4,6}$`),

	// Example: ABC123456
	"pattern2": regexp.MustCompile(`(?i)^[a-z]{3}[\d]{6}$`),

	// Example: 123-12345678
	"pattern3": regexp.MustCompile(`^[\d]{3}-[\d]{8}$`),

	// Example: ABC-DE-1234567
	"pattern4": regexp.MustCompile(`(?i)^[a-z]{2,3}-[a-z]{2,3}-[\d]{7}$`),

	// Example: 1234-12345
	"pattern5": regexp.MustCompile(`^[\d]{4}-[\d]{5}$`),

	// Example: 123456789
	// 9 and 10 in UPS as well !!!
	"numbers7":    regexp.MustCompile(`^[\d]{7}$`),
	"numbers9_10": regexp.MustCompile(`^[\d]{9,10}$`),
	"numbers14":   regexp.MustCompile(`^[\d]{14}$`),
}

type api interface {
	TrackingDocument(string) (*response, error)
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

func (c *Carrier) Track(trackNumber string) ([]carriers.Parcel, error) {
	response, err := c.api.TrackingDocument(trackNumber)
	if err != nil {
		return nil, err
	}

	parcels := make([]carriers.Parcel, len(response.Shipments))
	for i, d := range response.Shipments {
		parcels[i] = carriers.Parcel{
			Number:  d.Id,
			Address: d.Status.Location.StreetAddress,
			Status:  d.Status.Status,
		}
	}

	return parcels, nil
}
