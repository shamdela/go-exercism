// Package isogram determines, if a word or phrase is an isogram.
// An isogram (also known as a "nonpattern word") is a word or phrase without a repeating letter.
package isogram

import "strings"

// Constant declarations
const testVersion = 1

// IsIsogram function returns true if the string passed in has no repeating letters.
func IsIsogram(input string) bool {

	// Replacer replaces all occurrance spaces and '-' with an empty string
	replacer := strings.NewReplacer(" ", "", "-", "")
	input = replacer.Replace(input)

	// A map data collection to hold each char in the string as its key,
	// with an int value (representing no. of occurrences) as its value
	letterStore := make(map[rune]int)
	for _, char := range strings.ToLower(input) {
		if _, ok := letterStore[char]; ok {
			return false
		}
		letterStore[char]++
	}
	return true
}
