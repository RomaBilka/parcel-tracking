package determine_delivery

import (
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/me"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/np"
	np_shopping "github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/np-shopping"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/ups"
)

const NOVA_POSHTA = "NovaPoshta"
const NP_SHOPPING = "NPShopping"
const UPS = "UPS"
const MEEST_EXPRESS = "MeestExpress"

//var carriers map[string]D

type D interface {
	Detect() bool
}

func Detect(str string) string {
	carriers := make(map[string]D)

	carriers[NOVA_POSHTA] = &np.NP{str}
	carriers[NP_SHOPPING] = &np_shopping.NPShopping{str}
	carriers[UPS] = &ups.UPS{str}
	carriers[MEEST_EXPRESS] = &me.ME{str}

	for k, carrier := range carriers {
		if carrier.Detect() {
			return k
		}
	}

	return ""
}
