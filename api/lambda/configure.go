package lambda

import (
	"strconv"

	"github.com/RomaBilka/parcel-tracking/api/lambda/handlers"
	"github.com/RomaBilka/parcel-tracking/api/lambda/midllewares"
	"github.com/RomaBilka/parcel-tracking/dependencies"
)

func Configure(deps *dependencies.Deps) handlers.Handler {
	n, err := strconv.Atoi(deps.Config.MaximumNumberTrackingId)
	if err != nil {
		panic(err)
	}
	tracking := handlers.Tracking(deps.ParcelTracker, n)
	logging := midllewares.Logging(deps.Logger)
	panicRecovery := midllewares.PanicRecovery(deps.Logger.Sugar())

	return midllewares.RewriteInternalErrors(panicRecovery(logging(tracking)))
}
