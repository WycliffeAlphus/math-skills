package standard

import (
	"math"
	"math-skills/variance"
)

func Deviation(numb []int) float64 {

	v := variance.Variance(numb)
	standard := math.Sqrt(v)
	return standard
}
