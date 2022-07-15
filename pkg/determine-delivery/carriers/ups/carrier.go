package ups

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

var patterns = map[string]*regexp.Regexp{
	// 1Z12345E6605272234
	"start1z": regexp.MustCompile(`(?i)^1Z[\d]{5,6}[a-z]{1}[\d]{9,10}$`),

	"start1ZWX": regexp.MustCompile(`(?i)^1ZWX[\d]{4}[a-z]{2}[\d]{8}$`),

	"startER": regexp.MustCompile(`(?i)^ER[\d]{15}`),

	"start8": regexp.MustCompile(`^8[\d]{17}$`),

	"start9": regexp.MustCompile(`^9[\d]{17}$`),

	// Example: 123456789
	// 9 and 10 in DHL as well !!!
	"numbers9_10": regexp.MustCompile(`^[\d]{9,10}$`),

	"numbers22": regexp.MustCompile(`^[\d]{22}$`),

	"startCGISH": regexp.MustCompile(`(?i)^cgish[\d]{9}$`),
}

type Carrier struct{}

func (c *Carrier) Detect(trackId string) bool {
	for _, pattern := range patterns {
		if pattern.MatchString(trackId) {
			return true
		}
	}

	return false
}

func NewCarrier() *Carrier {
	return &Carrier{}
}

func (c *Carrier) Track(trackingId string) ([]carriers.Parcel, error) {
	return []carriers.Parcel{}, nil
}
