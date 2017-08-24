// Package to if a sentence is a pangram..
package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 2

// Abbreviate function to convert a long name to its acronym, using the regexp module
func Abbreviate(sentence string) string {
	var tla string

	// Compile the regular expression to search for.
	re := regexp.MustCompile(`[A-Z]?[a-z]+|[A-Z]+`)

	// Split based on pattern. Second argument means "no limit."
	words := re.FindAllString(sentence, -1)

	for _, word := range words{
		// Convert first letter of each word to Uppercase and append to tla string
		tla += strings.ToUpper(string(word[0]))
	}

	return tla
}
