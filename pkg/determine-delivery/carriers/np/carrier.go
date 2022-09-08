package np

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/RomaBilka/parcel-tracking/pkg/helpers"
)

var patterns = map[string]*regexp.Regexp{
	//Starts with 59, length 14, only numbers
	//59************
	"start59": regexp.MustCompile(`^59[\d]{12}$`),

	//Starts with 20, length 14, only numbers
	//20************
	"start20": regexp.MustCompile(`^20[\d]{12}$`),

	//Starts with 1, length 14, only numbers
	//1*************
	"start1": regexp.MustCompile(`^1[\d]{13}$`),
}

const layout = "08.03.2022 15:16:47"

type api interface {
	TrackByTrackingNumber(TrackingDocuments) (*TrackingDocumentsResponse, error)
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

	response, err := c.api.TrackByTrackingNumber(methodProperties)
	if err != nil {
		return nil, err
	}

	parcels := make([]carriers.Parcel, len(response.Data))

	for i, d := range response.Data {

		scheduledDeliveryDate, err := helpers.ParseTime(layout, d.ScheduledDeliveryDate)
		if err != nil {
			return nil, err
		}

		recipientDate, err := helpers.ParseTime(layout, d.RecipientDateTime)
		if err != nil {
			return nil, err
		}

		asctualDeliveryDate, err := helpers.ParseTime(layout, d.ActualDeliveryDate)
		if err != nil {
			return nil, err
		}

		sender := carriers.Place{
			City:    d.CitySender,
			Address: d.WarehouseSender,
			Date:    recipientDate,
		}

		recipient := carriers.Place{
			City:    d.CityRecipient,
			Address: d.WarehouseRecipient,
			Date:    scheduledDeliveryDate,
		}

		parcels[i] = carriers.Parcel{
			TrackingNumber: d.Number,
			Places: []carriers.Place{
				sender,
				recipient,
			},
			Status:       d.Status,
			DeliveryDate: asctualDeliveryDate,
		}
	}

	return parcels, nil
}
