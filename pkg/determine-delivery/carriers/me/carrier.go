package me

import (
	"errors"
	"regexp"
	"sync"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/RomaBilka/parcel-tracking/pkg/helpers"
)

var patterns = map[string]*regexp.Regexp{
	//Starts with CV, 9 numbers and 2 letters at the end
	//CV999999999AA
	"startCV": regexp.MustCompile(`(?i)^CV[\d]{9}[a-z]{2}$`),

	//Starts with MYCV, 9 numbers and 2 letters at the end
	//MYCV999999999AA
	"startMYCV": regexp.MustCompile(`(?i)^MYCV[\d]{9}[a-z]{2}$`),
}

const layout = "2016-06-30 13: 42: 11"

type api interface {
	TrackByTrackingNumber(string) (*ShipmentsTrackResponse, error)
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

			if len(response.ResultTable) == 0 {
				chanErr <- errors.New("empty result")
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

func prepareResponse(response *ShipmentsTrackResponse) (carriers.Parcel, error) {
	places, err := getPlaces(response.ResultTable)
	if err != nil {
		return carriers.Parcel{}, err
	}

	parcel := carriers.Parcel{
		TrackingNumber: response.ResultTable[0].ShipmentNumberSender,
		Places:         places,
	}

	return parcel, nil
}

func getPlaces(result []ShipmentTrackResponse) ([]carriers.Place, error) {
	places := make([]carriers.Place, len(result))

	for i, s := range result {
		date, err := helpers.ParseTime(layout, s.DateTimeAction)
		if err != nil {
			return nil, err
		}

		places[i] = carriers.Place{
			Country: s.Country,
			City:    s.City,
			Date:    date,
		}
	}

	return places, nil
}
