package ups

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

var patterns = map[string]*regexp.Regexp{
	//Starts with 1Z, length 18
	"start1z": regexp.MustCompile(`(?i)^1Z[\d]{16}$`),

	//Starts with 8, length 18
	"start8": regexp.MustCompile(`^8[\d]{17}$`),

	//Starts with 9, length 18
	"start9": regexp.MustCompile(`^9[\d]{17}$`),
}

type api interface {
}

type Carrier struct {
	api api
}

func (c *Carrier) Detect(trackId string) bool {
	for _, pattern := range patterns {
		if pattern.MatchString(trackId) {
			return true
		}
	}

	return false
}

func NewCarrier(api *Api) *Carrier {
	return &Carrier{
		api: api,
	}
}

func (c *Carrier) Track(trackingId string) ([]carriers.Parcel, error) {
	return []carriers.Parcel{}, nil
}
