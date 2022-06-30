package cmd

import (
	determine_delivery "github.com/RomaBilka/parcel-tracking/pkg/determine-delivery"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/me"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/np"
	"github.com/jessevdk/go-flags"
)

type opts struct {
	Port       string `short:"p" long:"port" description:"Port" required:"true" default:"8080" env:"PORT"`
	NP_API_URL string `long:"NP_API_URL" description:"nova poshta API URL" required:"true" default:"https://api.novaposhta.ua" env:"NP_API_URL"`
	NP_API_Key string `long:"NP_API_Key" description:"nova poshta API key"  default:"" env:"NP_API_KEY"`

	ME_API_URL  string `long:"ME_API_URL" description:"meest express API URL" required:"true" default:"https://apii.meest-group.com/T/1C_Query.php" env:"ME_API_URL"`
	ME_Login    string `long:"ME_Login" description:"meest express login" required:"true" default:"public" env:"ME_LOGIN"`
	ME_Password string `long:"ME_Password" description:"meest express password" required:"true" default:"PUBLIC" env:"ME_PASSWORD"`
	ME_ID       string `long:"ME_ID" description:"meest express ID" required:"true" default:"0xA79E003048D2B47311E26B7D4A430FFC" env:"ME_ID"`
}

var O opts

func init() {
	_, err := flags.Parse(&O)
	if err != nil {
		panic(err)
	}
}
func GetDetector() *determine_delivery.Detector {
	detector := determine_delivery.NewDetector()
	detector.Registry(np.NewCarrier(np.NewApi(O.NP_API_URL, O.NP_API_Key)))
	detector.Registry(me.NewCarrier(me.NewApi(O.ME_ID, O.ME_Login, O.ME_Password, O.ME_API_URL)))

	return detector
}
