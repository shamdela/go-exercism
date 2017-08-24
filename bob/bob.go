// package that defines a function to respond to questions as Bob would
package bob

import "strings"

const testVersion = 2 // same as targetTestVersion

// Hey function to reply to questions, given a set of rules
func Hey(question string) string {
	// Trim leading and trailing spaces on string
	question = strings.TrimSpace(question)

	if question == "" {
		return "Fine. Be that way!"
	} else if question == strings.ToUpper(question) && question != strings.ToLower(question) {
		// If string is all uppercase
		return "Whoa, chill out!"
	} else if strings.ContainsRune(string(question[len(question)-1]), '?') {
		// If last character of string is a ?
		return "Sure."
	}
	return "Whatever."
}