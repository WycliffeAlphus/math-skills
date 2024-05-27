package variance

import "math-skills/average"

func Variance(numbers []int) float64 {
	// Calculate the mean of the input numbers
	mean := average.Average(numbers)
	var sum float64 = 0
	// Calculate the sum of squares of the differences between each number and the mean
	for _, n := range numbers {
		sum += ((float64(n) - mean) * (float64(n) - mean)) // sum of squares
	}
	// Calculate the variance by dividing the sum of squares by the number of data points
	v := sum / float64(len(numbers))
	return v
}
