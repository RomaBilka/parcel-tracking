package fedex

import (
	"regexp"
	"strings"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

var patterns = map[string]*regexp.Regexp{
	//Numeric only with the length 12
	"numbers12": regexp.MustCompile(`^[\d]{12}$`),
	//Numeric only with the length 15
	"numbers15": regexp.MustCompile(`^[\d]{15}$`),
	//Numeric only with the length 20
	"numbers20": regexp.MustCompile(`^[\d]{20}$`),
	//Numeric only with the length 22
	//22 in UPS as well !!!
	"numbers22": regexp.MustCompile(`^[\d]{22}$`),
}

type api interface {
	TrackByTrackingNumber(TrackingRequest) (*TrackingResponse, error)
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
	trackingInfo := TrackingInfo{
		TrackingNumberInfo: TrackingNumberInfo{
			TrackingNumber: trackingId,
		},
	}

	trackingData := TrackingRequest{IncludeDetailedScans: true}
	trackingData.TrackingInfo = append(trackingData.TrackingInfo, trackingInfo)

	response, err := c.api.TrackByTrackingNumber(trackingData)
	if err != nil {
		return nil, err
	}

	parcels := make([]carriers.Parcel, len(response.Output.CompleteTrackResults))
	for i, d := range response.Output.CompleteTrackResults {
		parcels[i] = carriers.Parcel{
			TrackingNumber: d.TrackingNumber,
			Places:         getPlaces(d.TrackResults[0]),
			Status:         d.TrackResults[0].LatestStatusDetail.StatusByLocale,
			DeliveryDate:   d.TrackResults[0].EstimatedDeliveryTimeWindow.Window.Ends,
		}
	}

	return parcels, nil
}

func getPlaces(result TrackResult) []carriers.Place {
	address := []carriers.Place{}

	if result.OriginLocation.LocationContactAndAddress.Address.City != "" {
		address = append(address, preparePlace(result.OriginLocation.LocationContactAndAddress.Address))
	}

	if result.DeliveryDetails.ActualDeliveryAddress.City != "" {
		address = append(address, preparePlace(result.DeliveryDetails.ActualDeliveryAddress))
	}

	if result.RecipientInformation.Address.City != "" {
		address = append(address, preparePlace(result.RecipientInformation.Address))
	}

	return address
}

func preparePlace(a Address) carriers.Place {
	return carriers.Place{
		Street:  strings.Join(a.StreetLines, " "),
		City:    a.City,
		Country: a.CountryName,
	}
}
