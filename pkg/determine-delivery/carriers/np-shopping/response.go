package np_shopping

type TrackingDocumentResponse struct {
	WaybillNumber string `json:"waybill_number"`
	IewNum        string `json:"iew_num"`
	State         string `json:"state"`
	PickupAddress struct {
		Country string `json:"country"`
	} `json:"pickup_address"`
	DeliveryAddress struct {
		Country string `json:"country"`
	} `json:"delivery_address"`
	PickUpDate      string `json:"pick_up_date"`
	GeneralCostInfo struct {
		Cost     string `json:"cost"`
		Currency string `json:"currency"`
	} `json:"general_cost_info"`
	ShipmentInfo struct {
		SumActualWeight        string `json:"sum_actual_weight"`
		SumVolumeWeight        string `json:"sum_volume_weight"`
		SumControlActualWeight string `json:"sum_control_actual_weight"`
		SumControlVolumeWeight string `json:"sum_control_volume_weight"`
	} `json:"shipment_info"`
	Service struct {
		Name         string `json:"name"`
		PayerType    string `json:"payer_type"`
		PaymentType  string `json:"payment_type"`
		DeliveryType string `json:"delivery_type"`
	} `json:"service"`
	References []struct {
		Num  string `json:"num"`
		Type string `json:"type"`
	} `json:"references"`
	TrackingHistory []struct {
		Code         string `json:"code"`
		Description  string `json:"description"`
		Date         string `json:"date"`
		Country      string `json:"country"`
		Type         string `json:"type"`
		Translations []struct {
			Lang        string `json:"lang"`
			Description string `json:"description"`
		} `json:"translations"`
	} `json:"tracking_history"`
	Result        string `json:"result"`
	SearchHystory string `json:"search_hystory"`
}
