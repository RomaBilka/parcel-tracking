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

/*
type Parcel struct {
	Number  string
	Address string
	Status  string
}
*/
type Parcel struct {
	TrackingNumber string
	Places         []Place
	Status         string
	DeliveryDate   time.Time
}
type Place struct {
	County  string
	City    string
	Street  string
	Address string
	Comment string
	Date    time.Time
}
