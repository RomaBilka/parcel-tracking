package np

import "github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"

type Carrier struct {
	api *Api
}

func NewCarrier(api *Api) *Carrier {
	return &Carrier{
		api: api,
	}
}

func (c *Carrier) Tracking(trackingId string) ([]carriers.Parcel, error) {
	parcels := []carriers.Parcel{}

	document := TrackingDocument{
		DocumentNumber: trackingId,
	}
	methodProperties := TrackingDocuments{}
	methodProperties.Documents = append(methodProperties.Documents, document)
	methodProperties.CheckWeightMethod = "3"

	documents, err := c.api.TrackingDocument(methodProperties)
	if err != nil {
		return parcels, err
	}

	for _, d := range documents.Data {
		parcels = append(parcels, carriers.Parcel{
			Number:  d.Number,
			Address: d.CityRecipient,
		})
	}

	return parcels, nil
}
