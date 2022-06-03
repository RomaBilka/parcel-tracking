package carriers

import "regexp"

var npShopping *regexp.Regexp

func init() {
	//NP99999999999999NPG
	npShopping = regexp.MustCompile(`(?i)^NP[\d]{14}NPG$`)
}

func IsNpShopping(str string) bool {
	matched:=npShopping.MatchString(str)
	if matched {
		return true
	}

	return false
}
