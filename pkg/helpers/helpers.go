package helpers

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
