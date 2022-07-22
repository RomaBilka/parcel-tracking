package rest

import (
	"net/http"

	"github.com/RomaBilka/parcel-tracking/api/rest/handlers"
	"github.com/RomaBilka/parcel-tracking/api/rest/midllewares"
	"github.com/RomaBilka/parcel-tracking/dependencies"
)

func Configure(deps *dependencies.Deps) {
	tracking := handlers.Tracking(deps.ParcelTracker)
	logging := midllewares.Logging(deps.Logger)
	panicRecovery := midllewares.PanicRecovery(deps.Logger.Sugar())

	http.Handle("/tracking", midllewares.RewriteInternalErrors(panicRecovery(logging(tracking))))
}
