// Package slice which, given a string of digits, output all the contiguous substrings of length n in that string.
package slice

// Constant declarations
const testVersion = 1

// All function to loop through passed in strings' chars and return a slice of strings representing all the
// contiguous substrings of length n in that string.
func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}
	var result []string
	for idx := 0; idx < len(s)-(n-1); idx++ {
		// slice the string from idx to idx + n, and append result to the slice
		result = append(result, s[idx:idx+n])
	}
	return result
}

// UnsafeFirst function that returns the first n chars from the string s
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First function to check that the number wasn't greater than the len of string, then returns first n chars
// Returns first n chars of string plus a boolean to flag success or fail
func First(n int, s string) (first string, ok bool) {
	if len(s) >= n {
		first = s[:n]
		ok = true
	}
	return first, ok
}
