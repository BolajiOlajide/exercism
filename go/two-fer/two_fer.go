// Package twofer implements a solution to exercism two-fer's problem.
package twofer

import "fmt"

// ShareWith Share an item with somebody.
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
