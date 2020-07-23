//Package darts implements a function to score a dart throw
package darts

import (
	"math"
)

//Score returns the score of a dart throw based on passed coordinate parameters.
func Score(x, y float64) int {
	score := 0
	distance := math.Abs(math.Pow((x-0), 2)) + math.Abs(math.Pow((y-0), 2))
	switch {
	case distance > 100: //outside circle
		score = 0
	case distance > 25: //outer circle
		score = 1
	case distance > 1: //middle circle
		score = 5
	case distance <= 1: //inner circle
		score = 10
	}

	return score
}
