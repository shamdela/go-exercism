/* Package palindrome that given a range of integers, returns the smallest and largest
palindromic product within that range, along with all of it's factors.*/
package palindrome

import (
	"errors"
	"strconv"
)

// Constant declarations
const testVersion = 1

type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

// Products function to detect palindrome products in a given range.
func Products(fmin, fmax int) (pmin Product, pmax Product, err error) {
	var minSet bool

	if fmin > fmax {
		err = errors.New("fmin > fmax...")
	} else {
		for i := fmin; i <= fmax; i++ {
			for j := fmin; j <= fmax; j++ {
				product := i * j
				if isPalindrome(product) {
					// Check smallest palindromic number
					if !minSet || product < pmin.Product {
						pmin.Product = product
						// if numbers are already in min array, don't add them again
						if !contains(i, j, pmin.Factorizations) {
							pmin.Factorizations = [][2]int{{i, j}}
							minSet = true
						}
					} else if product == pmin.Product {
						pmin.Product = product
						// if numbers are already in min array, don't add them again
						if !contains(i, j, pmin.Factorizations) {
							pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
							minSet = true
						}
					}

					// Check largest palindromic number
					if product > pmax.Product {
						pmax.Product = product
						// if numbers are already in max array, don't add them again
						if !contains(i, j, pmax.Factorizations) {
							pmax.Factorizations = [][2]int{{i, j}}
						}

					} else if product == pmax.Product {
						pmax.Product = product
						// if numbers are already in max array, don't add them again
						if !contains(i, j, pmax.Factorizations) {
							pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
						}
					}
					continue
				}
				continue
			}
		}
		if len(pmin.Factorizations) <= 0 && len(pmax.Factorizations) <= 0 {
			err = errors.New("no palindromes...")
		}
	}
	return pmin, pmax, err
}

// contains function to check if 2 numbers are in a given 2d array
func contains(x, j int, factorizations [][2]int) bool {
	for intPair := range factorizations {
		if (x == factorizations[intPair][0] && j == factorizations[intPair][1]) ||
			(j == factorizations[intPair][0] && x == factorizations[intPair][1]) {
			return true
		}
	}
	return false
}

// isPalindrome function to check if a mumber is a Palindrome
func isPalindrome(number int) bool {
	numberStr := strconv.Itoa(number)
	halfNum := len(numberStr) / 2

	for i := 0; i < halfNum; i++ {
		if numberStr[i] != numberStr[len(numberStr)-i-1] {
			return false
		}
	}
	return true
}
