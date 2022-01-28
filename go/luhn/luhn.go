package luhn

import (
	"regexp"
	"strconv"
)

// Valid validates an account number to know if it's valid
func Valid(id string) bool {
	whiteSpaceRegex := regexp.MustCompile(`\s+`)
	trimmedID := whiteSpaceRegex.ReplaceAllString(id, "")
	lenOfID := len(trimmedID)

	if id == "" || lenOfID <= 1 {
		return false
	}

	var sum int
	var shouldTransform bool
	var newSequence []int

	for index := lenOfID - 1; index >= 0; index-- {
		integer, err := strconv.Atoi(string(trimmedID[index]))
		if err != nil {
			return false
		}

		if shouldTransform {
			integer = integer * 2

			if integer > 9 {
				integer -= 9
			}
		}

		sum += integer
		newSequence = append(newSequence, integer)
		shouldTransform = !shouldTransform
	}

	return sum%10 == 0
}
