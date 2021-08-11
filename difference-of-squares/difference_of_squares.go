package diffsquares

// SquareOfSum calculates the square of sum of N numbers
// Use the formula: N(N+1)/2
func SquareOfSum(N int) int {
	sum := N * (N + 1) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of squares of N numbers
// Use the formula: [N(N+1)(2N+1)]/6
func SumOfSquares(N int) int {
	return (N * (N + 1) * (2*N + 1)) / 6
}

// Difference returns the difference of Square of Sum and
// Sum of Squares for given N numbers
func Difference(N int) int {
	// Since (a + b ....+ N)^2 > a^2 + b^2 ..+N^2
	return SquareOfSum(N) - SumOfSquares(N)
}
