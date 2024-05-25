package median

func Median(numbers []int) float64 {
	var median float64

	currIndex := len(numbers) / 2
	if len(numbers)%2 == 0 { // if the number of available data points is even
		median = (float64(numbers[currIndex]) + float64(numbers[currIndex-1])) / 2
	} else if len(numbers)%2 != 0 {
		median = (float64(numbers[currIndex])) // if the number of available data points is odd
	}
	return median
}
