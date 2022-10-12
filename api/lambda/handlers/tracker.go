package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/aws/aws-lambda-go/events"
)

type Handler func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

type parcelTracker interface {
	TrackParcels(ctx context.Context, parcelIds []string) (map[string]carriers.Parcel, error)
}

func Tracking(t parcelTracker, maximumNumberTrackingId int) Handler {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

		fmt.Println(request.Body)
		/*
			id := request.QueryStringParameters["track_id"]

			if id == "" {
				return response(http.StatusBadRequest, api.Error{Message: "track_id cannot be empty"})
			}

			p, err := t.TrackParcel(ctx, id)
			if err != nil {
				return handleError(err)
			}
		*/
		s, _ := base64.StdEncoding.DecodeString(request.Body)
		return response(http.StatusOK, s)
	}
}

/*
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
*/
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
