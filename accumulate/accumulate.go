// Package accumulate which, given a collection and an operation to perform on each element of the collection,
// returns a new collection containing the result of applying that operation to each element of the input collection.
package accumulate

// Constant declarations
const testVersion = 1

// Accumulate function to take in a slice of string and a callback function to perform on each element of the slice,
// and returns a new slice of string containing the result of applying that callback function on each element of the input collection.
func Accumulate(given []string, converter func(string) string) []string {
	// Create a slice of string to store our results
	result := []string{}
	for _, str := range given {
		result = append(result, converter(str))
	}
	return result
}
