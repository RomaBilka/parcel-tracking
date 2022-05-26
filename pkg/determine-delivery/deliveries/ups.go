package deliveries

import "regexp"

func IsUPS(str string) (bool, error) {
	//1Z**************** length 18
	matched, err := regexp.MatchString(`(?i)^1Z[\d]{16}$`, str)
	if err != nil {
		return false, err
	} else if matched {
		return matched, nil
	}

	//8***************** length 18
	matched, err = regexp.MatchString(`^8[\d]{17}$`, str)
	if err != nil {
		return false, err
	} else if matched {
		return matched, nil
	}

	//9***************** length 18
	matched, err = regexp.MatchString(`^9[\d]{17}$`, str)
	if err != nil {
		return false, err
	} else if matched {
		return matched, nil
	}

	return false, nil
}
