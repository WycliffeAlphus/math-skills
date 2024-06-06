package variance

import "math-skills/average"

// Variance measures the spread of the data points from the mean.
//
// It is the average of the squared differences from the mean.
func Variance(numbers []int) float64 {
	if len(numbers) < 1 {
		return 0
	}
	mean := average.Average(numbers)
	var sum float64 = 0
	for _, n := range numbers {
		sum += ((float64(n) - mean) * (float64(n) - mean))
	}
	v := sum / float64(len(numbers))
	return v
}
