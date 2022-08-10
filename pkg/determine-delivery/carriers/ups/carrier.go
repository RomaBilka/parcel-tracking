package ups

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

var patterns = map[string]*regexp.Regexp{
	//Starts with 1Z, 5 or 6 digits, 1 letter and 9 or 10 digits, length 18
	//1Z12345E6605272234
	"start1z": regexp.MustCompile(`(?i)^1Z[\d]{5,6}[a-z]{1}[\d]{9,10}$`),

	//Starts with 1ZWX, 4 digits, 2 letters and 8 digits, length 18
	//1ZWX0692YP40636269
	"start1ZWX": regexp.MustCompile(`(?i)^1ZWX[\d]{4}[a-z]{2}[\d]{8}$`),

	//Starts with ER and 15 digits, length 17
	//ER751105042015062
	"startER": regexp.MustCompile(`(?i)^ER[\d]{15}`),

	//Numeric only, starts 8, length 18
	"start8": regexp.MustCompile(`^8[\d]{17}$`),

	//Numeric only, starts with 9, length 18
	"start9": regexp.MustCompile(`^9[\d]{17}$`),

	//Numeric only with the length 9 or 10
	//Example: 123456789
	//9 and 10 in DHL as well !!!
	"numbers9_10": regexp.MustCompile(`^[\d]{9,10}$`),

	//Numeric only with the length 22
	//22 in Fedex as well !!!
	"numbers22": regexp.MustCompile(`^[\d]{22}$`),

	//Starts with cgish and 9 digits
	//cgish000116630
	"startCGISH": regexp.MustCompile(`(?i)^cgish[\d]{9}$`),
}

type api interface {
	TrackByNumber(string) (*TrackResponse, error)
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

func (c *Carrier) Track(trackingNumber string) ([]carriers.Parcel, error) {
	res, err := c.api.TrackByNumber(trackingNumber)
	if err != nil {
		return nil, err
	}

	parcels := []carriers.Parcel{
		carriers.Parcel{
			Number:  res.Shipment.ShipmentIdentificationNumber,
			Status:  res.Shipment.Package.Activity[0].Status.StatusType.Description,
			Address: res.Shipment.Package.Activity[0].ActivityLocation.Address.City,
		},
	}

	return parcels, nil
}
