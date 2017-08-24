// Package that converts a give decimal number, to the appropriate sequence of events for a secret handshake.
package secret

import (
	"strconv"
)

// Constant declarations.
const (
	testVersion = 1
	reverseAt   = 10000
)

// Map variable to store handshake codes
var handshakeMap = map[int]string{
	1:    "wink",
	10:   "double blink",
	100:  "close your eyes",
	1000: "jump",
}

// Handshake function to a decimal number as an unsigned int and return a slice of string, representing a sequence
// of events for that secret handshake.
func Handshake(code uint) []string {
	var result = make([]string, 0, 5)
	var shouldReverse bool

	// Using FormatUint to convert uint to binary string,
	// Using Atoi to convert it from string to an int
	binaryCode, _ := strconv.Atoi(strconv.FormatUint(uint64(code), 2))
	binaryCode %= 100000 // Only 32 bits allowed. So any larger, we cut it off.

	// If binary code > 10000, subtract 10000 from binary code and we now reverse array
	if binaryCode >= reverseAt {
		shouldReverse = true
		binaryCode -= reverseAt
	}

	if binaryCode >= 1000 {
		binaryCode, result = generateResult(binaryCode, shouldReverse, result, 1000)
	}
	if binaryCode >= 100 {
		binaryCode, result = generateResult(binaryCode, shouldReverse, result, 100)
	}
	if binaryCode >= 10 {
		binaryCode, result = generateResult(binaryCode, shouldReverse, result, 10)
	}
	if binaryCode == 1 {
		binaryCode, result = generateResult(binaryCode, shouldReverse, result, 1)
	}
	return result
}

// generateResult function to add secret instruction to slice.
func generateResult(binaryCode int, shouldReverse bool, result []string, size int) (int, []string) {
	// Add to front or back of slice, depending on shouldRevese boolean
	if shouldReverse {
		// Append to end of slice
		result = append(result, handshakeMap[size])
	} else {
		// Pre-pend to start of slice
		result = append([]string{handshakeMap[size]}, result...)
	}
	// Subtract the size away from binaryCode value and continue
	binaryCode -= size

	// Multiple return values
	return binaryCode, result
}
