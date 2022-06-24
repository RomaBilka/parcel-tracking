package carriers

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
	Number  string
	Address string
	Status  string
}
