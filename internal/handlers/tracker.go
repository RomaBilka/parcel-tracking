package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

type Detector interface {
	Registry(carrier carriers.Carrier)
	Detect(string) (carriers.Carrier, error)
}

type Tracker struct {
	detector Detector
}

func NewTracker(detector Detector) *Tracker {
	return &Tracker{
		detector: detector,
	}
}

func (t *Tracker) Tracking(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	trackingId, ok := r.URL.Query()["tracking_id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", "tracking_id is empty")
		return
	}

	carrier, err := t.detector.Detect(trackingId[0])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	parcel, err := carrier.Track(trackingId[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(parcel)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
	}
}
