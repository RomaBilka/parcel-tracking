package np

type NovaPoshtaResponse struct {
	Success      bool                `json:"success"`
	Errors       []string            `json:"errors,omitempty"`
	Warnings     any                 `json:"warnings,omitempty"`
	Info         []map[string]string `json:"info,omitempty"`
	MessageCodes []string            `json:"messageCodes,omitempty"`
	ErrorCodes   []string            `json:"errorCodes,omitempty"`
	WarningCodes []map[string]string `json:"warningCodes,omitempty"`
	InfoCodes    []map[string]string `json:"infoCodes,omitempty"`
}

type TrackingDocumentsResponse struct {
	NovaPoshtaResponse
	Data []TrackingDocumentResponse `json:"data"`
}

type TrackingDocumentResponse struct {
	Number                               string  `json:"Number"`                               //"20400048799000"
	Redelivery                           float64 `json:"Redelivery"`                           //"0"
	RedeliverySum                        any     `json:"RedeliverySum"`                        //"0"
	RedeliveryNum                        string  `json:"RedeliveryNum"`                        //""
	RedeliveryPayer                      string  `json:"RedeliveryPayer"`                      //"Sender/Recipient"
	OwnerDocumentType                    string  `json:"OwnerDocumentType"`                    //""
	LastCreatedOnTheBasisDocumentType    string  `json:"LastCreatedOnTheBasisDocumentType"`    //""
	LastCreatedOnTheBasisPayerType       string  `json:"LastCreatedOnTheBasisPayerType"`       //""
	LastCreatedOnTheBasisDateTime        string  `json:"LastCreatedOnTheBasisDateTime"`        //""
	LastTransactionStatusGM              string  `json:"LastTransactionStatusGM"`              //""
	LastTransactionDateTimeGM            string  `json:"LastTransactionDateTimeGM"`            //""
	LastAmountTransferGM                 string  `json:"LastAmountTransferGM"`                 //""
	DateCreated                          string  `json:"DateCreated"`                          //"18-11-2021 11:52:42"
	DocumentWeight                       float64 `json:"DocumentWeight"`                       //"3"
	FactualWeight                        string  `json:"FactualWeight"`                        //"3"
	VolumeWeight                         string  `json:"VolumeWeight"`                         //"0.1",
	CheckWeight                          float64 `json:"CheckWeight"`                          //""
	DocumentCost                         string  `json:"DocumentCost"`                         //"51"
	SumBeforeCheckWeight                 float64 `json:"SumBeforeCheckWeight"`                 //""
	PayerType                            string  `json:"PayerType"`                            //"Sender"
	RecipientFullName                    string  `json:"RecipientFullName"`                    //"??????"
	RecipientDateTime                    string  `json:"RecipientDateTime"`                    //"21.11.2021 13:53:47"
	ScheduledDeliveryDate                string  `json:"ScheduledDeliveryDate"`                //"19.11.2021 13:53:47"
	PaymentMethod                        string  `json:"PaymentMethod"`                        //"Cash"
	CargoDescriptionString               string  `json:"CargoDescriptionString"`               //"????????"
	CargoType                            string  `json:"CargoType"`                            //"Cargo",
	CitySender                           string  `json:"CitySender"`                           //"????????"
	CityRecipient                        string  `json:"CityRecipient"`                        //"????????"
	WarehouseRecipient                   string  `json:"WarehouseRecipient"`                   // "???????????????????? ???101 (???? 15 ????), ????????-????????????????????: ??????. ???????????? ??????????????????????????, 143/2, (??????. "????????")"
	CounterpartyType                     string  `json:"CounterpartyType"`                     //"PrivatePerson"
	AfterpaymentOnGoodsCost              any     `json:"AfterpaymentOnGoodsCost"`              //"0"
	ServiceType                          string  `json:"ServiceType"`                          //"WarehouseWarehouse"
	UndeliveryReasonsSubtypeDescription  string  `json:"UndeliveryReasonsSubtypeDescription"`  //""
	WarehouseRecipientNumber             float64 `json:"WarehouseRecipientNumber"`             //"101"
	LastCreatedOnTheBasisNumber          string  `json:"LastCreatedOnTheBasisNumber"`          //""
	PhoneRecipient                       string  `json:"PhoneRecipient"`                       //"380600000000"
	RecipientFullNameEW                  string  `json:"RecipientFullNameEW"`                  //""
	WarehouseRecipientInternetAddressRef string  `json:"WarehouseRecipientInternetAddressRef"` //"00000000-0000-0000-0000-000000000000"
	MarketplacePartnerToken              string  `json:"MarketplacePartnerToken"`              //""
	ClientBarcode                        string  `json:"ClientBarcode"`                        //""
	RecipientAddress                     string  `json:"RecipientAddress"`                     //"??. ????????, ???????????????????? ???101 (???? 15 ????), ????????-????????????????????, ??????. ???????????? ??????????????????????????, 143/2"
	CounterpartyRecipientDescription     string  `json:"CounterpartyRecipientDescription"`     //"???????????????? ??????????"
	CounterpartySenderType               string  `json:"CounterpartySenderType"`               //"PrivatePerson",
	DateScan                             string  `json:"DateScan"`                             //"0001-01-01 00:00:00"
	PaymentStatus                        string  `json:"PaymentStatus"`                        //""
	PaymentStatusDate                    string  `json:"PaymentStatusDate"`                    //""
	AmountToPay                          string  `json:"AmountToPay"`                          //""
	AmountPaid                           string  `json:"AmountPaid"`                           //""
	Status                               string  `json:"Status"`                               //""
	StatusCode                           string  `json:"StatusCode"`                           //""
	RefEW                                string  `json:"RefEW"`                                //"00000000-0000-0000-0000-000000000000"
	BackwardDeliverySubTypesActions      string  `json:"BackwardDeliverySubTypesActions"`      //""
	BackwardDeliverySubTypesServices     string  `json:"BackwardDeliverySubTypesServices"`     //""
	UndeliveryReasons                    string  `json:"UndeliveryReasons"`                    //""
	DatePayedKeeping                     string  `json:"DatePayedKeeping"`                     //""
}
