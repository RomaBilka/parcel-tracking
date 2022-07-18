package dhl

type request struct {
	TrackingNumber       string `url:"trackingNumber"`
	Service              string `url:"service,omitempty"`
	OriginCountryCode    string `url:"originCountryCode,omitempty"`
	RequesterCountryCode string `url:"requesterCountryCode,omitempty"`
	RecipientPostalCode  string `url:"recipientPostalCode,omitempty"`
	Language             string `url:"language,omitempty"`
	Offset               string `url:"offset,omitempty"`
	Limit                string `url:"limit,omitempty"`
}
