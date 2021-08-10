package hamming

import "errors"

// Distance: Function to determine Hamming Distance
// for two DNA Strands of equal length
func Distance(a, b string) (int, error) {
	if len(a) == 0 && len(b) == 0 { // nothing provided
		return 0, nil
	}

	if (len(a) == 0 && len(b) != 0) || (len(a) != 0 && len(b) == 0) {
		return 0, errors.New("only one input strand provided")
	}

	if len(a) > 0 && len(b) > 0 {
		if len(a) == len(b) { // check only if the length of the strands are equal
			var hamming_distance int = 0 // Initial Hamming Distance
			for index := range a {       // since same length strands range over the length
				if a[index] != b[index] { // check if the letters don't match
					hamming_distance += 1 // Increase Hamming Distance if they mismatch
				}
			}
			return hamming_distance, nil
		}
		return 0, errors.New("unequal strand lengths provided") // return error if length mismatches
	}

	return 0, nil // default to 0 distance and no error
}
