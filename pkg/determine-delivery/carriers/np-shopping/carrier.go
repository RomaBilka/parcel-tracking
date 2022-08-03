package np_shopping

import (
	"regexp"
)

////Starts with NP, 14 numbers and NPG at the end NP99999999999999NPG
var npShopping = regexp.MustCompile(`(?i)^NP[\d]{14}NPG$`)

type Carrier struct{}

func NewCarrier() *Carrier {
	return &Carrier{}
}

func (c *Carrier) Detect(trackId string) bool {
	return npShopping.MatchString(trackId)
}
