// Package pascal to compute Pascal's triangle up to a given number of rows. In Pascal's Triangle each number is
// computed by adding the numbers to the right and left of the current position in the previous row
package pascal

// Constant declarations
const testVersion = 1

// Triangle function to output each number is computed by adding the numbers to the right and left of the current
// position in the previous row.
func Triangle(rows int) [][]int {
	// Construct the 2d list and initialise it with a slice with an int value of 1
	var twoDList [][]int
	twoDList = append(twoDList, []int{1})

	// Loop for number of rows passed into the function
	for row := 0; row < rows; row++ {
		// Construct an empty slice of int, representing the inner values of the triangle
		var list []int

		// Loop for the number of times + 1, of the value of each row
		// Eg, row 1 = loop twice, row 2 = loop trice, row 3 = loop 4 times etc
		for index := 0; index <= row; index++ {
			// We already have first row in list initialised with 1, so check where row > 0
			if row > 0 {
				var val1, val2 int // Initialises these variables with 0

				// If we have a value on previous row, at same index as current one - set it in val1
				// Otherwise val1 is 0
				if len(twoDList[row-1]) > index {
					val1 = twoDList[row-1][index]
				}

				// If we have a value on previous row, at index before current one - set it in val2
				// Otherwise val1 is 0
				if index > 0 {
					val2 = twoDList[row-1][index-1]
				}

				// Add val1 and val2 from previous row and set result in current position in list
				list = append(list, val1+val2)
			}
		}
		if row > 0 {
			// If we are passed the first row, append inner list to 2Dlist of reults
			twoDList = append(twoDList, list)
		}
	}
	return twoDList
}
