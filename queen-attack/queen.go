// Package queenattack which, given the position of two queens on a chess board, indicate
// whether or not they are positioned so that they can attack each other.
package queenattack

import (
	"errors"
	"math"
)

// Constant declarations
const testVersion = 2

// CanQueenAttack function will calculate, given two strings representing a chessboard position of two queens,
// whether they can attack each other or not. Returns true if they may attack, and false if not. Error also returned as
// part of multiple return value.
func CanQueenAttack(whiteQueen string, blackQueen string) (attack bool, err error) {
	if whiteQueen == blackQueen {
		// If both queens are on same board position, throw an error
		err = errors.New("same position")

	} else if whiteQueen[0] > 'h' || blackQueen[0] > 'h' ||
		whiteQueen[1] < '1' || blackQueen[1] < '1' ||
		whiteQueen[1] > '8' || blackQueen[1] > '8' {
		// If either queen are in a position that's off the board
		err = errors.New("off the board")

	} else if whiteQueen[0] == blackQueen[0] || whiteQueen[1] == blackQueen[1] {
		// If queens are on same row OR same column, they can attack each other
		attack = true
	} else {
		// Otherwise, find difference between positions of col and row for each queen and
		// convert these to absolute values.
		rowDiff := math.Abs(float64(int(whiteQueen[0]) - int(blackQueen[0])))
		colDiff := math.Abs(float64(int(whiteQueen[1]) - int(blackQueen[1])))

		// If the differences are equal, the queens are on diagonal positions and may attack
		if rowDiff == colDiff {
			attack = true
		}
	}
	return attack, err
}
