package fedex

import (
	"encoding/json"
	"errors"
	"time"
)

type authResponse struct {
	AccessToken string  `json:"access_token"`
	TokenType   string  `json:"token_type"`
	ExpiresIn   Expires `json:"expires_in"`
	Scope       string  `json:"scope"`
}

type Expires struct {
	time.Time
}

func (e *Expires) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		seconds := time.Duration(value) * time.Second
		e.Time = time.Now().Local().Add(seconds)
		return nil
	default:
		return errors.New("invalid duration")
	}
}

type Error struct {
	Code          string          `json:"code"`
	Message       string          `json:"message"`
	ParameterList []ParameterList `json:"parameterList"`
}

type ParameterList struct {
	Value string `json:"value"`
	Key   string `json:"key"`
}

type Response struct {
	TransactionId         string  `json:"transactionId"`
	CustomerTransactionId string  `json:"customerTransactionId"`
	Errors                []Error `json:"errors"`
}

type TrackingResponse struct {
	Response
	Output Output `json:"output"`
}

type Output struct {
	CompleteTrackResults []CompleteTrackResult `json:"completeTrackResults"`
	Alerts               string                `json:"alerts"`
}

type CompleteTrackResult struct {
	TrackingNumber string        `json:"trackingNumber"`
	TrackResults   []TrackResult `json:"trackResults"`
}
type TrackResult struct {
	TrackingNumberInfo struct {
		TrackingNumber         string `json:"trackingNumber"`
		CarrierCode            string `json:"carrierCode"`
		TrackingNumberUniqueId string `json:"trackingNumberUniqueId"`
	} `json:"trackingNumberInfo"`
	AdditionalTrackingInfo        AdditionalTrackingInfo `json:"additionalTrackingInfo"`
	DistanceToDestination         DistanceToDestination  `json:"distanceToDestination"`
	ConsolidationDetail           []ConsolidationDetail  `json:"consolidationDetail"`
	MeterNumber                   string                 `json:"meterNumber"`
	ReturnDetail                  ReturnDetail           `json:"returnDetail"`
	ServiceDetail                 ServiceDetail          `json:"serviceDetail"`
	DestinationLocation           DestinationLocation    `json:"destinationLocation"`
	LatestStatusDetail            LatestStatusDetail     `json:"latestStatusDetail"`
	ServiceCommitMessage          ServiceCommitMessage   `json:"serviceCommitMessage"`
	InformationNotes              []InformationNote      `json:"informationNotes"`
	Error                         Error                  `json:"error"`
	SpecialHandlings              []SpecialHandling      `json:"specialHandlings"`
	AvailableImages               []AvailableImage       `json:"availableImages"`
	DeliveryDetails               DeliveryDetails        `json:"deliveryDetails"`
	ScanEvents                    []ScanEvent            `json:"scanEvents"`
	DateAndTimes                  []DateAndTime          `json:"dateAndTimes"`
	PackageDetails                PackageDetails         `json:"packageDetails"`
	GoodsClassificationCode       string                 `json:"goodsClassificationCode"`
	HoldAtLocation                HoldAtLocation         `json:"holdAtLocation"`
	CustomDeliveryOptions         []CustomDeliveryOption `json:"customDeliveryOptions"`
	EstimatedDeliveryTimeWindow   DescriptionWindow      `json:"estimatedDeliveryTimeWindow"`
	PieceCounts                   []PieceCount           `json:"pieceCounts"`
	OriginLocation                OriginLocation         `json:"originLocation"`
	RecipientInformation          ContactAndAddress      `json:"recipientInformation"`
	StandardTransitTimeWindow     DescriptionWindow      `json:"standardTransitTimeWindow"`
	ShipmentDetails               ShipmentDetails        `json:"shipmentDetails"`
	ReasonDetail                  ReasonDetail           `json:"reasonDetail"`
	AvailableNotifications        []string               `json:"availableNotifications"`
	ShipperInformation            ContactAndAddress      `json:"shipperInformation"`
	LastUpdatedDestinationAddress Address                `json:"lastUpdatedDestinationAddress"`
}

type AdditionalTrackingInfo struct {
	HasAssociatedShipments bool                `json:"hasAssociatedShipments"`
	Nickname               string              `json:"nickname"`
	PackageIdentifiers     []PackageIdentifier `json:"packageIdentifiers"`
	ShipmentNotes          string              `json:"shipmentNotes"`
}

