// Package to validate given 16 digit number as string
// via Luhn Algorithm
package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

// Valid implements the Luhn Algorithm
func Valid(input_number string) bool {

	// remove whitespaces between the number sets
	scrubbed := strings.ReplaceAll(input_number, " ", "")
	total_digits := len(scrubbed)

	if total_digits < 2 {
		return false
	}

	special_char_regex := regexp.MustCompile(`\D`)

	if special_char_regex.MatchString(scrubbed) {
		return false
	}

	var sum int
	var parity bool

	for idx := total_digits - 1; idx >= 0; idx-- {
		digit, _ := strconv.Atoi(string(scrubbed[idx])) // start rightmost side
		if parity {
			digit *= 2 // double the number
		}
		parity = !parity // flip parity for the next bit
		if digit > 9 {
			digit -= 9
		}
		sum += digit
	}

	return sum%10 == 0
}
