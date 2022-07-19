package fedex

type Request struct {
	IncludeDetailedScans bool           `json:"includeDetailedScans"`
	TrackingInfo         []TrackingInfo `json:"trackingInfo"`
}

type TrackingInfo struct {
	ShipDateBegin      string             `json:"shipDateBegin"`
	ShipDateEnd        string             `json:"shipDateEnd"`
	TrackingNumberInfo TrackingNumberInfo `json:"trackingNumberInfo"`
}

type TrackingNumberInfo struct {
	TrackingNumber         string `json:"trackingNumber"`
	CarrierCode            string `json:"carrierCode"`
	TrackingNumberUniqueId string `json:"trackingNumberUniqueId"`
}
