package ups

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

//Starts with 1Z, length 18
var start1z = regexp.MustCompile(`(?i)^1Z[\d]{16}$`)

//Starts with 8, length 18
var start8 = regexp.MustCompile(`^8[\d]{17}$`)

//Starts with 9, length 18
var start9 = regexp.MustCompile(`^9[\d]{17}$`)

type Carrier struct {
	api *Api
}

func (c *Carrier) Detect(trackId string) bool {
	if start1z.MatchString(trackId) {
		return true
	}

	if start8.MatchString(trackId) {
		return true
	}

	if start9.MatchString(trackId) {
		return true
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
