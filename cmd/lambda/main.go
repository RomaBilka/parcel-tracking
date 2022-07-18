package main

import (
	api "github.com/RomaBilka/parcel-tracking/api/lambda"
	"github.com/RomaBilka/parcel-tracking/dependencies"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	deps, err := dependencies.InitDeps()
	if err != nil {
		panic(err)
	}
	defer deps.TearDown()

	h := api.Configure(deps)
	lambda.Start(h)
}
