package average

// It returns the mean of given data sets
func Average(numbers []int) float64 {
	if len(numbers) < 1 {
		return 0
	}
	var mean float64
	var total float64

	for _, n := range numbers {
		total += float64(n)
		mean = total / float64(len(numbers))
	}
	return mean
}
