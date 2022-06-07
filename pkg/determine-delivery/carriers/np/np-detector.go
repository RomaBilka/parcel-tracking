package np

import "regexp"

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

type NP struct {
	TrackId string
}

func (n *NP) Detect() bool {
	matched := start59.MatchString(n.TrackId)
	if matched {
		return true
	}

	matched = start20.MatchString(n.TrackId)
	if matched {
		return true
	}

	matched = start1.MatchString(n.TrackId)
	if matched {
		return true
	}

	return false
}