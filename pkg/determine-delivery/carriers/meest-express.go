package carriers

import "regexp"

var startCV *regexp.Regexp
var startMYCV *regexp.Regexp

func init() {
	//CV999999999ZZ
	startCV = regexp.MustCompile(`(?i)^CV[\d]{9}[a-z][a-z]$`)

	//MYCV999999999ZZ
	startMYCV = regexp.MustCompile(`(?i)^MYCV[\d]{9}[a-z][a-z]$`)
}

func IsMeestExpress(str string) bool {

	matched := startCV.MatchString(str)
	if matched {
		return true
	}

	matched = startMYCV.MatchString(str)
	if matched {
		return true
	}

	return false
}
