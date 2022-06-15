package carriers

type Carrier interface {
	Tracking(string) ([]Parcel, error)
}

type Parcel struct {
	Number  string
	Address string
}
