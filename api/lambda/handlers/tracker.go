package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/RomaBilka/parcel-tracking/api"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	response_errors "github.com/RomaBilka/parcel-tracking/pkg/response-errors"
	"github.com/aws/aws-lambda-go/events"
)

type Handler func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

type parcelTracker interface {
	TrackParcels(ctx context.Context, parcelIds []string) (map[string]carriers.Parcel, error)
}

type trackData struct {
	Ids []string `json:"track_id"`
}

func Tracking(t parcelTracker, maximumNumberTrackingId int) Handler {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		data := &trackData{}
		if err := json.Unmarshal([]byte(request.Body), data); err != nil {
			return handleError(err)
		}

		if len(data.Ids) < 1 {
			return response(http.StatusBadRequest, api.Error{Message: "track_id cannot be empty"})
		}

		for _, id := range data.Ids {
			if id == "" {
				return response(http.StatusBadRequest, api.Error{Message: "track_id cannot be empty"})
			}
		}

		if len(data.Ids) > maximumNumberTrackingId {
			return response(http.StatusBadRequest, api.Error{Message: "too many track numbers"})
		}

		parcels, err := t.TrackParcels(ctx, data.Ids)
		if err != nil {
			return handleError(err)
		}

		return response(http.StatusOK, parcels)
	}
}

func handleError(err error) (events.APIGatewayProxyResponse, error) {
	switch {
	case errors.Is(err, response_errors.NotFound):
		return response(http.StatusNotFound, api.Error{Message: err.Error()})
	case errors.Is(err, response_errors.CarrierNotFound):
		return response(http.StatusNotFound, api.Error{Message: err.Error()})
	case errors.Is(err, response_errors.InvalidNumber):
		return response(http.StatusBadRequest, api.Error{Message: err.Error()})
	default:
		return response(http.StatusInternalServerError, api.Error{Message: err.Error()})
	}
}

func response(status int, body interface{}) (events.APIGatewayProxyResponse, error) {
	resp, err := json.Marshal(body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(resp),
	}, nil
}
