package dhl

type response struct {
	Url                            string      `json:"url"`
	PrevUrl                        string      `json:"prevUrl"`
	NextUrl                        string      `json:"nextUrl"`
	FirstUrl                       string      `json:"firstUrl"`
	LastUrl                        string      `json:"lastUrl"`
	PossibleAdditionalShipmentsUrl []string    `json:"possibleAdditionalShipmentsUrl"`
	Shipments                      []shipments `json:"shipments"`
}

type shipments struct {
	Id                            string `json:"id"`
	Service                       string `json:"service"`
	EstimatedTimeOfDelivery       string `json:"estimatedTimeOfDelivery"`
	EstimatedTimeOfDeliveryRemark string `json:"estimatedTimeOfDeliveryRemark"`
	ServiceUrl                    string `json:"serviceUrl"`
	RerouteUrl                    string `json:"rerouteUrl"`

	Origin struct {
		Address address `json:"address"`
	} `json:"origin"`

	Destination struct {
		Address address `json:"address"`
	} `json:"destination"`

	Status struct {
		Timestamp string `json:"timestamp"`
	} `json:"status"`

	EstimatedDeliveryTimeFrame struct {
		EstimatedFrom    string `json:"estimatedFrom"`
		EstimatedThrough string `json:"estimatedThrough"`
	} `json:"estimatedDeliveryTimeFrame"`

	Details struct {
		TotalNumberOfPieces int      `json:"totalNumberOfPieces"`
		PieceIds            []string `json:"pieceIds"`
		Weight              struct {
			Value    float64 `json:"value"`
			UnitText string  `json:"unitText"`
		} `json:"weight"`
		Volume        map[string]string `json:"volume"`
		LoadingMeters string            `json:"loadingMeters"`

		Dimensions struct {
			Width  map[string]string `json:"width"`
			Height map[string]string `json:"height"`
			Length map[string]string `json:"length"`
		} `json:"dimensions"`

		References struct {
			Number string `json:"number"`
			Type   string `json:"type"`
		} `json:"references"`

		Routes struct {
			VesselName             string `json:"dgf:vesselName"`
			VoyageFlightNumber     string `json:"dgf:voyageFlightNumber"`
			EstimatedDepartureDate string `json:"dgf:estimatedDepartureDate"`
			EstimatedArrivalDate   string `json:"dgf:estimatedArrivalDate"`
		} `json:"dgf:routes"`

		AirportOfDeparture airport `json:"dgf:airportOfDeparture"`

		AirportOfDestination airport `json:"dgf:airportOfDestination"`

		PlaceOfAcceptance struct {
			LocationName string `json:"dgf:locationName"`
		} `json:"dgf:placeOfAcceptance"`

		PortOfLoading struct {
			LocationName string `json:"dgf:locationName"`
		} `json:"dgf:portOfLoading"`

		PortOfUnloading struct {
			LocationName string `json:"dgf:locationName"`
		} `json:"dgf:portOfUnloading"`

		PlaceOfDelivery struct {
			LocationName string `json:"dgf:locationName"`
		} `json:"dgf:placeOfDelivery"`
	} `json:"details"`

	Events struct {
		Timestamp   string   `json:"timestamp"`
		StatusCode  string   `json:"statusCode"`
		Status      string   `json:"status"`
		Description string   `json:"description"`
		PieceIds    []string `json:"pieceIds"`
		Remark      string   `json:"remark"`
		NextSteps   string   `json:"nextSteps"`

		Location struct {
			Address struct {
				CountryCode     string `json:"countryCode"`
				PostalCode      string `json:"postalCode"`
				AddressLocality string `json:"addressLocality"`
				StreetAddress   string `json:"streetAddress"`
			} `json:"address"`
		} `json:"location"`
	} `json:"events"`
}

type airport struct {
	LocationName string `json:"dgf:locationName"`
	LocationCode string `json:"dgf:locationCode"`
	CountryCode  string `json:"dgf:countryCode"`
}

type address struct {
	CountryCode     string `json:"countryCode"`
	PostalCode      string `json:"postalCode"`
	AddressLocality string `json:"addressLocality"`
	StreetAddress   string `json:"streetAddress"`
}