type PackageIdentifier struct {
	Type                   string `json:"type"`
	Value                  string `json:"value"`
	TrackingNumberUniqueId string `json:"trackingNumberUniqueId"`
}

type DistanceToDestination struct {
	Units string  `json:"units"`
	Value float64 `json:"value"`
}

type ConsolidationDetail struct {
	TimeStamp       time.Time    `json:"timeStamp"`
	ConsolidationID string       `json:"consolidationID"`
	ReasonDetail    ReasonDetail `json:"reasonDetail"`
	PackageCount    int          `json:"packageCount"`
	EventType       string       `json:"eventType"`
}

type ReasonDetail struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

type ReturnDetail struct {
	AuthorizationName string         `json:"authorizationName"`
	ReasonDetail      []ReasonDetail `json:"reasonDetail"`
}

type ServiceDetail struct {
	Description      string `json:"description"`
	ShortDescription string `json:"shortDescription"`
	Type             string `json:"type"`
}

type DestinationLocation struct {
	LocationId                string            `json:"locationId"`
	LocationContactAndAddress ContactAndAddress `json:"locationContactAndAddress"`
	LocationType              string            `json:"locationType"`
}

type Contact struct {
	PersonName  string `json:"personName"`
	PhoneNumber string `json:"phoneNumber"`
	CompanyName string `json:"companyName"`
}

type Address struct {
	Classification      string   `json:"classification"`
	Residential         bool     `json:"residential"`
	StreetLines         []string `json:"streetLines"`
	City                string   `json:"city"`
	UrbanizationCode    string   `json:"urbanizationCode"`
	StateOrProvinceCode string   `json:"stateOrProvinceCode"`
	PostalCode          string   `json:"postalCode"`
	CountryCode         string   `json:"countryCode"`
	CountryName         string   `json:"countryName"`
}

type LatestStatusDetail struct {
	ScanLocation     ScanLocation      `json:"scanLocation"`
	Code             string            `json:"code"`
	DerivedCode      string            `json:"derivedCode"`
	AncillaryDetails []AncillaryDetail `json:"ancillaryDetails"`
	StatusByLocale   string            `json:"statusByLocale"`
	Description      string            `json:"description"`
	DelayDetail      DelayDetail       `json:"delayDetail"`
}

type ScanLocation struct {
	Classification      string   `json:"classification"`
	Residential         bool     `json:"residential"`
	StreetLines         []string `json:"streetLines"`
	City                string   `json:"city"`
	UrbanizationCode    string   `json:"urbanizationCode"`
	StateOrProvinceCode string   `json:"stateOrProvinceCode"`
	PostalCode          string   `json:"postalCode"`
	CountryCode         string   `json:"countryCode"`
	CountryName         string   `json:"countryName"`
}

type AncillaryDetail struct {
	Reason            string `json:"reason"`
	ReasonDescription string `json:"reasonDescription"`
	Action            string `json:"action"`
	ActionDescription string `json:"actionDescription"`
}

type DelayDetail struct {
	Type    string `json:"type"`
	SubType string `json:"subType"`
	Status  string `json:"status"`
}

type ServiceCommitMessage struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

type InformationNote struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type SpecialHandling struct {
	Description string `json:"description"`
	Type        string `json:"type"`
	PaymentType string `json:"paymentType"`
}

type AvailableImage struct {
	Size string `json:"size"`
	Type string `json:"type"`
}

type DeliveryDetails struct {
	ReceivedByName                    string                            `json:"receivedByName"`
	DestinationServiceArea            string                            `json:"destinationServiceArea"`
	DestinationServiceAreaDescription string                            `json:"destinationServiceAreaDescription"`
	LocationDescription               string                            `json:"locationDescription"`
	ActualDeliveryAddress             ActualDeliveryAddress             `json:"actualDeliveryAddress"`
	DeliveryToday                     bool                              `json:"deliveryToday"`
	LocationType                      string                            `json:"locationType"`
	SignedByName                      string                            `json:"signedByName"`
	OfficeOrderDeliveryMethod         string                            `json:"officeOrderDeliveryMethod"`
	DeliveryAttempts                  string                            `json:"deliveryAttempts"`
	DeliveryOptionEligibilityDetails  []DeliveryOptionEligibilityDetail `json:"deliveryOptionEligibilityDetails"`
}

