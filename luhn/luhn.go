// Package luhn which given a number determine whether or not it is valid per the Luhn formula
package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

// Constant declarations
const testVersion = 2

// Valid function which given string input, checks if number is a valid according to Luhn algorithm
func Valid(input string) bool {

	var runningTotal int

	// Remove spaces
	normalisedString := strings.Replace(input, " ", "", -1)

	// Validate input
	if checkValidInput(normalisedString) {
		return false
	}

	// Perform Luhn algorithm
	for i := len(normalisedString) - 1; i >= 0; i-- {
		number, _ := strconv.Atoi(string(normalisedString[i]))
		if (len(normalisedString)-i)%2 == 0 {
			number *= 2
			if number > 9 {
				number -= 9
			}
		}
		runningTotal += number
	}

	if runningTotal%10 == 0 {
		return true
	}

	return false
}

// checkValidInput function which given a string, returns true if it has a length less than 2, or if it contains non
// numeric characters.
func checkValidInput(input string) bool {

	re := regexp.MustCompile(`[^\d ]+`) // find all digits - 0-9
	match := re.MatchString(input)
	invalidLength := len(input) <= 1
	return match || invalidLength
}
