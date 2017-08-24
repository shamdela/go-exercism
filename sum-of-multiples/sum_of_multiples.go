// Package summultiples which given a number, finds the sum of all the multiples of particular numbers up to but not
// including that number.
package summultiples

// Constant declarations
const testVersion = 1

// SumMultiples function to iterate up to the passed in 'limit' and finds the sum of all the multiples of particular
// numbers up to but not including that number.
func SumMultiples(limit int, divisors ...int) int {
	var setOfMultiples []int // Slice of int to store multiples
	var sum int
	// Iterate from 1 upto the number 'limit', and then iterate over the divisors in the passed in list
	// If index is a multiple of a divisor, and its not already in out list, add it to our list.
	for i := 1; i < limit; i++ {
		for _, divisor := range divisors {
			if i%divisor == 0 {
				if !contains(setOfMultiples, i) {
					setOfMultiples = append(setOfMultiples, i)
					sum += i
				}
			}
		}
	}
	return sum
}

// contains function to iterate over the setOfMultiples to check that if a number has already been added to the slice
func contains(setOfM []int, value int) bool {
	for _, val := range setOfM {
		if val == value {
			return true
		}
	}
	return false
}
