package dependencies

import (
	"github.com/RomaBilka/parcel-tracking/logic"
	determine_delivery "github.com/RomaBilka/parcel-tracking/pkg/determine-delivery"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/dhl"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/me"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/np"
)

type Deps struct {
	Config        *Config
	ParcelTracker *logic.ParcelsTracker
}

func InitDeps() (*Deps, error) {
	config, err := initConfig()
	if err != nil {
		return nil, err
	}

	detector := determine_delivery.NewDetector()
	detector.Registry(np.NewCarrier(np.NewApi(config.NovaPoshta.ApiURL, config.NovaPoshta.ApiKey)))

	meApi := me.NewApi(config.MeestExpress.ID, config.MeestExpress.Login, config.MeestExpress.Password, config.MeestExpress.ApiURL)
	detector.Registry(me.NewCarrier(meApi))

	dhlApi := dhl.NewApi(config.DHL.ApiURL, config.DHL.ApiKey)
	detector.Registry(dhl.NewCarrier(dhlApi))

	return &Deps{
		Config:        config,
		ParcelTracker: logic.NewParcelsTracker(detector),
	}, nil
}