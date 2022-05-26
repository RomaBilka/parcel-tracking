package deliveries

import "regexp"

func IsNpShopping(str string) (bool, error) {
	matched, err := regexp.MatchString(`^[\d]{14}$`, str)
	if err != nil {
		return false, err
	} else if matched {
		return matched, err
	}

	return false, nil
}
