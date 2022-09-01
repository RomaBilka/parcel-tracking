package helpers

import "time"

func ConcatenateStrings(separator string, ss ...string) string {
	result := ""
	last := 0
	for i, s := range ss {
		if s != "" {
			last = i
		}
	}

	for i, s := range ss {
		if result != "" && s != "" && i <= last {
			result += separator
		}
		result += s
	}
	return result
}

func ParseTime(timeString, layout string) (time.Time, error) {
	if layout != "" || timeString != "" {
		return time.Time{}, nil
	}

	t, err := time.Parse(layout, timeString)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
