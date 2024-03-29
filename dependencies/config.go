package dependencies

import (
	"github.com/jessevdk/go-flags"
)

type (
	NovaPoshta struct {
		ApiURL string `long:"NP_API_URL" description:"nova poshta API URL" required:"true" default:"https://api.novaposhta.ua" env:"NP_API_URL"`
		ApiKey string `long:"NP_API_KEY" description:"nova poshta API key" env:"NP_API_KEY"`
	}

	NovaPoshtaShopping struct {
		ApiURL string `long:"NP_SHOPPING_API_URL" description:"nova poshta shopping API URL" required:"true" default:"https://novaposhtaglobal.ua/ajax.php" env:"NP_SHOPPING_API_URL"`
	}

	MeestExpress struct {
		ApiURL   string `long:"ME_API_URL" description:"meest express API URL" required:"true" default:"https://apii.meest-group.com/T/1C_Query.php" env:"ME_API_URL"`
		ID       string `long:"ME_ID" description:"meest express ID" required:"true" default:"0xA79E003048D2B47311E26B7D4A430FFC" env:"ME_ID"`
		Login    string `long:"ME_LOGIN" description:"meest express login" required:"true" default:"public" env:"ME_LOGIN"`
		Password string `long:"ME_PASSWORD" description:"meest express password" required:"true" default:"PUBLIC" env:"ME_PASSWORD"`
	}

	DHL struct {
		ApiURL string `long:"DHL_URL" description:"DHL API URL" required:"true" default:"https://api-eu.dhl.com" env:"DHL_API_URL"`
		ApiKey string `long:"DHL_API_KEY" description:"DHL API key" required:"true" default:"demo-key" env:"DHL_API_KEY"`
	}

	Fedex struct {
		ApiURL          string `long:"FEDEX_URL" description:"FEDEX API URL" required:"true" default:"https://apis-sandbox.fedex.com" env:"FEDEX_API_URL"`
		ApiKey          string `long:"FEDEX_API_KEY" description:"Fedex client id" env:"FEDEX_API_KEY"`
		GrantType       string `long:"FEDEX_GRANT_TYPE" description:"Fedex grant type" env:"FEDEX_GRANT_TYPE"`
		ShippingAccount string `long:"FEDEX_SHIPPING_ACCOUNT" description:"Fedex shipping account" env:"FEDEX_SHIPPING_ACCOUNT"`
	}

	USPS struct {
		ApiURL   string `long:"USPS_APIV2_URL" description:"USPS express API TrackV2 URL" required:"true" default:"http://production.shippingapis.com/ShippingAPI.dll?API=TrackV2" env:"USPS_APIV2_URL"`
		UserID   string `long:"USPS_ID" description:"USPS user ID" env:"USPS_ID"`
		SourceId string `long:"USPS_SOURCE_ID" description:"USPS password" env:"USPS_SOURCE_ID"`
	}

	UPS struct {
		ApiURL              string `long:"UPS_API_URL" description:"UPS API URL" required:"true" default:"https://wwwcie.ups.com/ups.app/xml/" env:"UPS_API_URL"`
		UserID              string `long:"UPS_USER_ID" description:"UPS user ID"  env:"UPS_USER_ID"`
		AccessLicenseNumber string `long:"UPS_ACCESS_LICENSE_NUMBER" description:"UPS Access License Number" env:"UPS_ACCESS_LICENSE_NUMBER"`
		Password            string `long:"UPS_PASSWORD" description:"UPS password" env:"UPS_PASSWORD"`
	}

	UkrPoshta struct {
		ApiURL   string `long:"UKRPOSHTA_API_URL" description:"Ukrposhta API URL" required:"true" default:"https://www.ukrposhta.ua" env:"UKRPOSHTA_API_URL"`
		ApiToken string `long:"UKRPOSHTA_API_TOKEN" description:"Ukrposhta API token" env:"UKRPOSHTA_API_TOKEN"`
	}

	Config struct {
		Port                    string `short:"p" long:"port" description:"Port" required:"true" default:"8080" env:"PORT"`
		MaximumNumberTrackingId string `short:"n" long:"MAXIMUM_NUMBER_TRACKING_ID" description:"Maximum of number tracking Id" required:"true" default:"10" env:"MAXIMUM_NUMBER_TRACKING_ID"`
		NovaPoshta              NovaPoshta
		NovaPoshtaShopping      NovaPoshtaShopping
		MeestExpress            MeestExpress
		DHL                     DHL
		Fedex                   Fedex
		USPS                    USPS
		UPS                     UPS
		UkrPoshta               UkrPoshta
	}
)

func initConfig() (*Config, error) {
	cfg := &Config{}
	if _, err := flags.NewParser(cfg, flags.Default|flags.IgnoreUnknown).Parse(); err != nil {
		return nil, err
	}
	return cfg, nil
}
