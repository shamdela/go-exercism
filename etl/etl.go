// Package etl which, extracts some scrabble scores from a legacy system.
package etl

import "strings"

// Constant declarations
const testVersion = 1

// Transform function to loop through passed in map of slices based on int keys.
// The output will stores the score per letter, which makes it much faster and easier to calculate the score for a word.
// It also stores the letters in lower-case regardless of the case of the input letter
func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)
	for k, v := range input {
		for i := 0; i < len(v); i++ {
			output[strings.ToLower(v[i])] = k
		}
	}
	return output
}
