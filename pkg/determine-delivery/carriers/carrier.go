package carriers

import "time"

type Tracker interface {
	Track(string) ([]Parcel, error)
}

type Detector interface {
	Detect(string) bool
}

type Carrier interface {
	Tracker
	Detector
}

type Parcel struct {
	TrackingNumber string    `json:"TrackingNumber,omitempty"`
	Places         []Place   `json:"Places,omitempty"`
	Status         string    `json:"Status,omitempty"`
	DeliveryDate   time.Time `json:"DeliveryDate,omitempty"`
}
type Place struct {
	Country string    `json:"Country,omitempty"`
	City    string    `json:"City,omitempty"`
	Street  string    `json:"Street,omitempty"`
	Address string    `json:"Address,omitempty"`
	Comment string    `json:"Comment,omitempty"`
	Date    time.Time `json:"Date,omitempty"`
}