type ActualDeliveryAddress struct {
	Classification      string   `json:"classification"`
	Residential         bool     `json:"residential"`
	StreetLines         []string `json:"streetLines"`
	City                string   `json:"city"`
	UrbanizationCode    string   `json:"urbanizationCode"`
	StateOrProvinceCode string   `json:"stateOrProvinceCode"`
	PostalCode          string   `json:"postalCode"`
	CountryCode         string   `json:"countryCode"`
	CountryName         string   `json:"countryName"`
}

type DeliveryOptionEligibilityDetail struct {
	Option      string `json:"option"`
	Eligibility string `json:"eligibility"`
}

type ScanEvent struct {
	Date                 time.Time         `json:"date"`
	DerivedStatus        string            `json:"derivedStatus"`
	ScanLocation         ScanEventLocation `json:"scanLocation"`
	ExceptionDescription string            `json:"exceptionDescription"`
	EventDescription     string            `json:"eventDescription"`
	EventType            string            `json:"eventType"`
	DerivedStatusCode    string            `json:"derivedStatusCode"`
	ExceptionCode        string            `json:"exceptionCode"`
	DelayDetail          DelayDetail       `json:"delayDetail"`
}

type ScanEventLocation struct {
	LocationId                string            `json:"locationId"`
	LocationContactAndAddress ContactAndAddress `json:"locationContactAndAddress"`
	LocationType              string            `json:"locationType"`
}

type DateAndTime struct {
	DateTime string `json:"dateTime"`
	Type     string `json:"type"`
}

type PackageDetails struct {
	PhysicalPackagingType string               `json:"physicalPackagingType"`
	SequenceNumber        string               `json:"sequenceNumber"`
	UndeliveredCount      string               `json:"undeliveredCount"`
	PackagingDescription  PackagingDescription `json:"packagingDescription"`
	Count                 string               `json:"count"`
	WeightAndDimensions   WeightAndDimensions  `json:"weightAndDimensions"`
	PackageContent        []string             `json:"packageContent"`
	ContentPieceCount     string               `json:"contentPieceCount"`
	DeclaredValue         DeclaredValue        `json:"declaredValue"`
}

type PackagingDescription struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

type WeightAndDimensions struct {
	Weight     []Weight    `json:"weight"`
	Dimensions []Dimension `json:"dimensions"`
}

type Weight struct {
	Unit  string `json:"unit"`
	Value string `json:"value"`
}

type Dimension struct {
	Length int    `json:"length"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Units  string `json:"units"`
}

type DeclaredValue struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}

type HoldAtLocation struct {
	LocationId                string            `json:"locationId"`
	LocationContactAndAddress ContactAndAddress `json:"locationContactAndAddress"`
	LocationType              string            `json:"locationType"`
}

type CustomDeliveryOption struct {
	RequestedAppointmentDetail RequestedAppointmentDetail `json:"requestedAppointmentDetail"`
	Description                string                     `json:"description"`
	Type                       string                     `json:"type"`
	Status                     string                     `json:"status"`
}

type RequestedAppointmentDetail struct {
	Date   string              `json:"date"`
	Window []DescriptionWindow `json:"window"`
}

type DescriptionWindow struct {
	Description string `json:"description"`
	Window      Window `json:"window"`
	Type        string `json:"type"`
}

type Window struct {
	Begins string    `json:"begins"`
	Ends   time.Time `json:"ends"`
}

type PieceCount struct {
	Count       string `json:"count"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type OriginLocation struct {
	LocationId                string            `json:"locationId"`
	LocationContactAndAddress ContactAndAddress `json:"locationContactAndAddress"`
	LocationType              string            `json:"locationType"`
}

type ContactAndAddress struct {
	Contact Contact `json:"contact"`
	Address Address `json:"address"`
}

type ShipmentDetails struct {
	Contents               []Content       `json:"contents"`
	BeforePossessionStatus bool            `json:"beforePossessionStatus"`
	Weight                 []Weight        `json:"weight"`
	ContentPieceCount      string          `json:"contentPieceCount"`
	SplitShipments         []SplitShipment `json:"splitShipments"`
}

type Content struct {
	ItemNumber       string `json:"itemNumber"`
	ReceivedQuantity string `json:"receivedQuantity"`
	Description      string `json:"description"`
	PartNumber       string `json:"partNumber"`
}

type SplitShipment struct {
	PieceCount        string `json:"pieceCount"`
	StatusDescription string `json:"statusDescription"`
	Timestamp         string `json:"timestamp"`
	StatusCode        string `json:"statusCode"`
}
