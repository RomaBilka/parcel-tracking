package logic

import (
	"context"
	"fmt"

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

func (p ParcelsTracker) TrackParcel(_ context.Context, parcelId string) (carriers.Parcel, error) {
	carrier, err := p.detector.Detect(parcelId)
	if err != nil {
		return carriers.Parcel{}, err
	}

	parcels, err := carrier.Track(parcelId)
	if err != nil {
		return carriers.Parcel{}, err
	}
	if len(parcels) != 1 {
		return carriers.Parcel{}, fmt.Errorf("invalid number of parcels, expected 1 - got %d", len(parcels))
	}

	return parcels[0], nil
}
