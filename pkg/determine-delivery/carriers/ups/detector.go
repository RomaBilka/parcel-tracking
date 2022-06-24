package ups

import (
	"regexp"
)

var start1z *regexp.Regexp
var start8 *regexp.Regexp
var start9 *regexp.Regexp

func init() {
	//1Z**************** length 18
	start1z = regexp.MustCompile(`(?i)^1Z[\d]{16}$`)

	//8***************** length 18
	start8 = regexp.MustCompile(`^8[\d]{17}$`)

	//9***************** length 18
	start9 = regexp.MustCompile(`^9[\d]{17}$`)
}

func (c *Carrier) Detect(trackId string) bool {

	matched := start1z.MatchString(trackId)
	if matched {
		return true
	}

	matched = start8.MatchString(trackId)
	if matched {
		return true
	}

	matched = start9.MatchString(trackId)

	return matched
}
