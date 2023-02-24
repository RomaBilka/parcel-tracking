package helpers

func ConcatenateStrings(separator string, ss ...string) string {
	var result string
	for i, s := range ss {
		if s == "" {
			continue
		}
		if i > 0 && result != "" {
			result += separator
		}
		result += s
	}
	return result
}
