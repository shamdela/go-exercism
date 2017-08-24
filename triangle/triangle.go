// Package to determine if a triangle is equilateral, isosceles, or scalene.
package triangle

import "math"

const testVersion = 3
const (
	NaT Kind = "NaT" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)

type Kind string

// KindFromSides function got return the kind of triangle from given sides
func KindFromSides(a, b, c float64) Kind {
	if a>0 && b>0 && c>0 {
		//  Checks whether the floats are positive infinity. If so -> not a triangle
		if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
			return Kind(NaT)
		}

		if (a+b >= c) && (a+c >= b) && (b+c >= a) {
			if a == b && a == c {
				return Kind(Equ)
			} else if a == b || a == c || b == c {
				return Kind(Iso)
			}
			return Kind(Sca)
		}
	}
	return Kind(NaT)
}
