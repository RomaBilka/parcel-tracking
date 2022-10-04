package ups

import (
	"regexp"
	"sync"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/RomaBilka/parcel-tracking/pkg/helpers"
)

var patterns = map[string]*regexp.Regexp{
	//Starts with 1Z, 5 or 6 digits, 1 letter and 9 or 10 digits, length 18
	//1Z12345E6605272234
	"start1z": regexp.MustCompile(`(?i)^1Z[\d]{5,6}[a-z]{1}[\d]{9,10}$`),

	//Starts with 1Z, 4 digits, WWE and 9 digits, length 18
	//1Z4861WWE194914215
	"WWE": regexp.MustCompile(`(?i)^1Z[\d]{4}WWE[\d]{9}$`),

	//Starts with 1ZWX, 4 digits, 2 letters and 8 digits, length 18
	//1ZWX0692YP40636269
	"start1ZWX": regexp.MustCompile(`(?i)^1ZWX[\d]{4}[a-z]{2}[\d]{8}$`),

	//Starts with ER and 15 digits, length 17
	//ER751105042015062
	"startER": regexp.MustCompile(`(?i)^ER[\d]{15}`),

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
	TrackByTrackingNumber(string) (*TrackResponse, error)
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

func (c *Carrier) Track(trackNumbers []string) ([]carriers.Parcel, error) {
	var wg sync.WaitGroup
	chanErr := make(chan error)
	chanParcel := make(chan carriers.Parcel)
	defer close(chanErr)
	defer close(chanParcel)
	var parcels []carriers.Parcel

	go func() {
		for _, trackNumber := range trackNumbers {
			wg.Add(1)
			go func(trackNumber string) {
				defer wg.Done()

				response, err := c.api.TrackByTrackingNumber(trackNumber)
				if err != nil {
					chanErr <- err
					return
				}

				parcel := prepareResponse(response)

				chanParcel <- parcel
			}(trackNumber)
		}
	}()

	select {
	case err := <-chanErr:
		return nil, err
	case p := <-chanParcel:
		parcels = append(parcels, p)
	}

	wg.Wait()

	return parcels, nil
}

func prepareResponse(response *TrackResponse) carriers.Parcel {
	parcel := carriers.Parcel{
		TrackingNumber: response.Shipment.ShipmentIdentificationNumber,
		Places: []carriers.Place{
			carriers.Place{
				Country: response.Shipment.Shipper.Address.CountryCode,
				City:    response.Shipment.Shipper.Address.City,
				Address: getAddress(response.Shipment.Shipper.Address),
			},
			carriers.Place{
				Country: response.Shipment.ShipTo.Address.CountryCode,
				City:    response.Shipment.ShipTo.Address.City,
				Address: getAddress(response.Shipment.ShipTo.Address),
			},
		},
		Status: response.Shipment.CurrentStatus.Description,
	}

	return parcel
}

func getAddress(a Address) string {
	address := helpers.ConcatenateStrings(", ", a.AddressLine1, a.AddressLine2, a.AddressLine3)

	return address
}
