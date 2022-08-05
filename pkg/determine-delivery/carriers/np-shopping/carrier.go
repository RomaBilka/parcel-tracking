package np_shopping

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

////Starts with NP, 14 numbers and NPG at the end NP99999999999999NPG
var npShopping = regexp.MustCompile(`(?i)^NPI[\d]{14}$`)

type api interface {
	TrackingDocument(string) (*TrackingDocumentResponse, error)
}

type Carrier struct {
	api api
}

func NewCarrier(api api) *Carrier {
	return &Carrier{
		api: api,
	}
}

func (c *Carrier) Detect(trackId string) bool {
	return npShopping.MatchString(trackId)
}

func (c *Carrier) Track(trackId string) ([]carriers.Parcel, error) {
	doc, err := c.api.TrackingDocument(trackId)
	if err != nil {
		return nil, err
	}

	p := carriers.Parcel{
		Number:  doc.WaybillNumber,
		Address: doc.PickupAddress.Country,
		Status:  doc.State,
	}

	return []carriers.Parcel{p}, nil
}
