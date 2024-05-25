package average

import (
	"fmt"
	"math"
)

func Average(number []int) float64 {
	var mean float64
	var total float64

	for _, n := range number {
		if n <= math.MaxInt || n >= math.MinInt {
			total += float64(n)
			mean = total / float64(len(number))
		} else {
			fmt.Printf("The %v is not within the expected range\n", number)
		}
	}
	return mean
}
