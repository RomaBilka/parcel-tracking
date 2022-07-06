package np

import (
	"regexp"
)

var start59 *regexp.Regexp
var start20 *regexp.Regexp
var start1 *regexp.Regexp

func init() {
	//59************ length 14
	start59 = regexp.MustCompile(`^59[\d]{12}$`)

	//20************ length 14
	start20 = regexp.MustCompile(`^20[\d]{12}$`)

	//1************* length 14
	start1 = regexp.MustCompile(`^1[\d]{13}$`)
}

func (c *Carrier) Detect(trackId string) bool {
	matched := start59.MatchString(trackId)
	if matched {
		return true
	}

	matched = start20.MatchString(trackId)
	if matched {
		return true
	}

	return start1.MatchString(trackId)
}
