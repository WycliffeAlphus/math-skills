package variance

import "math-skills/average"

func Variance(numbers []int) float64 {
// variance is the the sum of the difference between each data point and the mean divide by the number of data points
	mean := average.Average(numbers)
	var sum float64 = 0
	for _, n := range numbers {
		sum += ((float64(n) - mean) * (float64(n) - mean)) // sum of squares
	}
	v := sum / float64(len(numbers))
	return v
}
