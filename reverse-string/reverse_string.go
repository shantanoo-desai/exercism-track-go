// reverse package to reverse the string
package reverse

func Reverse(word string) string {
	r := []rune(word)   // convert incoming word into an array of runes
	var reversed []rune // resultant reversed rune array

	if len(word) == 0 { // if empty string
		return string(reversed)
	}

	for i := len(r) - 1; i >= 0; i-- {
		reversed = append(reversed, r[i])
	}
	return string(reversed)
}
