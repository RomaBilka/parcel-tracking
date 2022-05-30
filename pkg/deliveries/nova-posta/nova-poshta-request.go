package nova_posta

type novaPoshtaRequest struct {
	ApiKey           string      `json:"apiKey,omitempty"`
	ModelName        string      `json:"modelName,omitempty"`
	CalledMethod     string      `json:"calledMethod,omitempty"`
	MethodProperties interface{} `json:"methodProperties,omitempty"`
}

type trackingDocuments struct {
	Documents         []trackingDocument `json:"Documents,omitempty"`
	CheckWeightMethod string             `json:"CheckWeightMethod"`
	CalculatedWeight  string             `json:"CalculatedWeight"`
}

type trackingDocument struct {
	DocumentNumber string `json:"DocumentNumber,omitempty"`
	Phone          string `json:"Phone,omitempty"`
}
