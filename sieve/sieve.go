// Package sieve which implements - Sieve of Eratosthenes algorithm, to find all the primes from 2 up to a given number.
package sieve

import "sort"

// Constant declarations
const testVersion = 1

// Sieve function to iteratively mark as composite (i.e. not prime) the multiples of each prime,
// starting with the multiples of 2.
func Sieve(s int) []int {
	var result []int
	compositeNumbers := make(map[int]bool)

	// Starting at two and continuing up to and including the given limit.
	for i := 2; i <= s; i++ {
		// Initialise map with the list of numbers and false as value, only if they are not already added to map.
		if !compositeNumbers[i] {
			compositeNumbers[i] = false
		}
		// Take the next available unmarked number in your list (it is prime)
		// Mark all the multiples of that number (they are not prime)
		for j := i; j <= s; j += i {
			if j > i {
				compositeNumbers[j] = true
			}
		}
	}

	// Loop through map to find numbers that are not marked as prime, and append them to our slice.
	for k, v := range compositeNumbers {
		if v == false {
			result = append(result, k)
		}
	}

	// Sort slice of int
	sort.Ints(result)
	return result
}
