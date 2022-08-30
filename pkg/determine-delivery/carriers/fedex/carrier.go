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
			Number:  d.TrackingNumber,
			Address: d.TrackResults[0].LatestStatusDetail.ScanLocation.CountryName + " " + d.TrackResults[0].LatestStatusDetail.ScanLocation.City,
			Status:  d.TrackResults[0].LatestStatusDetail.StatusByLocale,
		}
	}

	return parcels, nil
}

func (c *Carrier) Track_draft(trackingId string) ([]carriers.Parcel_draft, error) {
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

	parcels := make([]carriers.Parcel_draft, len(response.Output.CompleteTrackResults))
	for i, d := range response.Output.CompleteTrackResults {
		parcels[i] = carriers.Parcel_draft{
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

	return append(address,
		preparePlace(result.OriginLocation.LocationContactAndAddress.Address),
		preparePlace(result.DeliveryDetails.ActualDeliveryAddress),
		preparePlace(result.RecipientInformation.Address))
}

func preparePlace(a Address) carriers.Place {
	return carriers.Place{
		Street: strings.Join(a.StreetLines, " "),
		City:   a.City,
		County: a.CountryName,
	}
}
