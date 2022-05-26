package deliveries

import "regexp"

func IsUPS(str string) (bool, error) {
	matched, err := regexp.MatchString(`^1Z[\d]{16}`, str)
	if err != nil {
		return false, err
	} else if matched {
		return matched, nil
	}

	matched, err = regexp.MatchString(`^8[\d]{17}`, str)
	if err != nil {
		return false, err
	} else if matched {
		return matched, nil
	}

	matched, err = regexp.MatchString(`^9[\d]{17}`, str)
	if err != nil {
		return false, err
	} else if matched {
		return matched, nil
	}

	return false, nil
}

