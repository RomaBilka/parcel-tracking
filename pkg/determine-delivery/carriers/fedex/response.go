package fedex

import "time"

type authResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

type Response struct {
	TransactionId         string `json:"transactionId"`
	CustomerTransactionId string `json:"customerTransactionId"`
	Output                struct {
		CompleteTrackResults []struct {
			TrackingNumber string `json:"trackingNumber"`
			TrackResults   []struct {
				TrackingNumberInfo struct {
					TrackingNumber         string `json:"trackingNumber"`
					CarrierCode            string `json:"carrierCode"`
					TrackingNumberUniqueId string `json:"trackingNumberUniqueId"`
				} `json:"trackingNumberInfo"`
				AdditionalTrackingInfo struct {
					HasAssociatedShipments bool   `json:"hasAssociatedShipments"`
					Nickname               string `json:"nickname"`
					PackageIdentifiers     []struct {
						Type                   string `json:"type"`
						Value                  string `json:"value"`
						TrackingNumberUniqueId string `json:"trackingNumberUniqueId"`
					} `json:"packageIdentifiers"`
					ShipmentNotes string `json:"shipmentNotes"`
				} `json:"additionalTrackingInfo"`
				DistanceToDestination struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"distanceToDestination"`
				ConsolidationDetail []struct {
					TimeStamp       time.Time `json:"timeStamp"`
					ConsolidationID string    `json:"consolidationID"`
					ReasonDetail    struct {
						Description string `json:"description"`
						Type        string `json:"type"`
					} `json:"reasonDetail"`
					PackageCount int    `json:"packageCount"`
					EventType    string `json:"eventType"`
				} `json:"consolidationDetail"`
				MeterNumber  string `json:"meterNumber"`
				ReturnDetail struct {
					AuthorizationName string `json:"authorizationName"`
					ReasonDetail      []struct {
						Description string `json:"description"`
						Type        string `json:"type"`
					} `json:"reasonDetail"`
				} `json:"returnDetail"`
				ServiceDetail struct {
					Description      string `json:"description"`
					ShortDescription string `json:"shortDescription"`
					Type             string `json:"type"`
				} `json:"serviceDetail"`
				DestinationLocation struct {
					LocationId                string `json:"locationId"`
					LocationContactAndAddress struct {
						Contact struct {
							PersonName  string `json:"personName"`
							PhoneNumber string `json:"phoneNumber"`
							CompanyName string `json:"companyName"`
						} `json:"contact"`
						Address struct {
							Classification      string   `json:"classification"`
							Residential         bool     `json:"residential"`
							StreetLines         []string `json:"streetLines"`
							City                string   `json:"city"`
							UrbanizationCode    string   `json:"urbanizationCode"`
							StateOrProvinceCode string   `json:"stateOrProvinceCode"`
							PostalCode          string   `json:"postalCode"`
							CountryCode         string   `json:"countryCode"`
						} `json:"address"`
					} `json:"locationContactAndAddress"`
					LocationType string `json:"locationType"`
				} `json:"destinationLocation"`
				LatestStatusDetail struct {
					ScanLocation struct {
						Classification      string   `json:"classification"`
						Residential         bool     `json:"residential"`
						StreetLines         []string `json:"streetLines"`
						City                string   `json:"city"`
						UrbanizationCode    string   `json:"urbanizationCode"`
						StateOrProvinceCode string   `json:"stateOrProvinceCode"`
						PostalCode          string   `json:"postalCode"`
						CountryCode         string   `json:"countryCode"`
						CountryName         string   `json:"countryName"`
					} `json:"scanLocation"`
					Code             string `json:"code"`
					DerivedCode      string `json:"derivedCode"`
					AncillaryDetails []struct {
						Reason            string `json:"reason"`
						ReasonDescription string `json:"reasonDescription"`
						Action            string `json:"action"`
						ActionDescription string `json:"actionDescription"`
					} `json:"ancillaryDetails"`
					StatusByLocale string `json:"statusByLocale"`
					Description    string `json:"description"`
					DelayDetail    struct {
						Type    string `json:"type"`
						SubType string `json:"subType"`
						Status  string `json:"status"`
					} `json:"delayDetail"`
				} `json:"latestStatusDetail"`
				ServiceCommitMessage struct {
					Message string `json:"message"`
					Type    string `json:"type"`
				} `json:"serviceCommitMessage"`
				InformationNotes []struct {
					Code        string `json:"code"`
					Description string `json:"description"`
				} `json:"informationNotes"`
				Error struct {
					Code          string `json:"code"`
					ParameterList []struct {
						Value string `json:"value"`
						Key   string `json:"key"`
					} `json:"parameterList"`
					Message string `json:"message"`
				} `json:"error"`
				SpecialHandlings []struct {
					Description string `json:"description"`
					Type        string `json:"type"`
					PaymentType string `json:"paymentType"`
				} `json:"specialHandlings"`
				AvailableImages []struct {
					Size string `json:"size"`
					Type string `json:"type"`
				} `json:"availableImages"`
				DeliveryDetails struct {
					ReceivedByName                    string `json:"receivedByName"`
					DestinationServiceArea            string `json:"destinationServiceArea"`
					DestinationServiceAreaDescription string `json:"destinationServiceAreaDescription"`
					LocationDescription               string `json:"locationDescription"`
					ActualDeliveryAddress             struct {
						Classification      string   `json:"classification"`
						Residential         bool     `json:"residential"`
						StreetLines         []string `json:"streetLines"`
						City                string   `json:"city"`
						UrbanizationCode    string   `json:"urbanizationCode"`
						StateOrProvinceCode string   `json:"stateOrProvinceCode"`
						PostalCode          string   `json:"postalCode"`
						CountryCode         string   `json:"countryCode"`
						CountryName         string   `json:"countryName"`
					} `json:"actualDeliveryAddress"`
					DeliveryToday                    bool   `json:"deliveryToday"`
					LocationType                     string `json:"locationType"`
					SignedByName                     string `json:"signedByName"`
					OfficeOrderDeliveryMethod        string `json:"officeOrderDeliveryMethod"`
					DeliveryAttempts                 string `json:"deliveryAttempts"`
					DeliveryOptionEligibilityDetails []struct {
						Option      string `json:"option"`
						Eligibility string `json:"eligibility"`
					} `json:"deliveryOptionEligibilityDetails"`
				} `json:"deliveryDetails"`
				ScanEvents []struct {
					Date          time.Time `json:"date"`
					DerivedStatus string    `json:"derivedStatus"`
					ScanLocation  struct {
						LocationId                string `json:"locationId"`
						LocationContactAndAddress struct {
							Contact struct {
								PersonName  string `json:"personName"`
								PhoneNumber string `json:"phoneNumber"`
								CompanyName string `json:"companyName"`
							} `json:"contact"`
							Address struct {
								Classification      string   `json:"classification"`
								Residential         bool     `json:"residential"`
								StreetLines         []string `json:"streetLines"`
								City                string   `json:"city"`
								UrbanizationCode    string   `json:"urbanizationCode"`
								StateOrProvinceCode string   `json:"stateOrProvinceCode"`
								PostalCode          string   `json:"postalCode"`
								CountryCode         string   `json:"countryCode"`
							} `json:"address"`
						} `json:"locationContactAndAddress"`
						LocationType string `json:"locationType"`
					} `json:"scanLocation"`
					ExceptionDescription string `json:"exceptionDescription"`
					EventDescription     string `json:"eventDescription"`
					EventType            string `json:"eventType"`
					DerivedStatusCode    string `json:"derivedStatusCode"`
					ExceptionCode        string `json:"exceptionCode"`
					DelayDetail          struct {
						Type    string `json:"type"`
						SubType string `json:"subType"`
						Status  string `json:"status"`
					} `json:"delayDetail"`
				} `json:"scanEvents"`
				DateAndTimes []struct {
					DateTime string `json:"dateTime"`
					Type     string `json:"type"`
				} `json:"dateAndTimes"`
				PackageDetails struct {
					PhysicalPackagingType string `json:"physicalPackagingType"`
					SequenceNumber        string `json:"sequenceNumber"`
					UndeliveredCount      string `json:"undeliveredCount"`
					PackagingDescription  struct {
						Description string `json:"description"`
						Type        string `json:"type"`
					} `json:"packagingDescription"`
					Count               string `json:"count"`
					WeightAndDimensions struct {
						Weight []struct {
							Unit  string `json:"unit"`
							Value string `json:"value"`
						} `json:"weight"`
						Dimensions []struct {
							Length int    `json:"length"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
							Units  string `json:"units"`
						} `json:"dimensions"`
					} `json:"weightAndDimensions"`
					PackageContent    []string `json:"packageContent"`
					ContentPieceCount string   `json:"contentPieceCount"`
					DeclaredValue     struct {
						Currency string  `json:"currency"`
						Value    float64 `json:"value"`
					} `json:"declaredValue"`
				} `json:"packageDetails"`
				GoodsClassificationCode string `json:"goodsClassificationCode"`
				HoldAtLocation          struct {
					LocationId                string `json:"locationId"`
					LocationContactAndAddress struct {
						Contact struct {
							PersonName  string `json:"personName"`
							PhoneNumber string `json:"phoneNumber"`
							CompanyName string `json:"companyName"`
						} `json:"contact"`
						Address struct {
							Classification      string   `json:"classification"`
							Residential         bool     `json:"residential"`
							StreetLines         []string `json:"streetLines"`
							City                string   `json:"city"`
							UrbanizationCode    string   `json:"urbanizationCode"`
							StateOrProvinceCode string   `json:"stateOrProvinceCode"`
							PostalCode          string   `json:"postalCode"`
							CountryCode         string   `json:"countryCode"`
							CountryName         string   `json:"countryName"`
						} `json:"address"`
					} `json:"locationContactAndAddress"`
					LocationType string `json:"locationType"`
				} `json:"holdAtLocation"`
				CustomDeliveryOptions []struct {
					RequestedAppointmentDetail struct {
						Date   string `json:"date"`
						Window []struct {
							Description string `json:"description"`
							Window      struct {
								Begins string    `json:"begins"`
								Ends   time.Time `json:"ends"`
							} `json:"window"`
							Type string `json:"type"`
						} `json:"window"`
					} `json:"requestedAppointmentDetail"`
					Description string `json:"description"`
					Type        string `json:"type"`
					Status      string `json:"status"`
				} `json:"customDeliveryOptions"`
				EstimatedDeliveryTimeWindow struct {
					Description string `json:"description"`
					Window      struct {
						Begins string    `json:"begins"`
						Ends   time.Time `json:"ends"`
					} `json:"window"`
					Type string `json:"type"`
				} `json:"estimatedDeliveryTimeWindow"`
				PieceCounts []struct {
					Count       string `json:"count"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"pieceCounts"`
				OriginLocation struct {
					LocationId                string `json:"locationId"`
					LocationContactAndAddress struct {
						Contact struct {
							PersonName  string `json:"personName"`
							PhoneNumber string `json:"phoneNumber"`
							CompanyName string `json:"companyName"`
						} `json:"contact"`
						Address struct {
							Classification      string   `json:"classification"`
							Residential         bool     `json:"residential"`
							StreetLines         []string `json:"streetLines"`
							City                string   `json:"city"`
							UrbanizationCode    string   `json:"urbanizationCode"`
							StateOrProvinceCode string   `json:"stateOrProvinceCode"`
							PostalCode          string   `json:"postalCode"`
							CountryCode         string   `json:"countryCode"`
						} `json:"address"`
					} `json:"locationContactAndAddress"`
					LocationType string `json:"locationType"`
				} `json:"originLocation"`
				RecipientInformation struct {
					Contact struct {
						PersonName  string `json:"personName"`
						PhoneNumber string `json:"phoneNumber"`
						CompanyName string `json:"companyName"`
					} `json:"contact"`
					Address struct {
						Classification      string   `json:"classification"`
						Residential         bool     `json:"residential"`
						StreetLines         []string `json:"streetLines"`
						City                string   `json:"city"`
						UrbanizationCode    string   `json:"urbanizationCode"`
						StateOrProvinceCode string   `json:"stateOrProvinceCode"`
						PostalCode          string   `json:"postalCode"`
						CountryCode         string   `json:"countryCode"`
						CountryName         string   `json:"countryName"`
					} `json:"address"`
				} `json:"recipientInformation"`
				StandardTransitTimeWindow struct {
					Description string `json:"description"`
					Window      struct {
						Begins string    `json:"begins"`
						Ends   time.Time `json:"ends"`
					} `json:"window"`
					Type string `json:"type"`
				} `json:"standardTransitTimeWindow"`
				ShipmentDetails struct {
					Contents []struct {
						ItemNumber       string `json:"itemNumber"`
						ReceivedQuantity string `json:"receivedQuantity"`
						Description      string `json:"description"`
						PartNumber       string `json:"partNumber"`
					} `json:"contents"`
					BeforePossessionStatus bool `json:"beforePossessionStatus"`
					Weight                 []struct {
						Unit  string `json:"unit"`
						Value string `json:"value"`
					} `json:"weight"`
					ContentPieceCount string `json:"contentPieceCount"`
					SplitShipments    []struct {
						PieceCount        string `json:"pieceCount"`
						StatusDescription string `json:"statusDescription"`
						Timestamp         string `json:"timestamp"`
						StatusCode        string `json:"statusCode"`
					} `json:"splitShipments"`
				} `json:"shipmentDetails"`
				ReasonDetail struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"reasonDetail"`
				AvailableNotifications []string `json:"availableNotifications"`
				ShipperInformation     struct {
					Contact struct {
						PersonName  string `json:"personName"`
						PhoneNumber string `json:"phoneNumber"`
						CompanyName string `json:"companyName"`
					} `json:"contact"`
					Address struct {
						Classification      string   `json:"classification"`
						Residential         bool     `json:"residential"`
						StreetLines         []string `json:"streetLines"`
						City                string   `json:"city"`
						UrbanizationCode    string   `json:"urbanizationCode"`
						StateOrProvinceCode string   `json:"stateOrProvinceCode"`
						PostalCode          string   `json:"postalCode"`
						CountryCode         string   `json:"countryCode"`
						CountryName         string   `json:"countryName"`
					} `json:"address"`
				} `json:"shipperInformation"`
				LastUpdatedDestinationAddress struct {
					Classification      string   `json:"classification"`
					Residential         bool     `json:"residential"`
					StreetLines         []string `json:"streetLines"`
					City                string   `json:"city"`
					UrbanizationCode    string   `json:"urbanizationCode"`
					StateOrProvinceCode string   `json:"stateOrProvinceCode"`
					PostalCode          string   `json:"postalCode"`
					CountryCode         string   `json:"countryCode"`
					CountryName         string   `json:"countryName"`
				} `json:"lastUpdatedDestinationAddress"`
			} `json:"trackResults"`
		} `json:"completeTrackResults"`
		Alerts string `json:"alerts"`
	} `json:"output"`
	Errors Errors `json:"errors"`
}

type Errors []struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
