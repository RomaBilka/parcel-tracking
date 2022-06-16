package handlers

import (
	"net/http"

	determine_delivery "github.com/RomaBilka/parcel-tracking/pkg/determine-delivery"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

type T *interface {
	Registry(c determine_delivery.D)
	Detect(string) (carriers.Carrier, error)
}

type Tracker struct {
	detector determine_delivery.D
}

func NewTracker(detector determine_delivery.D) *Tracker {
	return &Tracker{
		detector: detector,
	}
}

func (t *Tracker) Tracking(w http.ResponseWriter, r *http.Request) {

}
