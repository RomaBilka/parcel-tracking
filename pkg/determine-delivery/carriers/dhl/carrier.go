package dhl

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

var patterns = map[string]*regexp.Regexp{
	//It starts with 000
	"start000": regexp.MustCompile(`^000`),

	//It starts with JVGL
	"startJVGL": regexp.MustCompile(`(?i)^JVGL`),

	//It starts with GM
	"startGM": regexp.MustCompile(`(?i)^GM`),

	//It starts with LX
	"startLX": regexp.MustCompile(`(?i)^LX`),

	//It starts with RX
	"startRX": regexp.MustCompile(`(?i)^RX`),

	//It starts with 3S
	"start3S": regexp.MustCompile(`(?i)^3S`),

	//It starts with JJD
	"startJJD": regexp.MustCompile(`(?i)^JJD`),

	//Starts with 1 number, followed by 2 letters and 4 to 6 numbers
	//Example: 1AB12345
	"pattern1": regexp.MustCompile(`(?i)^[\d]{1}[a-z]{2}[\d]{4,6}$`),

	//Starts with 3 to 4 letters
	//Example: ABC123456
	"pattern2": regexp.MustCompile(`(?i)^[a-z]{3}[\d]{6}$`),

	//Starts with 3-digit carrier code, followed by hyphen (-), followed by the 8-digit masterbill number.
	//Example: 123-12345678
	"pattern3": regexp.MustCompile(`^[\d]{3}-[\d]{8}$`),

	//Order Code: starts with 2 to 3 letters, followed by hyphen (-), 2 to 3 letters, hyphen (-) and 7 numbers
	//Example: ABC-DE-1234567
	"pattern4": regexp.MustCompile(`(?i)^[a-z]{2,3}-[a-z]{2,3}-[\d]{7}$`),

	//Starts with 4 numbers, followed by hyphen (-) and 5 numbers
	//Example: 1234-12345
	"pattern5": regexp.MustCompile(`^[\d]{4}-[\d]{5}$`),

	//Numeric only with the length 7,9,10 or 14
	//Example: 123456789
	//9 and 10 in UPS as well !!!
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
