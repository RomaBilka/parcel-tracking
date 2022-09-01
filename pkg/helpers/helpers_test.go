package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcatenateStrings(t *testing.T) {
	testCases := []struct {
		name    string
		strings []string
		result  string
	}{
		{
			name:    "s1, s2, s3",
			strings: []string{"s1", "s2", "s3"},
			result:  "s1, s2, s3",
		},
		{
			name:    "with empty",
			strings: []string{"", "s1", "s2", "", "s3", ""},
			result:  "s1, s2, s3",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			r := ConcatenateStrings(", ", testCase.strings...)
			assert.Equal(t, testCase.result, r)
		})
	}
}
