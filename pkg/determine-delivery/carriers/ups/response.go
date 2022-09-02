package ups

type TrackResponse struct {
	Response struct {
		ResponseStatusCode        string `xml:"ResponseStatusCode"`
		ResponseStatusDescription string `xml:"ResponseStatusDescription"`
		TransactionReference      struct {
			CustomerContext string `xml:"CustomerContext"`
			XpciVersion     string `xml:"XpciVersion"`
		} `xml:"TransactionReference"`
		Error struct {
			ErrorSeverity       string `xml:"ErrorSeverity"`
			ErrorCode           string `xml:"ErrorCode"`
			ErrorDescription    string `xml:"ErrorDescription"`
			MinimumRetrySeconds string `xml:"MinimumRetrySeconds"`
			ErrorDigest         string `xml:"ErrorDigest"`
			ErrorLocation       struct {
				ErrorLocationElementName   string `xml:"ErrorLocationElementName"`
				ErrorLocationAttributeName string `xml:"ErrorLocationAttributeName"`
			} `xml:"ErrorLocation"`
		} `xml:"Error"`
	} `xml:"Response"`
	Shipment Shipment `xml:"Shipment"`
}

type Shipment struct {
	ShipmentIdentificationNumber string `xml:"ShipmentIdentificationNumber"`
	CandidateBookmark            string `xml:"CandidateBookmark"`
	PickupDate                   string `xml:"PickupDate"`
	BillToName                   string `xml:"BillToName"`
	NumberOfPieces               string `xml:"NumberOfPieces"`
	NumberOfPallets              string `xml:"NumberOfPallets"`
	SignedForByName              string `xml:"SignedForByName"`
	DescriptionOfGoods           string `xml:"DescriptionOfGoods"`
	GraphicImage                 string `xml:"GraphicImage"`
	ScheduledDeliveryDate        string `xml:"ScheduledDeliveryDate"`
	ScheduledDeliveryTime        string `xml:"ScheduledDeliveryTime"`
	FileNumber                   string `xml:"FileNumber"`

	ReferenceNumber struct {
		Code        string `xml:"Code"`
		Description string `xml:"Description"`
		Value       string `xml:"Value"`
	} `xml:"ReferenceNumber"`

	CurrentStatus struct {
		Code        string `xml:"Code"`
		Description string `xml:"Description"`
	} `xml:"CurrentStatus"`

	InquiryNumber struct {
		Code        string `xml:"Code"`
		Description string `xml:"Description"`
		Value       string `xml:"Value"`
	} `xml:"InquiryNumber"`

	ShipmentType struct {
		Code        string `xml:"Code"`
		Description string `xml:"Description"`
	} `xml:"ShipmentType"`

	Shipper Shipper `xml:"Shipper"`

	ShipTo ShipTo `xml:"ShipTo"`

	ShipmentWeight struct {
		Weight            string `xml:"Weight"`
		UnitOfMeasurement struct {
			Code        string `xml:"Code"`
			Description string `xml:"Description"`
		} `xml:"UnitOfMeasurement"`
	} `xml:"ShipmentWeight"`

	DeliveryDetails struct {
		DeliveryDate struct {
			Date string `xml:"Date"`
			Time string `xml:"Time"`
		} `xml:"DeliveryDate"`

		ServiceCenter struct {
			City              string `xml:"City"`
			StateProvinceCode string `xml:"StateProvinceCode"`
		} `xml:"ServiceCenter"`
	} `xml:"DeliveryDetails"`

	DeliveryDateTime struct {
		Date string `xml:"Date"`
		Type struct {
			Code        string `xml:"Code"`
			Description string `xml:"Description"`
		} `xml:"Type"`
	} `xml:"DeliveryDateTime"`

	Volume struct {
		Code        string `xml:"Code"`
		Description string `xml:"Description"`
		Value       string `xml:"Value"`
	} `xml:"Volume"`

	PickUpServiceCenter struct {
		City              string `xml:"City"`
		StateProvinceCode string `xml:"StateProvinceCode"`
	} `xml:"PickUpServiceCenter"`

	ShipmentServiceOptions struct {
		COD struct {
			CODAmount struct {
				CurrencyCode  string `xml:"CurrencyCode"`
				MonetaryValue string `xml:"MonetaryValue"`
			} `xml:"CODAmount"`
		} `xml:"COD"`
	} `xml:"ShipmentServiceOptions"`

	EstimatedDeliveryDetails struct {
		Date          string `xml:"Date"`
		Time          string `xml:"Time"`
		ServiceCenter struct {
			City              string `xml:"City"`
			StateProvinceCode string `xml:"StateProvinceCode"`
		} `xml:"ServiceCenter"`
	} `xml:"EstimatedDeliveryDetails"`

	Service struct {
		Code        string `xml:"Code"`
		Description string `xml:"Description"`
	} `xml:"Service"`

	Activity struct {
		Description      string `xml:"Description"`
		Date             string `xml:"Date"`
		Time             string `xml:"Time"`
		Trailer          string `xml:"Trailer"`
		ActivityLocation struct {
			Address struct {
				City              string `xml:"City"`
				PostalCode        string `xml:"PostalCode"`
				StateProvinceCode string `xml:"StateProvinceCode"`
				CountryCode       string `xml:"CountryCode"`
			} `xml:"Address"`
		} `xml:"ActivityLocation"`
	} `xml:"Activity"`

	OriginPortDetails struct {
		OriginPort         string `xml:"OriginPort"`
		EstimatedDeparture struct {
			Date string `xml:"Date"`
			Time string `xml:"Time"`
		} `xml:"EstimatedDeparture"`
	} `xml:"OriginPortDetails"`

	DestinationPortDetails struct {
		DestinationPort  string `xml:"DestinationPort"`
		EstimatedArrival struct {
			Date string `xml:"Date"`
			Time string `xml:"Time"`
		} `xml:"EstimatedArrival"`
	} `xml:"DestinationPortDetails"`

	CargoReady struct {
		Date string `xml:"Date"`
		Time string `xml:"Time"`
	} `xml:"CargoReady"`

	Manifest struct {
		Date string `xml:"Date"`
		Time string `xml:"Time"`
	} `xml:"Manifest"`

	CarrierActivityInformation struct {
		CarrierId       string `xml:"CarrierId"`
		Description     string `xml:"Description"`
		Status          string `xml:"Status"`
		OriginPort      string `xml:"OriginPort"`
		DestinationPort string `xml:"DestinationPort"`
		Arrival         struct {
			Date string `xml:"Date"`
			Time string `xml:"Time"`
		} `xml:"Arrival"`
		Departure struct {
			Date string `xml:"Date"`
			Time string `xml:"Time"`
		} `xml:"Departure"`
	} `xml:"CarrierActivityInformation"`

	Document struct {
		Type struct {
			Code        string `xml:"Code"`
			Description string `xml:"Description"`
		} `xml:"Type"`
		Format struct {
			Code        string `xml:"Code"`
			Description string `xml:"Description"`
		} `xml:"Format"`
	} `xml:"Document"`

	DeliveryDateUnavailable struct {
		Type        string `xml:"Type"`
		Description string `xml:"Description"`
	} `xml:"DeliveryDateUnavailable"`

	Appointment struct {
		BeginTime string `xml:"BeginTime"`
		EndTime   string `xml:"EndTime"`
		Made      struct {
			Date string `xml:"Date"`
			Time string `xml:"Time"`
		} `xml:"Made"`
		Requested struct {
			Date string `xml:"Date"`
			Time string `xml:"Time"`
		} `xml:"Requested"`
	} `xml:"Appointment"`

	Package Package `xml:"Package"`
}

