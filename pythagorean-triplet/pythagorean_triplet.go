// Package pythagorean which, finds the product a * b * c.
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
package pythagorean

// Constant declarations
const testVersion = 1

// Triplet representation - slice of 3 ints
type Triplet [3]int

// Sum function returns a list of all Pythagorean triplets where the sum a+b+c (the perimeter) is equal to p.
func Sum(p int) []Triplet {

	var tripletHolder []Triplet
	tripletsInRange := Range(1, p)

	for _, triplet := range tripletsInRange {
		sum := 0
		for _, num := range triplet {
			sum += num
		}
		if sum == p {
			tripletHolder = append(tripletHolder, triplet)
		}
	}
	return tripletHolder
}

// Range function returns a list of all Pythagorean triplets with sides in the range min to max inclusive.
// Formula = a**2 + b**2 = c**2
func Range(min, max int) []Triplet {

	var tripletHolder []Triplet

	for a := min; a < max+1; a++ {
		for b := a + 1; b < max+1; b++ {
			for c := b + 1; c < max+1; c++ {
				if a*a+b*b == c*c {
					triplet := Triplet{a, b, c}
					tripletHolder = append(tripletHolder, triplet)
				}
			}
		}
	}
	return tripletHolder
}
