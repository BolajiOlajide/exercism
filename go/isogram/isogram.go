package isogram

import (
	"regexp"
	"strings"
)

// bolaji
var invalidCharRegex *regexp.Regexp = regexp.MustCompile(`[- ]`)

// IsIsogram returns a boolean to determine if the argument is an isogram
func IsIsogram(input string) bool {
	sanitizedInput := invalidCharRegex.ReplaceAllString(strings.ToLower(input), "")

	var tempMap = make(map[rune]int)

	for _, char := range sanitizedInput {
		if tempMap[char] == 0 {
			tempMap[char] = 1
			continue
		}
		return false
	}

	return true
}
