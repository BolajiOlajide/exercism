package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns a boolean to determine if the argument is an isogram
func IsIsogram(input string) bool {
	var tempMap = make(map[rune]int)

	for _, char := range strings.ToLower(input) {
		if !unicode.IsLetter(char) {
			continue
		}

		if tempMap[char] == 0 {
			tempMap[char] = 1
			continue
		}

		return false
	}

	return true
}
