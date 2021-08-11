// Package acronym creates an acronym based on Capitalized Letters
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate: creates an acronym for a full form string
func Abbreviate(s string) string {
	var abbreviation = ""
	re := regexp.MustCompile(`\w'\w|(?:_|\b)([A-Za-z])`)
	for _, match := range re.FindAllStringSubmatch(s, -1) {
		abbreviation = abbreviation + match[1]
	}
	return strings.ToUpper(abbreviation)
}