type ShipTo struct {
	Address Address `xml:"Address"`
}

type Shipper struct {
	ShipperNumber string  `xml:"ShipperNumber"`
	Address       Address `xml:"Address"`
}

type Package struct {
	TrackingNumber          string `xml:"TrackingNumber"`
	DeliveryIndicator       string `xml:"DeliveryIndicator"`
	DeliveryDate            string `xml:"DeliveryDate"`
	RescheduledDeliveryDate string `xml:"RescheduledDeliveryDate"`
	RescheduledDeliveryTime string `xml:"RescheduledDeliveryTime"`
	SRSizeCode              string `xml:"SRSizeCode"`
	Redirect                struct {
		CompanyName  string  `xml:"CompanyName"`
		LocationID   string  `xml:"LocationID"`
		PickupDate   string  `xml:"PickupDate"`
		UPSAPAddress Address `xml:"UPSAPAddress"`
	} `xml:"Redirect"`
	Reroute struct {
		Address Address `xml:"Address"`
	} `xml:"Reroute"`

	ReturnTo struct {
		Address Address `xml:"Address"`
	} `xml:"ReturnTo"`
	PackageServiceOptions struct {
		ImportControl            string `xml:"ImportControl"`
		CommercialInvoiceRemoval string `xml:"CommercialInvoiceRemoval"`
		UPScarbonneutral         string `xml:"UPScarbonneutral"`
		USPSPICNumber            string `xml:"USPSPICNumber"`
		ExchangeBased            string `xml:"ExchangeBased"`
		PackAndCollect           string `xml:"PackAndCollect"`
		COD                      struct {
			ControlNumber string `xml:"ControlNumber"`
			CODStatus     string `xml:"CODStatus"`
			CODAmount     struct {
				CurrencyCode  string `xml:"CurrencyCode"`
				MonetaryValue string `xml:"MonetaryValue"`
			} `xml:"CODAmount"`
		} `xml:"COD"`
		SignatureRequired struct {
			Code string `xml:"Code"`
		} `xml:"SignatureRequired"`
	} `xml:"PackageServiceOptions"`

	Activity []Activity `xml:"Activity"`

	Message struct {
		Code        string   `xml:"Code"`
		Description struct{} `xml:"Description"`
	} `xml:"Message"`

	ReferenceNumber struct {
		Code  string   `xml:"Code"`
		Value struct{} `xml:"Value"`
	} `xml:"ReferenceNumber"`

	ProductType struct {
		Code        string   `xml:"Code"`
		Description struct{} `xml:"Description"`
	} `xml:"ProductType"`

	Accessorial struct {
		Code        string   `xml:"Code"`
		Description struct{} `xml:"Description"`
	} `xml:"Accessorial"`

	AlternateTrackingInfo struct {
		Type  string   `xml:"Type"`
		Value struct{} `xml:"Value"`
	} `xml:"AlternateTrackingInfo"`

	PackageWeight struct {
		Weight            string `xml:"Weight"`
		UnitOfMeasurement struct {
			Code        string `xml:"Code"`
			Description string `xml:"Description"`
		} `xml:"UnitOfMeasurement"`
	} `xml:"PackageWeight"`

	LocationAssured                string `xml:"LocationAssured"`
	AlternateTrackingNumber        string `xml:"AlternateTrackingNumber"`
	DimensionalWeightScanIndicator string `xml:"DimensionalWeightScanIndicator"`

	PreauthorizedReturnInformation struct {
		ReturnEligibilityIndicator string `xml:"ReturnEligibilityIndicator"`
		ReturnExpirationDate       string `xml:"ReturnExpirationDate"`
		ReturnRequestURL           string `xml:"ReturnRequestURL"`
		OriginalTrackingNumber     string `xml:"OriginalTrackingNumber"`
		ReturnTrackingNumber       string `xml:"ReturnTrackingNumber"`
	} `xml:"PreauthorizedReturnInformation"`

	UPSPremierAccessorial struct {
		UPSPremierCode        string `xml:"UPSPremierCode"`
		UPSPremierDescription string `xml:"UPSPremierDescription"`
	} `xml:"UPSPremierAccessorial"`
}

