package usps

import (
	"regexp"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
)

var patterns = map[string]*regexp.Regexp{
	//	9400100000000000000000 length 22
	"start94001": regexp.MustCompile(`(?i)^94001[\d]{17}$`),

	//	9205500000000000000000 length 22
	"start92055": regexp.MustCompile(`(?i)^92055[\d]{17}$`),

	//	9407300000000000000000 length 22
	"start94073": regexp.MustCompile(`(?i)^94073[\d]{17}$`),

	//	9303300000000000000000 length 22
	"start93033": regexp.MustCompile(`(?i)^93033[\d]{17}$`),

	//	9208800000000000000000 length 22
	"start92088": regexp.MustCompile(`(?i)^92088[\d]{17}$`),

	//	9202100000000000000000 length 22
	"start92021": regexp.MustCompile(`(?i)^92021[\d]{17}$`),

	//	9270100000000000000000 length 22
	"start92701": regexp.MustCompile(`(?i)^92701[\d]{17}$`),

	//	EC000000000US length 13
	"startEC_endUS": regexp.MustCompile(`(?i)^EC[\d]{9}US$`),

	//	EA000000000US length 13
	"startEA_endUS": regexp.MustCompile(`(?i)^EA[\d]{9}US$`),

	//	CP000000000US length 13
	"startCP_endUS": regexp.MustCompile(`(?i)^CP[\d]{9}US$`),

	//	8200000000 length 10
	//10 in UPS and DHL as well !!!
	"start80": regexp.MustCompile(`(?i)^82[\d]{8}$`),
}

type api interface {
	TrackByTrackingNumber(trackNumbers []TrackID) (*TrackResponse, error)
}

type Carrier struct {
	api api
}

func NewCarrier(api *Api) *Carrier {
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

func (c *Carrier) Track(trackingIds []string) ([]carriers.Parcel, error) {
	/*trackNumbers := []TrackID{TrackID{Id: trackNumber}}
	resp, err := c.api.TrackByTrackingNumber(trackNumbers)
	if err != nil {
		return nil, err
	}

	parcels := make([]carriers.Parcel, len(resp.TrackInfo))
	for i, info := range resp.TrackInfo {
		parcels[i] = carriers.Parcel{
			TrackingNumber: info.TrackId,
			Status:         info.Status,
		}
	}

	return parcels, nil*/
	return nil, nil
}
