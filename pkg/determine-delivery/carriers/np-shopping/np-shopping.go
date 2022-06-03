package np_shopping

import "regexp"

var npShopping *regexp.Regexp

func init() {
	//NP99999999999999NPG
	npShopping = regexp.MustCompile(`(?i)^NP[\d]{14}NPG$`)
}

type NPShopping struct {
	TrackId string
}

func (n *NPShopping) Detect() bool {
	matched := npShopping.MatchString(n.TrackId)
	if matched {
		return true
	}

	return false
}
