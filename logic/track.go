package logic

import (
	"context"
	"sync"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

type Detector interface {
	Detect(trackId string) (carriers.Carrier, error)
}

type ParcelsTracker struct {
	detector Detector
}

func NewParcelsTracker(detector Detector) *ParcelsTracker {
	return &ParcelsTracker{detector: detector}
}

func track(idsToCarriers map[carriers.Carrier][]string, chanParcels chan map[string]carriers.Parcel, chanErr chan error) {
	mapParcels := make(map[string]carriers.Parcel)
	var mu sync.Mutex
	var wg sync.WaitGroup
	for carrier, ids := range idsToCarriers {
		wg.Add(1)

		go func(carrier carriers.Carrier, ids []string) {
			defer wg.Done()
			parcels, err := carrier.Track(ids)

			if err != nil {
				chanErr <- err
			}

			mu.Lock()
			for _, p := range parcels {
				mapParcels[p.TrackingNumber] = p
			}
			mu.Unlock()
		}(carrier, ids)
	}
	wg.Wait()

	chanParcels <- mapParcels
}

func (p ParcelsTracker) TrackParcels(_ context.Context, parcelIds []string) (map[string]carriers.Parcel, error) {
	chanErr := make(chan error)
	chanIdsToCarriers := make(chan map[carriers.Carrier][]string)
	chanParcels := make(chan map[string]carriers.Parcel)

	go p.matchParcelIdsToCarriers(parcelIds, chanIdsToCarriers, chanErr)
	for {
		select {
		case err := <-chanErr:
			{
				return nil, err
			}
		case idsToCarriers := <-chanIdsToCarriers:
			{
				go track(idsToCarriers, chanParcels, chanErr)
			}
		case parcels := <-chanParcels:
			return parcels, nil
		}
	}
}

func (p ParcelsTracker) matchParcelIdsToCarriers(parcelIds []string, chanIdsToCarriers chan map[carriers.Carrier][]string, chanErr chan error) {
	idsToCarriers := make(map[carriers.Carrier][]string)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, parcelId := range parcelIds {
		wg.Add(1)
		go func(parcelId string) {
			defer wg.Done()
			carrier, err := p.detector.Detect(parcelId)
			if err != nil {
				chanErr <- err
				return
			}

			mu.Lock()
			idsToCarriers[carrier] = append(idsToCarriers[carrier], parcelId)
			mu.Unlock()
		}(parcelId)

	}

	wg.Wait()
	chanIdsToCarriers <- idsToCarriers
}
