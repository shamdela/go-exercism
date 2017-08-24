// Package cryptosquare to implement the classic method for composing secret messages called a square code.
package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

// Constant declarations
const testVersion = 2

// Encode function which given an English text, outputs the encoded version of that text.
func Encode(pt string) string {
	var result []rune

	normalisedString := normalise(pt)
	k := int(math.Ceil(math.Pow(float64(len(normalisedString)), 0.5)))

	for x := 0; x < k; x++ {
		for y := 0; y < k; y++ {
			i := y*k + x
			if i < len(normalisedString) {
				result = append(result, rune(normalisedString[i]))
			}
		}
		if x < k-1 {
			result = append(result, ' ')
		}
	}
	return string(result)
}

// normalise function which given a string, removes the spaces and punctuation and converts to lowercase.
func normalise(pt string) string {
	re := regexp.MustCompile(`[\w]*`)                  // find all word characers - a-z, A-Z, 0-9
	match := re.FindAllString(strings.ToLower(pt), -1) // convert string to lowercase
	return strings.Join(match, "")
}
