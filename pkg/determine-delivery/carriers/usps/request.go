package usps

type TrackFieldRequest struct {
	UserId   string    `xml:"USERID,attr"`
	Revision int       `xml:"Revision"`
	ClientIp string    `xml:"ClientIp,omitempty"`
	SourceId string    `xml:"SourceId"`
	TrackID  []TrackID `xml:"TrackID"`
}

type TrackID struct {
	Id                 string `xml:"ID,attr"`
	DestinationZipCode string `xml:"DestinationZipCode,omitempty"`
	MailingDate        string `xml:"MailingDate,omitempty"`
}
