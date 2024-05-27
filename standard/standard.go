package standard

import (
	"math"

	"math-skills/variance"
)

func Deviation(numb []int) float64 {
	// Calculate the variance of the input numbers; call the variance function with the numb slice as the argument
	v := variance.Variance(numb)

	// Calculate the standard deviation by taking the square root of the variance
	standard := math.Sqrt(v)

	// Return the standard deviation
	return standard
}
