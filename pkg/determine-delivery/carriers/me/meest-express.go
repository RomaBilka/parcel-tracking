package me

import "regexp"

var startCV *regexp.Regexp
var startMYCV *regexp.Regexp

func init() {
	//CV999999999ZZ
	startCV = regexp.MustCompile(`(?i)^CV[\d]{9}[a-z][a-z]$`)

	//MYCV999999999ZZ
	startMYCV = regexp.MustCompile(`(?i)^MYCV[\d]{9}[a-z][a-z]$`)
}

type ME struct {
	TrackId string
}

func (m *ME) Detect() bool {
	matched := startCV.MatchString(m.TrackId)
	if matched {
		return true
	}

	matched = startMYCV.MatchString(m.TrackId)
	if matched {
		return true
	}

	return false
}
