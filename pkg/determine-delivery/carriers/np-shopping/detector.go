package np_shopping

import (
	"regexp"
)

var npShopping *regexp.Regexp

func init() {
	//NP99999999999999NPG
	npShopping = regexp.MustCompile(`(?i)^NP[\d]{14}NPG$`)
}

func (c *Carrier) Detect(trackId string) bool {
	matched := npShopping.MatchString(trackId)

	return matched
}
