// Isogram package to check whether a worth comprises
// only of unique letters and hence is an Isogram
package isogram

import "strings"

// IsIsogram: Checks whether all the letters in the word
// are unique and none are repeating
func IsIsogram(word string) bool {

	if len(word) == 0 {
		return true // empty string is an isogram
	}

	// Create a Dynamic memory for the Counter Map
	letter_count_map := make(map[string]int)

	// Scrub off all `-` from the words
	scrubbed_word := strings.ReplaceAll(strings.ToLower(word), "-", "")

	// Remove spaces between words
	scrubbed_word = strings.ReplaceAll(scrubbed_word, " ", "")

	// Iterate over all letters in the scrubbed word and create a count map
	for _, each_letter := range scrubbed_word {
		// create a map of each letter and the count of occurence
		letter_count_map[string(each_letter)] = strings.Count(scrubbed_word, string(each_letter))
	}

	// parse through all the counted values and check if all are unique
	for _, count := range letter_count_map {
		if count > 1 {
			return false // if any letter occurs more than once -> not an isogram
		}
	}
	return true
}
