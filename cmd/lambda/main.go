package main

import (
	"github.com/RomaBilka/parcel-tracking/api/lambda/handlers"
	"github.com/RomaBilka/parcel-tracking/dependencies"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	deps, err := dependencies.InitDeps()
	if err != nil {
		panic(err)
	}
	lambda.Start(handlers.Tracking(deps.ParcelTracker))
}
