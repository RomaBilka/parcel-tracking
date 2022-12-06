package rest

import (
	"net/http"
	"strconv"

	"github.com/RomaBilka/parcel-tracking/api/rest/handlers"
	"github.com/RomaBilka/parcel-tracking/api/rest/midllewares"
	"github.com/RomaBilka/parcel-tracking/dependencies"
)

func Configure(deps *dependencies.Deps) {
	n, err := strconv.Atoi(deps.Config.MaximumNumberTrackingId)
	if err != nil {
		panic(err)
	}
	tracking := handlers.Tracking(deps.ParcelTracker, n)

	logging := midllewares.Logging(deps.Logger)
	panicRecovery := midllewares.PanicRecovery(deps.Logger.Sugar())

	http.Handle("/tracking", midllewares.RewriteInternalErrors(panicRecovery(logging(tracking))))
}
