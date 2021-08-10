// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Responses by Bob
const (
	sure     = "Sure."
	whoa     = "Whoa, chill out!"
	calm     = "Calm down, I know what I'm doing!"
	fine     = "Fine. Be that way!"
	whatever = "Whatever."
)

// Bob's Lackadaisical Function
// Response to Question: 'Sure'
// Response to Shouts: 'Whoa, chill out!'
// Response to Shouted Questions: 'Calm down, I know what I'm doing!'
// Response to no input: 'Fine, Be that way!'
// Response to anything else: 'Whatever.'
func Hey(remark string) string {

	s := strings.TrimSpace(remark)
	if len(s) == 0 { // Said Nothing
		return fine
	}
	if s[len(s)-1] == '?' {
		// check for questions
		if strings.ToUpper(s) == s && strings.ToLower(s) != s { // shouted question
			return calm
		}
		return sure // normal question
	}

	if strings.ToUpper(s) == s && strings.ToLower(s) != s { // Shout
		return whoa
	}
	return whatever
}
