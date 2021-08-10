// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "strings"

// ShareWith: A function that shares coordially with everyone
// If `name` exist then -> Output: One for `name`, one for me.
// If `name` does not exist -> Output: One for you, one for me.
func ShareWith(name string) string {
	s := strings.TrimSpace(name) // remove unnecessary whitespaces, if any

	if len(s) == 0 {
		return "One for you, one for me."
	}

	return "One for " + s + ", one for me."
}
