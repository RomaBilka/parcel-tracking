package rest

import (
	"net/http"

	"github.com/RomaBilka/parcel-tracking/api/rest/handlers"
	"github.com/RomaBilka/parcel-tracking/api/rest/midllewares"
	"github.com/RomaBilka/parcel-tracking/dependencies"
)

func Configure(deps *dependencies.Deps) {
	http.Handle("/tracking",
		midllewares.RewriteInternalErrors(
			midllewares.Logging(deps.Logger)(
				handlers.Tracking(deps.ParcelTracker),
			),
		),
	)

}
