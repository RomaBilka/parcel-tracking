package np_shopping

import (
	"regexp"
	"strconv"
	"sync"
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

func (c *Carrier) Track(trackNumbers []string) ([]carriers.Parcel, error) {
	chanErr := make(chan error)
	chanParcels := make(chan []carriers.Parcel)

	go c.track(trackNumbers, chanParcels, chanErr)

	for {
		select {
		case err := <-chanErr:
			return nil, err
		case p := <-chanParcels:
			return p, nil
		}
	}
}

func (c *Carrier) track(trackNumbers []string, chanParcels chan []carriers.Parcel, chanErr chan error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	parcels := make([]carriers.Parcel, len(trackNumbers))
	for i := range trackNumbers {
		wg.Add(1)
		go func(trackNumber string, i int) {
			defer wg.Done()
			response, err := c.api.TrackByTrackingNumber(trackNumber)
			if err != nil {
				chanErr <- err
				return
			}

			parcel, err := prepareResponse(response)
			if err != nil {
				chanErr <- err
				return
			}
			mu.Lock()
			defer mu.Unlock()
			parcels[i] = parcel
		}(trackNumbers[i], i)
	}

	wg.Wait()
	chanParcels <- parcels
}

func prepareResponse(response *TrackingDocumentResponse) (carriers.Parcel, error) {
	places := make([]carriers.Place, len(response.TrackingHistory))
	for i, d := range response.TrackingHistory {
		timestamp, err := strconv.ParseInt(d.Date, 10, 64)
		if err != nil {
			return carriers.Parcel{}, err
		}

		places[i] = carriers.Place{
			Country: d.Country,
			Comment: d.Description,
			Date:    time.Unix(timestamp, 0),
		}
	}

	parcel := carriers.Parcel{
		TrackingNumber: response.WaybillNumber,
		Status:         response.State,
		Places:         places,
	}

	return parcel, nil
}
