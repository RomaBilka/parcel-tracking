package deliveries

import "regexp"

func IsNovaPoshta(str string) (bool, error) {
	//59************ length 14
	matched, err := regexp.MatchString(`^59[\d]{12}$`, str)
	if err != nil {
		return false, err
	} else if matched {
		return matched, nil
	}

	//20************ length 14
	matched, err = regexp.MatchString(`^20[\d]{12}$`, str)
	if err != nil {
		return false, err
	} else if matched {
		return matched, nil
	}

	//1************* length 14
	matched, err = regexp.MatchString(`^1[\d]{13}$`, str)
	if err != nil {
		return false, err
	} else if matched {
		return matched, nil
	}

	return false, nil
}
