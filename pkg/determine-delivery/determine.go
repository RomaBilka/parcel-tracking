package determine_delivery

import (
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/deliveries"
)

const NOVA_POSHTA = "NovaPoshta"
const NP_SHOPPING = "NPShopping"
const UPS = "UPS"

func Determine(str string) (string, error) {
	ok, err := deliveries.IsNovaPoshta(str)
	if err != nil {
		return "", err
	} else if ok {
		return NOVA_POSHTA, nil
	}

	ok, err = deliveries.IsNpShopping(str)
	if err != nil {
		return "", err
	} else if ok {
		return NP_SHOPPING, nil
	}

	ok, err = deliveries.IsUPS(str)
	if err != nil {
		return "", err
	} else if ok {
		return UPS, nil
	}

	return "", nil
}
