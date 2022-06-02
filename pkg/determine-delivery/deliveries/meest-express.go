package deliveries

import "regexp"

func IsMeestExpress(str string) (bool, error) {
	//CV999999999ZZ
	matched, err := regexp.MatchString(`(?i)^CV[\d]{9}[a-z][a-z]$`, str)
	if err != nil {
		return false, err
	} else if matched {
		return matched, nil
	}

	//MYCV999999999ZZ
	matched, err = regexp.MatchString(`(?i)^MYCV[\d]{9}[a-z][a-z]$`, str)
	if err != nil {
		return false, err
	} else if matched {
		return matched, nil
	}

	return false, nil
}
