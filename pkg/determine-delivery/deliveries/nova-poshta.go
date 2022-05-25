package deliveries

import "regexp"

func IsNovaPoshta(str string) (bool, error) {
	matched, err := regexp.MatchString(`^59[\d]{12}`, str)
	if err != nil {
		return false, err
	} else if matched {
		return matched, nil
	}

	matched, err = regexp.MatchString(`^20[\d]{12}`, str)
	if err != nil {
		return false, err
	} else if matched {
		return matched, nil
	}

	matched, err = regexp.MatchString(`^1[\d]{13}`, str)
	if err != nil {
		return false, err
	} else if matched {
		return matched, nil
	}

	return false, nil
}
