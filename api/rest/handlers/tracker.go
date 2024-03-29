package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/RomaBilka/parcel-tracking/api"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	response_errors "github.com/RomaBilka/parcel-tracking/pkg/response-errors"
)

type parcelTracker interface {
	TrackParcels(ctx context.Context, parcelIds []string) (map[string]carriers.Parcel, error)
}

type trackData struct {
	Ids []string `json:"track_id"`
}

func Tracking(t parcelTracker, maximumNumberTrackingId int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		data := &trackData{}
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		if len(data.Ids) < 1 {
			writeErrorResponse(w, http.StatusBadRequest, errors.New("track_id cannot be empty"))
			return
		}

		for _, id := range data.Ids {
			if id == "" {
				writeErrorResponse(w, http.StatusBadRequest, errors.New("track_id cannot be empty"))
				return
			}
		}

		if len(data.Ids) > maximumNumberTrackingId {
			writeErrorResponse(w, http.StatusBadRequest, errors.New("too many track numbers"))
			return
		}

		ctx := r.Context()
		parcels, err := t.TrackParcels(ctx, data.Ids)
		if err != nil {
			handleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(parcels); err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err)
		}
	}
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, response_errors.NotFound):
		writeErrorResponse(w, http.StatusNotFound, err)
	case errors.Is(err, response_errors.CarrierNotFound):
		writeErrorResponse(w, http.StatusNotFound, err)
	case errors.Is(err, response_errors.InvalidNumber):
		writeErrorResponse(w, http.StatusBadRequest, err)
	default:
		writeErrorResponse(w, http.StatusInternalServerError, err)
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
