package isogram

import (
	"regexp"
	"strings"
)

// IsIsogram returns a boolean to determine if the argument is an isogram
func IsIsogram(input string) bool {
	inputInByte := []byte(strings.ToLower(input))

	invalidCharRegex := regexp.MustCompile(`[- ]`)
	sanitizedInput := invalidCharRegex.ReplaceAll(inputInByte, []byte(""))

	var tempMap = make(map[byte]int)

	for _, char := range sanitizedInput {
		if tempMap[char] == 0 {
			tempMap[char] = 1
			continue
		}
		return false
	}

	return true
}
