package fedex

type authorizeRequest struct {
	GrantType    string `url:"grant_type"`
	ClientId     string `url:"client_id"`
	ClientSecret string `url:"client_secret"`
}

type TrackingRequest struct {
	IncludeDetailedScans bool           `json:"includeDetailedScans"`
	TrackingInfo         []TrackingInfo `json:"trackingInfo"`
}

type TrackingInfo struct {
	ShipDateBegin      string             `json:"shipDateBegin,omitempty"`
	ShipDateEnd        string             `json:"shipDateEnd,omitempty"`
	TrackingNumberInfo TrackingNumberInfo `json:"trackingNumberInfo"`
}

type TrackingNumberInfo struct {
	TrackingNumber         string `json:"trackingNumber"`
	CarrierCode            string `json:"carrierCode,omitempty"`
	TrackingNumberUniqueId string `json:"trackingNumberUniqueId,omitempty"`
}
