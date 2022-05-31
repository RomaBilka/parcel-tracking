package nova_poshta

type novaPoshtaRequest struct {
	ApiKey           string      `json:"apiKey,omitempty"`
	ModelName        string      `json:"modelName,omitempty"`
	CalledMethod     string      `json:"calledMethod,omitempty"`
	MethodProperties interface{} `json:"methodProperties,omitempty"`
}

type TrackingDocuments struct {
	Documents         []TrackingDocument `json:"Documents,omitempty"`
	CheckWeightMethod string             `json:"CheckWeightMethod"`
	CalculatedWeight  string             `json:"CalculatedWeight"`
}

type TrackingDocument struct {
	DocumentNumber string `json:"DocumentNumber,omitempty"`
	Phone          string `json:"Phone,omitempty"`
}
