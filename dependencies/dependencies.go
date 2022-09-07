package dependencies

import (
	"github.com/RomaBilka/parcel-tracking/logic"
	determine_delivery "github.com/RomaBilka/parcel-tracking/pkg/determine-delivery"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/dhl"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/fedex"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/me"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/np"
	np_shopping "github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/np-shopping"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/ups"
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

	meApi := me.NewApi(config.MeestExpress.ApiURL, config.MeestExpress.ID, config.MeestExpress.Login, config.MeestExpress.Password)
	detector.Registry(me.NewCarrier(meApi))

	uspsApi := usps.NewApi(config.USPS.ApiURL, config.USPS.UserID, config.USPS.SourceId)
	detector.Registry(usps.NewCarrier(uspsApi))

	upsApi := ups.NewApi(config.UPS.ApiURL, config.UPS.UserID, config.UPS.AccessLicenseNumber, config.UPS.Password)
	detector.Registry(ups.NewCarrier(upsApi))

	fedexApi := fedex.NewApi(config.Fedex.ApiURL, config.Fedex.ApiKey, config.Fedex.GrantType, config.Fedex.ShippingAccount)
	detector.Registry(fedex.NewCarrier(fedexApi))

	dhlApi := dhl.NewApi(config.DHL.ApiURL, config.DHL.ApiKey)
	detector.Registry(dhl.NewCarrier(dhlApi))

	npShoppingApi := np_shopping.NewApi(config.NovaPoshtaShopping.ApiURL)
	detector.Registry(np_shopping.NewCarrier(npShoppingApi))

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
