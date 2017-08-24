// A Package containing a function that can return if a year is a leap year or not.
// A leap year is;
// 	- a year that is evenly divisible by 4
// 	- except every year that is evenly divisible by 100
//	- unless the year is also evenly divisible by 400
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 3

// A function to take a year as param and return a boolean if the year is a leap year or not
func IsLeapYear(year int) bool {
	if year % 4 == 0 && (year % 100 != 0 || year % 400 == 0) {
		return true
	}
	return false
}
