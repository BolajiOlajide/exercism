package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns a boolean to determine if the argument is an isogram
func IsIsogram(input string) bool {
	var tempMap = make(map[rune]bool)

	for _, char := range strings.ToLower(input) {
		if !unicode.IsLetter(char) {
			continue
		}

		if tempMap[char] {
			return false
		}

		tempMap[char] = true
		continue
	}

	return true
}
