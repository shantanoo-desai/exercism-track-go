package raindrops

import "strconv"

// Convert: Create Raindrop sounds based on
// whether the number is a factor of 3, 5, 7
// if factor of 3: 'Pling'
// if factor of 5: 'Plang'
// if factor of 7: 'Plong'
func Convert(input_number int) string {
	var raindrop_sound string

	if input_number%3 != 0 && input_number%5 != 0 && input_number%7 != 0 {
		// return the number in string format
		raindrop_sound = strconv.Itoa(input_number)
	}

	// Check all factors individually

	if input_number%3 == 0 {
		raindrop_sound += "Pling"
	}

	if input_number%5 == 0 {
		raindrop_sound += "Plang"
	}

	if input_number%7 == 0 {
		raindrop_sound += "Plong"
	}

	return raindrop_sound
}
