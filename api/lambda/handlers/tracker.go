package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/RomaBilka/parcel-tracking/api"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	response_errors "github.com/RomaBilka/parcel-tracking/pkg/response-errors"
	"github.com/aws/aws-lambda-go/events"
)

type Handler func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

type parcelTracker interface {
	TrackParcel(ctx context.Context, parcelId string) (carriers.Parcel, error)
}

func Tracking(t parcelTracker) Handler {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

		fmt.Println(request.Body)

		id := request.QueryStringParameters["track_id"]

		if id == "" {
			return response(http.StatusBadRequest, api.Error{Message: "track_id cannot be empty"})
		}

		p, err := t.TrackParcel(ctx, id)
		if err != nil {
			return handleError(err)
		}
		return response(http.StatusOK, p)
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
