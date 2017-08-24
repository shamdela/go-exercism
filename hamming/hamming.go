// Package hamming to calculate the Hamming difference between two DNA strands.
package hamming

import "errors"

// Constant declaration.
const testVersion = 5

// Distance Function to calculate the hamming distance by counting the number of differences between two homologous DNA strands
func Distance(a, b string) (int, error) {
	// The Hamming distance is only defined for sequences of equal length
	if len(a) != len(b){
		return -1, errors.New("String are unequal in length")
	}

	var counter int // Counter to store the number of differences between each strand
	for idx := range a{
		if a[idx] != b[idx]{
			counter++
		}
	}
	return counter, nil
}
