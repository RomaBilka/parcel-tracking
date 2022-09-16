package ukrposhta

type TrackingDocuments struct {
	Documents         []TrackingDocument `json:"Documents,omitempty"`
	CheckWeightMethod string             `json:"CheckWeightMethod"`
	CalculatedWeight  string             `json:"CalculatedWeight"`
}

type TrackingDocument struct {
	DocumentNumber string `json:"DocumentNumber,omitempty"`
	Phone          string `json:"Phone,omitempty"`
}
