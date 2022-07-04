package logic

import (
	"context"

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

func (p ParcelsTracker) TrackParcel(_ context.Context, parcelID string) (carriers.Parcel, error) {
	carrier, err := p.detector.Detect(parcelID)
	if err != nil {
		return carriers.Parcel{}, err
	}

	parcel, err := carrier.Track(parcelID)
	if err != nil {
		return carriers.Parcel{}, err
	}

	return parcel[0], nil
}
