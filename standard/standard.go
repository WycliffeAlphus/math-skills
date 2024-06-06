package standard

import (
	"math"

	"math-skills/variance"
)

// Returns the standard deviation
//
// Square Root of the variance
func Deviation(numb []int) float64 {
	v := variance.Variance(numb)
	return math.Sqrt(v)
}
