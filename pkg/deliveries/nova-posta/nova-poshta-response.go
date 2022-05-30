package nova_posta

type novaPoshtaResponse struct {
	Success      bool                `json:"success"`
	Data         []interface{}       `json:"data"`
	Errors       []string            `json:"errors"`
	Warnings     []map[string]string `json:"warnings"`
	Info         []map[string]string `json:"info"`
	MessageCodes []string            `json:"messageCodes"`
	ErrorCodes   []string            `json:"errorCodes"`
	WarningCodes []map[string]string `json:"warningCodes"`
	InfoCodes    []map[string]string `json:"infoCodes"`
}