type Address struct {
	AddressLine1      string `xml:"AddressLine1"`
	AddressLine2      string `xml:"AddressLine2"`
	AddressLine3      string `xml:"AddressLine3"`
	City              string `xml:"City"`
	StateProvinceCode string `xml:"StateProvinceCode"`
	PostalCode        string `xml:"PostalCode"`
	CountryCode       string `xml:"CountryCode"`
}
type Activity struct {
	Date                              string `xml:"Date"`
	Time                              string `xml:"Time"`
	GMTDate                           string `xml:"GMTDate"`
	GMTTime                           string `xml:"GMTTime"`
	GMTOffset                         string `xml:"GMTOffset"`
	DeliveryDateFromManifestIndicator string `xml:"DeliveryDateFromManifestIndicator"`
	SensorEventIndicator              string `xml:"SensorEventIndicator"`

	NextScheduleActivity struct {
		Date string `xml:"Date"`
		Time string `xml:"Time"`
	} `xml:"NextScheduleActivity"`

	AlternateTrackingInfo struct {
		Type  string `xml:"Type"`
		Value string `xml:"Value"`
	} `xml:"AlternateTrackingInfo"`

	ActivityLocation struct {
		Address Address `xml:"Address"`

		AddressArtifactFormat struct {
			StreetNumberLow    string `xml:"StreetNumberLow"`
			StreetPrefix       string `xml:"StreetPrefix"`
			StreetName         string `xml:"StreetName"`
			StreetSuffix       string `xml:"StreetSuffix"`
			StreetType         string `xml:"StreetType"`
			PoliticalDivision2 string `xml:"PoliticalDivision2"`
			PoliticalDivision1 string `xml:"PoliticalDivision1"`
			PostcodePrimaryLow string `xml:"PostcodePrimaryLow"`
			CountryCode        string `xml:"CountryCode"`
		} `xml:"AddressArtifactFormat"`

		TransportFacility struct {
			Type  string `xml:"Type"`
			Value string `xml:"Value"`
		} `xml:"TransportFacility"`

		Code            string `xml:"Code"`
		Description     string `xml:"Description"`
		SignedForByName string `xml:"SignedForByName"`

		SignatureImage struct {
			GraphicImage string `xml:"GraphicImage"`
			ImageFormat  struct {
				Code        string `xml:"Code"`
				Description string `xml:"Description"`
			} `xml:"ImageFormat"`
		} `xml:"SignatureImage"`

		PODLetter struct {
			HTMLImage string `xml:"HTMLImage"`
		} `xml:"PODLetter"`

		ElectronicDeliveryNotification struct {
			Name string `xml:"Name"`
		} `xml:"ElectronicDeliveryNotification"`
	} `xml:"ActivityLocation"`

	Status struct {
		Date       string `xml:"Date"`
		Time       string `xml:"Time"`
		StatusType struct {
			Code        string `xml:"Code"`
			Description string `xml:"Description"`
		} `xml:"StatusType"`
		StatusCode struct {
			Code string `xml:"Code"`
		} `xml:"StatusCode"`
	} `xml:"Status"`
}
