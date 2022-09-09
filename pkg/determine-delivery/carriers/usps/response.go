package usps

type TrackResponse struct {
	TrackInfo []TrackInfo `xml:"TrackInfo"`
}

type TrackInfo struct {
	TrackId           string       `xml:"ID,attr"`
	CarrierRelease    string       `xml:"CarrierRelease"`
	Class             string       `xml:"Class"`
	ClassOfMailCode   string       `xml:"ClassOfMailCode"`
	DestinationCity   string       `xml:"DestinationCity"`
	DestinationState  string       `xml:"DestinationState"`
	DestinationZip    string       `xml:"DestinationZip"`
	EmailEnabled      string       `xml:"EmailEnabled"`
	KahalaIndicator   string       `xml:"KahalaIndicator"`
	MailTypeCode      string       `xml:"MailTypeCode"`
	MPDATE            string       `xml:"MPDATE"`
	MPSUFFIX          string       `xml:"MPSUFFIX"`
	PodEnabled        string       `xml:"PodEnabled"`
	TPodEnabled       string       `xml:"TPodEnabled"`
	RedeliveryEnabled string       `xml:"RedeliveryEnabled"`
	RestoreEnabled    string       `xml:"RestoreEnabled"`
	RramEnabled       string       `xml:"RramEnabled"`
	RreEnabled        string       `xml:"RreEnabled"`
	Service           []Service    `xml:"Service"`
	ServiceTypeCode   string       `xml:"ServiceTypeCode"`
	Status            string       `xml:"Status"`
	StatusCategory    string       `xml:"StatusCategory"`
	StatusSummary     string       `xml:"StatusSummary"`
	TableCode         string       `xml:"TABLECODE"`
	TrackSummary      TrackSummary `xml:"TrackSummary"`
	Error             Error        `xml:"Error"`
}

type Service struct {
	Service string `xml:"Service"`
}

type TrackSummary struct {
	EventTime           string `xml:"EventTime"`
	EventDate           string `xml:"EventDate"`
	Event               string `xml:"Event"`
	EventCity           string `xml:"EventCity"`
	EventState          string `xml:"EventState"`
	EventZIPCode        string `xml:"EventZIPCode"`
	EventCountry        string `xml:"EventCountry"`
	FirmName            string `xml:"FirmName"`
	Name                string `xml:"Name"`
	AuthorizedAgent     string `xml:"AuthorizedAgent"`
	EventCode           string `xml:"EventCode"`
	EventStatusCategory string `xml:"EventStatusCategory"`
}

type Error struct {
	Number      string `xml:"Number"`
	Description string `xml:"Description"`
	HelpFile    string `xml:"HelpFile"`
	HelpContext string `xml:"HelpContext"`
}
