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

				places := make([]carriers.Place, len(response.TrackingHistory))
				for i, d := range response.TrackingHistory {
					timestamp, err := strconv.ParseInt(d.Date, 10, 64)
					if err != nil {
						chanErr <- err
						return
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

				mu.Lock()
				parcels = append(parcels, parcel)
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
