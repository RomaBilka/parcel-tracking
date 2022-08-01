package dependencies

import (
	"github.com/RomaBilka/parcel-tracking/logic"
	determine_delivery "github.com/RomaBilka/parcel-tracking/pkg/determine-delivery"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/dhl"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/me"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/np"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/usps"
	"go.uber.org/zap"
)

type Deps struct {
	Config        *Config
	ParcelTracker *logic.ParcelsTracker
	Logger        *zap.Logger
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

	uspsApi := usps.NewApi(config.USPS.UserID, config.USPS.Password, config.USPS.URL)
	detector.Registry(usps.NewCarrier(uspsApi))

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &Deps{
		Config:        config,
		ParcelTracker: logic.NewParcelsTracker(detector),
		Logger:        logger,
	}, nil
}

func (d Deps) TearDown() {
	_ = d.Logger.Sync()
}
