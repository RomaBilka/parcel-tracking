package dependencies

import (
	"github.com/jessevdk/go-flags"
)

type NovaPoshta struct {
	ApiURL string `long:"NP_API_URL" description:"nova poshta API URL" required:"true" default:"https://api.novaposhta.ua" env:"NP_API_URL"`
	ApiKey string `long:"NP_API_Key" description:"nova poshta API key"  default:"" env:"NP_API_KEY"`
}

type MeestExpress struct {
	ID       string `long:"ME_ID" description:"meest express ID" required:"true" default:"0xA79E003048D2B47311E26B7D4A430FFC" env:"ME_ID"`
	Login    string `long:"ME_Login" description:"meest express login" required:"true" default:"public" env:"ME_LOGIN"`
	Password string `long:"ME_Password" description:"meest express password" required:"true" default:"PUBLIC" env:"ME_PASSWORD"`
	ApiURL   string `long:"ME_API_URL" description:"meest express API URL" required:"true" default:"https://apii.meest-group.com/T/1C_Query.php" env:"ME_API_URL"`
}

type DHL struct {
	ApiURL string `long:"DHL_URL" description:"DHL API URL" required:"true" default:"https://api-eu.dhl.com" env:"DHL_API_URL"`
	ApiKey string `long:"DHL_API_Key" description:"DHL API key"  default:"demo-key" env:"DHL_API_KEY"`
}

type USPS struct {
	UserID   string `long:"USPS_ID" description:"USPS user ID" required:"true" default:"302ROMAN1277" env:"USPS_ID"`
	Password string `long:"USPS_Password" description:"USPS password" required:"false" default:"536KD44UF932" env:"USPS_Password"`
	URL      string `long:"USPS_APIV2_URL" description:"USPS express API TrackV2 URL" required:"true" default:"http://production.shippingapis.com/ShippingAPI.dll?API=TrackV2" env:"USPS_APIV2_URL"`
	ApiURL   string `long:"ME_API_URL" description:"meest express API URL" required:"true" default:"https://apii.meest-group.com/T/1C_Query.php" env:"ME_API_URL"`
}

type Config struct {
	Port         string `short:"p" long:"port" description:"Port" required:"true" default:"8080" env:"PORT"`
	NovaPoshta   NovaPoshta
	MeestExpress MeestExpress
	USPS         USPS
	DHL          DHL
}

func initConfig() (*Config, error) {
	cfg := &Config{}
	if _, err := flags.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
