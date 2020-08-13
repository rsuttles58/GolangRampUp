package triangle

import (
	"sort"
	"math"
)

type Kind string

const (
	NaT = "Not a triangle"
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	//if any side is zero, not a number, or infinity, 
	//return Not a Triangle
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}

	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return NaT
	}

	//if the two shortest sides don't equal longest, Not a Triangle
	sides := []float64{a, b, c}
	sort.Float64s(sides)

	if sides[0]+sides[1] < sides[2]  {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	if a != b && b != c && a != c {
		return Sca
	}

	return NaT
}
