// Package containing a function to determine if a sentence is a pangram.
// A pangram is a sentence using every letter of the alphabet at least once.
package pangram

import "strings"

// Constant declarations.
const (
	testVersion = 1
	lowercaseA = 97
	lowercaseZ = 122
)

// IsPangram function to determine if a sentence is a pangram.
func IsPangram(sentence string) bool {
	// Check if sentence passed in contains each of the letters between ASCII 97 and 122
	for i := lowercaseA; i <= lowercaseZ; i++ {
		if !strings.ContainsRune(strings.ToLower(sentence), rune(i)) {
			return false
		}
	}
	// All letters are present at least once -> sentence == pangram
	return true
}
