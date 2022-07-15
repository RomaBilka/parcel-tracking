package np

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

var patterns = map[string]*regexp.Regexp{
	// 59************
	"start59": regexp.MustCompile(`^59[\d]{12}$`),

	"start20": regexp.MustCompile(`^20[\d]{12}$`),

	"start1": regexp.MustCompile(`^1[\d]{13}$`),
}

type api interface {
	TrackingDocument(TrackingDocuments) (*TrackingDocumentsResponse, error)
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
	for _, pattern := range patterns {
		if pattern.MatchString(trackId) {
			return true
		}
	}

	return false
}

func (c *Carrier) Track(trackingId string) ([]carriers.Parcel, error) {
	document := TrackingDocument{
		DocumentNumber: trackingId,
	}
	methodProperties := TrackingDocuments{}
	methodProperties.Documents = append(methodProperties.Documents, document)
	methodProperties.CheckWeightMethod = "3"

	documents, err := c.api.TrackingDocument(methodProperties)
	if err != nil {
		return nil, err
	}

	parcels := make([]carriers.Parcel, len(documents.Data))
	for i, d := range documents.Data {
		parcels[i] = carriers.Parcel{
			Number:  d.Number,
			Address: d.CityRecipient + " " + d.WarehouseRecipient,
			Status:  d.Status,
		}
	}

	return parcels, nil
}
