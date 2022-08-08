package ups

type AccessRequest struct {
	AccessLicenseNumber string `xml:"AccessLicenseNumber"`
	UserId              string `xml:"UserId"`
	Password            string `xml:"Password"`
}

type TrackRequest struct {
	TrackingNumber                     string `xml:"TrackingNumber"`
	ShipmentIdentificationNumber       string `xml:"ShipmentIdentificationNumber,omitempty"`
	CandidateBookmark                  string `xml:"CandidateBookmark,omitempty"`
	ShipperNumber                      string `xml:"ShipperNumber,omitempty"`
	DestinationPostalCode              string `xml:"DestinationPostalCode,omitempty"`
	DestinationCountryCode             string `xml:"DestinationCountryCode,omitempty"`
	OriginPostalCode                   string `xml:"OriginPostalCode,omitempty"`
	OriginCountryCode                  string `xml:"OriginCountryCode,omitempty"`
	IncludeMailInnovationIndicator     string `xml:"IncludeMailInnovationIndicator,omitempty"`
	TrackingOption                     string `xml:"TrackingOption,omitempty"`
	UPSWorldWideExpressFreightShipment string `xml:"UPSWorldWideExpressFreightShipment,omitempty"`
	IncludeFreight                     string `xml:"IncludeFreight,omitempty"`
	PreauthorizedReturnIndicator       string `xml:"PreauthorizedReturnIndicator,omitempty"`
	//Request                            Request            `xml:"Request,omitempty"`
	//ShipperAccountInfo                 ShipperAccountInfo `xml:"ShipperAccountInfo,omitempty"`
	//ShipmentType                       ShipmentType       `xml:"ShipmentType,omitempty"`
	//ReferenceNumber                    ReferenceNumber    `xml:"ReferenceNumber,omitempty"`
	//PickupDateRange                    PickupDateRange    `xml:"PickupDateRange,omitempty"`
}

type Request struct {
	RequestAction        string               `xml:"RequestAction,omitempty"`
	RequestOption        string               `xml:"RequestOption,omitempty"`
	SubVersion           string               `xml:"SubVersion,omitempty"`
	TransactionReference TransactionReference `xml:"TransactionReference,omitempty"`
}

type TransactionReference struct {
	CustomerContext       string `xml:"CustomerContext,omitempty"`
	TransactionIdentifier string `xml:"TransactionIdentifier,omitempty"`
	ToolVersion           string `xml:"ToolVersion,omitempty"`
}

type ShipperAccountInfo struct {
	PostalCode  string `xml:"PostalCode,omitempty"`
	CountryCode string `xml:"CountryCode,omitempty"`
}

type ShipmentType struct {
	Code        string `xml:"Code,omitempty"`
	Description string `xml:"Description,omitempty"`
}

type PickupDateRange struct {
	BeginDate string `xml:"BeginDate,omitempty"`
	EndDate   string `xml:"EndDate,omitempty"`
}

type ReferenceNumber struct {
	Value string `xml:"Value,omitempty"`
}
