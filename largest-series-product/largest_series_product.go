// Package lsproduct, when given a string of digits, calculate the largest product for a contiguous
// substring of digits of length n.
package lsproduct

import (
	"errors"
	"sort"
	"strconv"
)

// Constant declarations
const testVersion = 5

// LargestSeriesProduct function to calculate the largest product for a contiguous substring of digits of length n.
func LargestSeriesProduct(digits string, span int) (int, error) {
	var runningTotal int  // holds current product sum
	var totalHolder []int // holds all sums

	if span > len(digits) || span < 0 {
		return -1, errors.New("Invalid span value provided")
	}

	// Loop through string input of digits
	for idx, _ := range digits {
		runningTotal = 1
		maxRange := idx + span

		if maxRange <= len(digits) {
			// Take a slice of the digits from current index, up to maxRange/span
			nums := digits[idx:maxRange]

			// Loop through these nums
			for i, _ := range nums {
				// Convert each string digit to an int
				val, err := strconv.Atoi(string(nums[i]))
				if err != nil {
					return -1, errors.New("Error encountered parsing digit input")
				}
				runningTotal *= val
			}
			totalHolder = append(totalHolder, runningTotal)
		}
	}

	// If totalHolder slice contains no values, and no errors have occurred, return positively
	if len(totalHolder) < 1 {
		return 1, nil
	}

	// Sort contents of totalHolder slice in descending order and return highest value
	sort.Sort(sort.Reverse(sort.IntSlice(totalHolder)))
	return int(totalHolder[0]), nil
}
