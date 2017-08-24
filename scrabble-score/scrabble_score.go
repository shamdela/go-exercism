// Package scrabble which, given a word, computes the scrabble score for that word.
package scrabble

import "strings"

// Constant declarations
const testVersion = 4

// A map of letter-score values.
var scoreMap = map[rune]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1, 'l': 1, 'n': 1, 'r': 1, 's': 1, 't': 1,
	'd': 2, 'g': 2,
	'b': 3, 'c': 3, 'm': 3, 'p': 3,
	'f': 4, 'h': 4, 'v': 4, 'w': 4, 'y': 4,
	'k': 5,
	'j': 8, 'x': 8,
	'q': 10, 'z': 10,
}

// Score function to compute the scrabble score for a given word based off a Letter-Value map.
// A string parameter representing the word is passed in to the function and an int value representing the total score
// for that word, is returned.
func Score(input string) int {
	var score int
	for _, letter := range strings.ToLower(input) {
		score += scoreMap[letter]
	}
	return score
}
