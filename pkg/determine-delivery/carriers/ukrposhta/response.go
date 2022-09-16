package ukrposhta

type Response struct {
	Found    map[string][]Item `json:"found"`
	NotFound []string          `json:"notFound"`
}

type Item struct {
	Barcode       string `json:"barcode"`
	Step          int    `json:"step"`
	Date          string `json:"date"`
	Index         string `json:"index"`
	Name          string `json:"name"`
	Event         string `json:"event"`
	EventName     string `json:"eventName"`
	Country       string `json:"country"`
	EventReason   string `json:"eventReason"`
	EventReasonId int    `json:"eventReason_id"`
	MailType      int    `json:"mailType"`
	IndexOrder    int    `json:"indexOrder"`
}
