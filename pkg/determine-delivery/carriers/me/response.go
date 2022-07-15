package me

type meestExpressResponse struct {
	Api        string `xml:"api"`
	ApiVersion string `xml:"apiversion"`
	ResultUuid string `xml:"result_uuid"`
	Errors     Errors `xml:"errors"`
}

type ShipmentsTrackResponse struct {
	meestExpressResponse
	ResultTable []ShipmentTrackResponse `xml:"result_table>items"`
}

type ShipmentTrackResponse struct {
	DocumentIdRef         string `xml:"DocumentIdRef"`        // 0xa969003048d2b47311e3fc6109445fb4
	AgentsIdRef           string `xml:"AgentsIdRef"`          // 0xb95653547d06ad014b2580d2c4be55b2
	ShipmentNumberSender  string `xml:"ShipmentNumberSender"` // 001-0255371
	ShipmentNumberTransit string `xml:"ShipmentNumberTransit"`
	ShipmentClientID      string `xml:"ShipmentClientID"`

	Country            string `xml:"Country"`            // УКРАЇНА
	City               string `xml:"City"`               // Дніпродзержинськ
	ActionId           int    `xml:"ActionId"`           // 6
	StatusCode         string `xml:"StatusCode"`         // 606
	ActionMessages_UA  string `xml:"ActionMessages_UA"`  // Поступлення відправлення на відділення
	ActionMessages_RU  string `xml:"ActionMessages_RU"`  // Поступление отправления на отделение
	ActionMessages_EN  string `xml:"ActionMessages_EN"`  // Accepted by the branch
	DetailMessages_UA  string `xml:"DetailMessages_UA"`  // Дніпродзержинськ
	DetailMessages_RU  string `xml:"DetailMessages_RU"`  // Дніпродзержинськ
	DetailMessages_EN  string `xml:"DetailMessages_EN"`  // Дніпродзержинськ
	DetailPlacesAction string `xml:"DetailPlacesAction"` // 1/1
	CountryDel         string `xml:"CountryDel"`         // UA
	Recipient_Country  string `xml:"Recipient_Country"`
	Deliverydate       string `xml:"Deliverydate"`
}

type Errors struct {
	Code string `xml:"code"`
	Name string `xml:"name"`
}
