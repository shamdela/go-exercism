// Package to convert a number to a string, the contents of which depend on the number's factors.
package raindrops

import "strconv"

// Constant declaration.
const testVersion = 2

// Function to calculate the hamming distance by counting the number of differences between two homologous DNA strands
func Convert(number int) string {
	var output string // Our output string
	var not357 bool   // State of whether number is a factor of 3, 5 or 7

	if number%3 == 0 {
		output += "Pling"
		not357 = true
	}
	if number%5 == 0 {
		output += "Plang"
		not357 = true
	}
	if number%7 == 0 {
		output += "Plong"
		not357 = true
	}
	if !not357 {
		// Number is not a factor of 3, 5 or 7, so convert the number to a string
		output += strconv.Itoa(number)
	}
	return output
}
