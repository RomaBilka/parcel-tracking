package fedex

import (
	"regexp"

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

	documents, err := c.api.TrackByTrackingNumber(trackingData)
	if err != nil {
		return nil, err
	}

	parcels := make([]carriers.Parcel, len(documents.Output.CompleteTrackResults))
	for i, d := range documents.Output.CompleteTrackResults {
		parcels[i] = carriers.Parcel{
			Number:  d.TrackingNumber,
			Address: d.TrackResults[0].LatestStatusDetail.ScanLocation.CountryName + " " + d.TrackResults[0].LatestStatusDetail.ScanLocation.City,
			Status:  d.TrackResults[0].LatestStatusDetail.StatusByLocale,
		}
	}

	return parcels, nil
}
