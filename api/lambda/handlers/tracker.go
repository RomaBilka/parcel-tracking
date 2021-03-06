package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/RomaBilka/parcel-tracking/api"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/aws/aws-lambda-go/events"
)

type Handler func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

type parcelTracker interface {
	TrackParcel(ctx context.Context, parcelId string) (carriers.Parcel, error)
}

func Tracking(t parcelTracker) Handler {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		id := request.QueryStringParameters["track_id"]
		if id == "" {
			return response(http.StatusBadRequest, api.Error{Message: "track_id cannot be empty"})
		}

		p, err := t.TrackParcel(ctx, id)
		if err != nil {
			return response(http.StatusBadRequest, api.Error{Message: err.Error()})
		}
		return response(http.StatusOK, p)
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
