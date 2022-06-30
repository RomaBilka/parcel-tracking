package main

import (
	"context"

	"github.com/RomaBilka/parcel-tracking/cmd"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleLambdaEvent(ctx context.Context, request events.APIGatewayProxyRequest) ([]carriers.Parcel, error) {
	detector := cmd.GetDetector()

	carrier, err := detector.Detect(request.QueryStringParameters["track_id"])

	if err != nil {
		return []carriers.Parcel{}, err
	}

	parcel, err := carrier.Track(request.QueryStringParameters["track_id"])
	if err != nil {
		return []carriers.Parcel{}, err
	}

	return parcel, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
