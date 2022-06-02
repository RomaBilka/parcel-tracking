package deliveries

import "regexp"

func IsNpShopping(str string) (bool, error) {
	//NP99999999999999NPG
	matched, err := regexp.MatchString(`(?i)^NP[\d]{14}NPG$`, str)
	if err != nil {
		return false, err
	} else if matched {
		return matched, err
	}

	return false, nil
}
