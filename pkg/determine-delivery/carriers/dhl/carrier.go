package dhl

import (
	"regexp"
	"sync"
	"time"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/RomaBilka/parcel-tracking/pkg/helpers"
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

const layout = "2006-01-02T15:04:05Z"

type api interface {
	TrackByTrackingNumber(string) (*response, error)
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
	var mu sync.Mutex
	var wg sync.WaitGroup
	chanErr := make(chan error)
	defer close(chanErr)
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
				p, err := prepareResponse(response)
				if err != nil {
					chanErr <- err
					return
				}
				mu.Lock()
				parcels = append(parcels, p...)
				mu.Lock()
			}(trackNumber)
		}
	}()

	if err := <-chanErr; err != nil {
		return nil, err
	}

	wg.Wait()

	return parcels, nil
}

func prepareResponse(response *response) ([]carriers.Parcel, error) {
	parcels := make([]carriers.Parcel, len(response.Shipments))
	for i, s := range response.Shipments {
		estimatedTimeOfDelivery, err := helpers.ParseTime(layout, s.EstimatedTimeOfDelivery)
		if err != nil && s.EstimatedTimeOfDelivery != "" {
			return nil, err
		}

		places := make([]carriers.Place, len(s.Events))

		for i, e := range s.Events {
			date, err := time.Parse(layout, e.Timestamp)
			if err != nil && e.Timestamp != "" {
				return nil, err
			}
			places[i] = carriers.Place{
				Country: e.Location.Address.CountryCode,
				Street:  e.Location.Address.StreetAddress,
				Address: e.Location.Address.AddressLocality,
				Date:    date,
				Comment: helpers.ConcatenateStrings(", ", e.Description, e.Remark),
			}
		}

		parcels[i] = carriers.Parcel{
			TrackingNumber: s.Id,
			DeliveryDate:   estimatedTimeOfDelivery,
			Status:         s.Status.Status,
			Places:         places,
		}
	}
	return parcels, nil
}
