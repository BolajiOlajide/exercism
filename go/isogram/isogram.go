package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns a boolean to determine if the argument is an isogram
func IsIsogram(input string) bool {
	var seen = map[rune]bool{}

	for _, r := range strings.ToLower(input) {
		if !unicode.IsLetter(r) {
			continue
		}

		if seen[r] {
			return false
		}

		seen[r] = true
	}

	return true
}
