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

				parcel, err := prepareResponse(response)
				if err != nil {
					chanErr <- err
					return
				}

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
