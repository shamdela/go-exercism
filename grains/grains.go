// Package grains which, calculates the number of grains of wheat on a chessboard
// given that the number on each square doubles.
package grains

import (
	"errors"
	"math"
)

// Constant declarations
const testVersion = 1

// Square function to calculate how many grains were on each square
func Square(input int) (actualVal uint64, actualError error) {
	if input < 1 || input > 64 {
		return 0, errors.New("Input is outside te bounds of a chessboard")
	}
	// Exponential calculation
	return uint64(math.Pow(2, float64(input-1))), nil
}

// Total function to calculate the total number of grains on all squares
func Total() uint64 {
	// 1<<64 - 1
	var total uint64
	for i := 1; i <= 64; i++ {
		sum, _ := Square(i)
		total += sum
	}
	return total
}
