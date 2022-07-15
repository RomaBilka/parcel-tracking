package lambda

import (
	"github.com/RomaBilka/parcel-tracking/api/lambda/handlers"
	"github.com/RomaBilka/parcel-tracking/api/lambda/midllewares"
	"github.com/RomaBilka/parcel-tracking/dependencies"
)

func Configure(deps *dependencies.Deps) handlers.Handler {
	tracking := handlers.Tracking(deps.ParcelTracker)
	logging := midllewares.Logging(deps.Logger)
	return midllewares.RewriteInternalErrors(logging(tracking))
}
