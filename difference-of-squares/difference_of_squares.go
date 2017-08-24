// Package to find the difference between the sum of the squares and the square of the sum of the first N natural numbers.
package diffsquares

// Constant declarations.
const testVersion = 1

// SquareOfSums function to take in a number as a parameter,
// and return the square, of the sum of all of the elements of this number
// Eg. 1 + 2 + ... + 10)² = 55² = 3025.
func SquareOfSums(number int) int {
	var sumOfNumber int
	for i := 1; i <= number; i++ {
		sumOfNumber += i
	}
	return sumOfNumber * sumOfNumber
}

// SumOfSquares function to take in a number as a parameter,
// and return the sum, of the square of each of the elements of this number
// Eg. 1² + 2² + ... + 10² = 385.
func SumOfSquares(number int) int {
	var squareOfNumber int
	for i := 1; i <= number; i++ {
		squareOfNumber += i * i
	}
	return squareOfNumber
}

// Difference function to subtract the SumOfSquares from the SquareOfSums of a given number
func Difference(number int) int {
	return SquareOfSums(number) - SumOfSquares(number)
}
