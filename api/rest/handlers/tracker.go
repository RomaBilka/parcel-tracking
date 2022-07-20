package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/RomaBilka/parcel-tracking/api"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

type parcelTracker interface {
	TrackParcel(ctx context.Context, parcelId string) (carriers.Parcel, error)
}

func Tracking(t parcelTracker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		trackingId := r.URL.Query().Get("track_id")
		if trackingId == "" {
			writeErrorResponse(w, http.StatusBadRequest, errors.New("track_id cannot be empty"))
			return
		}

		ctx := r.Context()
		parcel, err := t.TrackParcel(ctx, trackingId)
		if err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(parcel); err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err)
		}
	}
}

func writeErrorResponse(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	b, err := json.Marshal(api.Error{Message: err.Error()})
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		panic(err)
	}
}
